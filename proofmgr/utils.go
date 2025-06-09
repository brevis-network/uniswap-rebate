package proofmgr

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"slices"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/sdk/proto/commonproto"
	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"github.com/brevis-network/brevis-sdk/sdk/proto/sdkproto"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CustomInputEntry struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func buildCustomInput(appCircuit sdk.AppCircuit) (*sdkproto.CustomInput, error) {
	appCircuitRef := reflect.ValueOf(appCircuit)
	// deref until we get the actual value
	for appCircuitRef.Kind() == reflect.Pointer {
		appCircuitRef = appCircuitRef.Elem()
	}
	if appCircuitRef.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the concrete type of AppCircuit must be struct")
	}
	customInputMap := make(map[string]any)
	fields := reflect.VisibleFields(appCircuitRef.Type())
	for i, field := range fields {
		fieldValue := appCircuitRef.Field(i)
		switch field.Type.Kind() {
		case reflect.Array, reflect.Slice:
			var entries []CustomInputEntry
			for j := 0; j < fieldValue.Len(); j++ {
				entries = append(entries, CustomInputEntry{
					Type: field.Type.Elem().Name(),
					Data: fmt.Sprintf("%s", fieldValue.Index(j)),
				})
				customInputMap[field.Name] = entries
			}
		case reflect.Struct:
			data := fmt.Sprintf("%s", fieldValue)
			switch field.Type.Name() {
			case "Bytes32":
				data = "0x" + data
			}
			customInputMap[field.Name] = CustomInputEntry{
				Type: field.Type.Name(),
				Data: data,
			}
		}
	}
	jsonBytes, err := json.Marshal(customInputMap)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal err: %w", err)
	}

	return &sdkproto.CustomInput{JsonBytes: string(jsonBytes)}, nil
}

func convertSDKStorageToProtoStorage(storage sdk.StorageData) *sdkproto.StorageData {
	dataBytes, err := json.Marshal(storage)
	if err != nil {
		log.Errorf("marshal err %+v: %v", storage, err)
	}

	return &sdkproto.StorageData{
		BlockNum:           storage.BlockNum.Uint64(),
		Address:            hexutil.Encode(storage.Address.Bytes()),
		Slot:               storage.Slot.Hex(),
		Value:              storage.Value.Hex(),
		StorageDataJsonHex: hexutil.Encode(dataBytes),
	}
}

func convertSDKReceiptToProtoReceipt(receipt sdk.ReceiptData) *sdkproto.ReceiptData {
	fields := make([]*sdkproto.Field, len(receipt.Fields))
	for i, field := range receipt.Fields {
		fields[i] = convertSDKLogToProtoField(field)
	}

	dataBytes, err := json.Marshal(receipt)
	if err != nil {
		log.Errorf("fail to marshal %+v", receipt)
	}

	return &sdkproto.ReceiptData{
		BlockNum:           receipt.BlockNum.Uint64(),
		TxHash:             receipt.TxHash.Hex(),
		Fields:             fields,
		ReceiptDataJsonHex: hexutil.Encode(dataBytes),
	}
}

func convertSDKLogToProtoField(field sdk.LogFieldData) *sdkproto.Field {
	return &sdkproto.Field{
		IsTopic:    field.IsTopic,
		FieldIndex: uint32(field.FieldIndex),
		LogPos:     uint32(field.LogPos),
	}
}

func convertProtoReceiptToGWReceipt(receipt *sdkproto.ReceiptData) *gwproto.ReceiptInfo {
	fields := make([]*gwproto.LogExtractInfo, len(receipt.Fields))
	for i, field := range receipt.Fields {
		fields[i] = convertProtoFieldToGWField(field)
	}
	return &gwproto.ReceiptInfo{
		TransactionHash: receipt.TxHash,
		LogExtractInfos: fields,
	}
}

func convertProtoFieldToGWField(field *sdkproto.Field) *gwproto.LogExtractInfo {
	return &gwproto.LogExtractInfo{
		ValueFromTopic: field.IsTopic,
		ValueIndex:     uint64(field.FieldIndex),
		LogPos:         uint64(field.LogPos),
	}
}

func convertProtoStorageToGWStorage(storage *sdkproto.StorageData) *gwproto.StorageQueryInfo {
	return &gwproto.StorageQueryInfo{
		BlkNum:      storage.BlockNum,
		Account:     storage.Address,
		StorageKeys: []string{storage.Slot},
	}
}

// reqs and appInfos have same length
func buildGwQueries(reqs []*sdkproto.ProveRequest, appInfos []*commonproto.AppCircuitInfo) (ret []*gwproto.Query) {
	for i, req := range reqs {
		appCircuitInfo := appInfos[i]
		ret = append(ret, &gwproto.Query{
			ReceiptInfos: buildReceiptInfos(req),
			AppCircuitInfo: &commonproto.AppCircuitInfoWithProof{
				VkHash:           appCircuitInfo.VkHash,
				Toggles:          appCircuitInfo.Toggles,
				MaxReceipts:      appCircuitInfo.MaxReceipts,
				MaxStorage:       appCircuitInfo.MaxStorage,
				MaxTx:            appCircuitInfo.MaxTx,
				MaxNumDataPoints: appCircuitInfo.MaxNumDataPoints,
			},
		})
	}
	return
}

func buildReceiptInfos(req *sdkproto.ProveRequest) []*gwproto.ReceiptInfo {
	ret := make([]*gwproto.ReceiptInfo, len(req.Receipts))
	receipts := req.Receipts
	slices.SortFunc(receipts, func(a, b *sdkproto.IndexedReceipt) int {
		return int(a.Index) - int(b.Index)
	})
	for i, receipt := range receipts {
		ret[i] = convertProtoReceiptToGWReceipt(receipt.Data)
	}
	return ret
}

func buildStorageQueryInfos(req *sdkproto.ProveRequest) []*gwproto.StorageQueryInfo {
	ret := make([]*gwproto.StorageQueryInfo, len(req.Storages))
	storages := req.Storages
	slices.SortFunc(storages, func(a, b *sdkproto.IndexedStorage) int {
		return int(a.Index) - int(b.Index)
	})
	for i, storage := range storages {
		ret[i] = convertProtoStorageToGWStorage(storage.Data)
	}
	return ret
}

func getProverClient(proverRpc string) (sdkproto.ProverClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(proverRpc, opts...)
	if err != nil {
		return nil, fmt.Errorf("NewClient err: %w", err)
	}
	return sdkproto.NewProverClient(conn), nil
}

func toBigInt(i sdk.Uint248) *big.Int {
	b, _ := new(big.Int).SetString(i.String(), 0)
	return b
}

func toBigInts(is []sdk.Uint248) []*big.Int {
	var result []*big.Int
	for _, i := range is {
		result = append(result, toBigInt(i))
	}
	return result
}
