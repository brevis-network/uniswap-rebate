-- name: MonGet :one
SELECT * FROM monitor WHERE key = $1;

-- name: MonSet :exec
INSERT INTO monitor (key, blknum, blkidx) VALUES ($1, $2, $3) ON CONFLICT (key) DO UPDATE
SET blknum = excluded.blknum, blkidx = excluded.blkidx;

-- name: PoolAdd :exec
INSERT INTO pools (poolid, poolkey) VALUES ($1, $2);

-- name: PoolGet :one
SELECT poolkey FROM pools WHERE poolid = $1;

-- name: ReqAdd :exec
INSERT INTO reqs (id, txs) VALUES ($1, $2);