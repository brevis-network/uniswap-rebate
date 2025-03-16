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
INSERT INTO reqs (id, proofreq) VALUES ($1, $2);

-- name: ReqGet :one
SELECT * FROM reqs WHERE id = $1;