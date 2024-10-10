package binding

// to be saved in db. fields are args for ClaimWithZkProofs
type CallData struct {
	ProofIds          [][32]byte
	Proof             []byte
	ProofDataArray    []BrevisProofData
	AppCircuitOutputs [][]byte
}
