package webapi

import (
	"database/sql/driver"

	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"google.golang.org/protobuf/encoding/protojson"
)

// impl db required interface for objects
func (c *NewProofReq) Value() (driver.Value, error) {
	return protojson.Marshal(c)
}

// tried double pointer but not work
func (u *NewProofReq) Scan(value interface{}) error {
	if u == nil {
		u = new(NewProofReq)
	}
	return protojson.Unmarshal(value.([]byte), u)
}

// impl db required interface for objects
func (c *CallData) Value() (driver.Value, error) {
	return protojson.Marshal(c)
}

// tried double pointer but not work
func (u *CallData) Scan(value interface{}) error {
	if u == nil {
		u = new(CallData)
	}
	return protojson.Unmarshal(value.([]byte), u)
}

func GwRespToCalldata(resp *gwproto.GetQueryStatusResponse) *CallData {
	return &CallData{
		// ProofIds is not needed for single proof.
		Proof: resp.ProofWithPublicInputs,
		ProofDataArray: []*BrevisProofData{
			GwProofData(resp.ProofData),
		},
		AppCircuitOutputs: []string{resp.CircuitOutput},
	}
}

func GwProofData(d *gwproto.ProofData) *BrevisProofData {
	return &BrevisProofData{
		CommitHash:    d.CommitHash,
		AppCommitHash: d.AppCommitHash,
		AppVkHash:     d.AppVkHash,
		SmtRoot:       d.SmtRoot,
		DummyInput:    d.DummyInputCommitment,
	}
}
