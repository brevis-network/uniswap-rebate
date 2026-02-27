-- name: MonGet :one
SELECT * FROM monitor WHERE key = $1;

-- name: MonSet :exec
INSERT INTO monitor (key, blknum, blkidx) VALUES ($1, $2, $3) ON CONFLICT (key) DO UPDATE
SET blknum = excluded.blknum, blkidx = excluded.blkidx;

-- name: PoolAdd :exec
INSERT INTO pools (chid, poolid, poolkey) VALUES ($1, $2, $3);

-- name: PoolGet :one
SELECT poolkey FROM pools WHERE chid = $1 and poolid = $2;

-- name: Pools :many
SELECT poolid, poolkey FROM pools WHERE chid = $1;

-- name: ReqAdd :exec
INSERT INTO reqs (id, router, usr_req, proof_info) VALUES ($1, $2, $3, $4);

-- name: ReqGetCalldata :one
SELECT calldata FROM reqs WHERE id = $1;

-- name: ReqSetCalldata :exec
UPDATE reqs SET calldata = $1 WHERE id = $2;

-- name: ClaimerAdd :exec
INSERT INTO claimer (chid, router, beneficiary) VALUES ($1, $2, $3)
ON CONFLICT (chid, router) DO UPDATE
SET beneficiary = excluded.beneficiary;

-- name: ClaimerGet :one
SELECT beneficiary, fetch_blk FROM claimer WHERE chid = $1 and router = $2;

-- name: Claimers :many
SELECT chid, router, beneficiary, fetch_blk FROM claimer ORDER BY chid, router;

-- name: ClaimerSetFetchBlk :exec
UPDATE claimer
SET fetch_blk = GREATEST(fetch_blk, $1)
WHERE chid = $2 AND router = $3;

-- name: ProofAdd :exec
INSERT INTO proof (reqid, idx, app_prover, app_proof_id, app_circuit_info) VALUES ($1, $2, $3, $4, $5);

-- name: ProofSetGwInfo :exec
UPDATE proof
SET gateway_batch_id = $1,
    gateway_request_id = $2,
    gateway_nonce = $3
WHERE reqid = $4 AND idx = $5;

-- name: ProofSetAppProof :exec
UPDATE proof SET app_proof = $1, app_circuit_info = $2 WHERE app_proof_id = $3;

-- name: ProofSetGwResp :exec
UPDATE proof SET gateway_query_status = $1 WHERE gateway_request_id = $2 AND gateway_nonce = $3;

-- name: ProofGetIds :many
SELECT idx, app_prover, app_proof_id, gateway_batch_id, gateway_request_id, gateway_nonce FROM proof WHERE reqid = $1 ORDER BY idx;
