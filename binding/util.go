package binding

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"

	"github.com/brevis-network/uniswap-rebate/webapi"
)

func (k PoolKey) Value() (driver.Value, error) {
	return json.Marshal(k)
}

func (k *PoolKey) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), k)
}

// to be saved in db. fields are args for ClaimWithZkProofs
type CallData struct {
	ProofIds          [][32]byte
	Proof             []byte
	ProofDataArray    []ProofData
	AppCircuitOutputs [][]byte
}

// convert to webapi.CallData
func (d *CallData) ToWebCallData() *webapi.CallData {
	ret := new(webapi.CallData)
	ret.Proof = ToHex(d.Proof)
	for _, p := range d.ProofIds {
		ret.ProofIds = append(ret.ProofIds, ToHex(p))
	}
	for _, o := range d.AppCircuitOutputs {
		ret.AppCircuitOutputs = append(ret.AppCircuitOutputs, ToHex(o))
	}
	for _, d := range d.ProofDataArray {
		ret.ProofDataArray = append(ret.ProofDataArray, &webapi.BrevisProofData{
			CommitHash:    ToHex(d.CommitHash),
			AppCommitHash: ToHex(d.AppCommitHash),
			AppVkHash:     ToHex(d.AppVkHash),
			SmtRoot:       ToHex(d.SmtRoot),
		})
	}
	return ret
}

// []byte or fixed length array [32]byte or [20]byte, has 0x prefix
func ToHex[T ~[]byte | ~[32]byte | ~[20]byte](input T) string {
	switch v := any(input).(type) {
	case []byte:
		return "0x" + hex.EncodeToString(v)
	case [20]byte:
	case [32]byte:
		return "0x" + hex.EncodeToString(v[:])
	}
	return ""
}
