// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package dal

import (
	"context"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/webapi"
)

const claimerAdd = `-- name: ClaimerAdd :exec
INSERT INTO claimer (chid, router, evlog) VALUES ($1, $2, $3)
`

type ClaimerAddParams struct {
	Chid   uint64                   `json:"chid"`
	Router string                   `json:"router"`
	Evlog  binding.ClaimHelpClaimer `json:"evlog"`
}

func (q *Queries) ClaimerAdd(ctx context.Context, arg ClaimerAddParams) error {
	_, err := q.db.ExecContext(ctx, claimerAdd, arg.Chid, arg.Router, arg.Evlog)
	return err
}

const claimerGet = `-- name: ClaimerGet :one
SELECT evlog FROM claimer WHERE chid = $1 and router = $2
`

type ClaimerGetParams struct {
	Chid   uint64 `json:"chid"`
	Router string `json:"router"`
}

func (q *Queries) ClaimerGet(ctx context.Context, arg ClaimerGetParams) (binding.ClaimHelpClaimer, error) {
	row := q.db.QueryRowContext(ctx, claimerGet, arg.Chid, arg.Router)
	var evlog binding.ClaimHelpClaimer
	err := row.Scan(&evlog)
	return evlog, err
}

const monGet = `-- name: MonGet :one
SELECT key, blknum, blkidx FROM monitor WHERE key = $1
`

func (q *Queries) MonGet(ctx context.Context, key string) (Monitor, error) {
	row := q.db.QueryRowContext(ctx, monGet, key)
	var i Monitor
	err := row.Scan(&i.Key, &i.Blknum, &i.Blkidx)
	return i, err
}

const monSet = `-- name: MonSet :exec
INSERT INTO monitor (key, blknum, blkidx) VALUES ($1, $2, $3) ON CONFLICT (key) DO UPDATE
SET blknum = excluded.blknum, blkidx = excluded.blkidx
`

type MonSetParams struct {
	Key    string `json:"key"`
	Blknum uint64 `json:"blknum"`
	Blkidx int64  `json:"blkidx"`
}

func (q *Queries) MonSet(ctx context.Context, arg MonSetParams) error {
	_, err := q.db.ExecContext(ctx, monSet, arg.Key, arg.Blknum, arg.Blkidx)
	return err
}

const poolAdd = `-- name: PoolAdd :exec
INSERT INTO pools (chid, poolid, poolkey) VALUES ($1, $2, $3)
`

type PoolAddParams struct {
	Chid    uint64          `json:"chid"`
	Poolid  string          `json:"poolid"`
	Poolkey binding.PoolKey `json:"poolkey"`
}

func (q *Queries) PoolAdd(ctx context.Context, arg PoolAddParams) error {
	_, err := q.db.ExecContext(ctx, poolAdd, arg.Chid, arg.Poolid, arg.Poolkey)
	return err
}

const poolGet = `-- name: PoolGet :one
SELECT poolkey FROM pools WHERE chid = $1 and poolid = $2
`

type PoolGetParams struct {
	Chid   uint64 `json:"chid"`
	Poolid string `json:"poolid"`
}

func (q *Queries) PoolGet(ctx context.Context, arg PoolGetParams) (binding.PoolKey, error) {
	row := q.db.QueryRowContext(ctx, poolGet, arg.Chid, arg.Poolid)
	var poolkey binding.PoolKey
	err := row.Scan(&poolkey)
	return poolkey, err
}

const pools = `-- name: Pools :many
SELECT poolid, poolkey FROM pools WHERE chid = $1
`

type PoolsRow struct {
	Poolid  string          `json:"poolid"`
	Poolkey binding.PoolKey `json:"poolkey"`
}

func (q *Queries) Pools(ctx context.Context, chid uint64) ([]PoolsRow, error) {
	rows, err := q.db.QueryContext(ctx, pools, chid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PoolsRow
	for rows.Next() {
		var i PoolsRow
		if err := rows.Scan(&i.Poolid, &i.Poolkey); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const reqAdd = `-- name: ReqAdd :exec
INSERT INTO reqs (id, proofreq) VALUES ($1, $2)
`

type ReqAddParams struct {
	ID       int64               `json:"id"`
	Proofreq *webapi.NewProofReq `json:"proofreq"`
}

func (q *Queries) ReqAdd(ctx context.Context, arg ReqAddParams) error {
	_, err := q.db.ExecContext(ctx, reqAdd, arg.ID, arg.Proofreq)
	return err
}

const reqGet = `-- name: ReqGet :one
SELECT id, step, proofreq, calldata FROM reqs WHERE id = $1
`

func (q *Queries) ReqGet(ctx context.Context, id int64) (Req, error) {
	row := q.db.QueryRowContext(ctx, reqGet, id)
	var i Req
	err := row.Scan(
		&i.ID,
		&i.Step,
		&i.Proofreq,
		&i.Calldata,
	)
	return i, err
}
