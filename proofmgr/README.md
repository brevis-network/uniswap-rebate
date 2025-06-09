# prover flow
proofmgr handles prove flow. for each ProofInfo, split into multiple proofs if needed, send batch queries to gw and request app proof from app prover. Then submit app proof to gw and finally aggregate
- gw.SendBatchQueriesAsync, resp has batch_id, repeated reqid and nonce
- generate app proof, then gw.SubmitAppCircuitProof (may fail if gw isn't ready). QueryKey can be built by reqid or GetQueryKeysByBatchId
- gw.GetQueryStatus
```
message GetQueryStatusResponse {
    ErrMsg err = 1;
    QueryStatus status = 2;
    string tx_hash = 3;
    string proof = 4;
    ProofData proof_data = 5;
    string circuit_output = 6;
    string proof_with_public_inputs = 7;
}
```
proof table:
reqid, idx, gwResp(batch_id, repeated reqid and nonce), app proof?