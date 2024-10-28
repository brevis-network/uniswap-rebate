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
	ProofDataArray    []BrevisProofData
	AppCircuitOutputs [][]byte
}

// convert to webapi.CallData
func (d *CallData) ToWebCallData() *webapi.CallData {
	ret := new(webapi.CallData)
	ret.Proof = toHex(d.Proof)
	for _, p := range d.ProofIds {
		ret.ProofIds = append(ret.ProofIds, toHex2(p))
	}
	for _, o := range d.AppCircuitOutputs {
		ret.AppCircuitOutputs = append(ret.AppCircuitOutputs, toHex(o))
	}
	for _, d := range d.ProofDataArray {
		ret.ProofDataArray = append(ret.ProofDataArray, &webapi.BrevisProofData{
			CommitHash:    toHex2(d.CommitHash),
			VkHash:        toHex2(d.VkHash),
			AppCommitHash: toHex2(d.AppCommitHash),
			AppVkHash:     toHex2(d.AppVkHash),
			SmtRoot:       toHex2(d.SmtRoot),
		})
	}
	return ret
}

func toHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func toHex2(b [32]byte) string {
	return "0x" + hex.EncodeToString(b[:])
}
