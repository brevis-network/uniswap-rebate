package dal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	_ "github.com/lib/pq"
)

const (
	SqlUrlFmt = "postgresql://unirebate@%s/unirebate?sslmode=disable"
	TxTimeout = 5 * time.Second
)

var (
	bgCtx = context.Background()
)

type DAL struct {
	*Queries         // so all sqlc generated funcs are pass through
	raw      *sql.DB // underlying from sql.Open, don't use db as Queries also has db field
}

// host:port like localhost:26257
func NewDAL(hostport string) (*DAL, error) {
	sqldb, err := sql.Open("postgres", fmt.Sprintf(SqlUrlFmt, hostport))
	if err != nil {
		return nil, err
	}
	// sqldb.SetMaxOpenConns(20)
	return &DAL{
		Queries: &Queries{sqldb},
		raw:     sqldb,
	}, nil
}

// minimal wrap, caller should do dtx := dal.Queries{tx} or dtx := dal.New(tx) in fn and call all dtx.XXX
// crdb.ExecuteTx handles retry based on fn return error
// fn must take care when wrapping errors returned from the database driver with additional context.
// ie. fn MUST use fmt.Errorf("%w", err) to ensure original db error is wrapped in return and available
// to crdb.ExecuteTx to decide what to do
// fn SHOULD avoid change any state that's not db. if can't avoid, state change must be idempotent,
// because fn may be run multiple times. non-db state change MUST be reverted by fn or upper level code if error
func (d *DAL) DoTx(fn func(tx *sql.Tx) error) error {
	ctx, cancel := context.WithTimeout(bgCtx, TxTimeout)
	defer cancel()
	return crdb.ExecuteTx(ctx, d.raw, nil, fn)
}

func (d *DAL) Close() {
	if d.raw != nil {
		d.raw.Close()
	}
	d.raw = nil
	d.Queries = nil
}

// return epoch millisec
func Nowms() int64 {
	return time.Now().UnixNano() / 1e6
}

// impl mon2.DAL interface
func (d *DAL) GetMonitorBlock(key string) (uint64, int64, bool, error) {
	mon, err := d.MonGet(bgCtx, key)
	// first check if err is no rows
	found, err2 := ChkQueryRow(err)
	return mon.Blknum, mon.Blkidx, found, err2
}

func (d *DAL) SetMonitorBlock(key string, blockNum uint64, blockIdx int64) error {
	return d.MonSet(bgCtx, MonSetParams{
		Key:    key,
		Blknum: blockNum,
		Blkidx: blockIdx,
	})
}

// if err is sql.ErrNoRows, return false, nil
func ChkQueryRow(err error) (bool, error) {
	found := false
	if err == nil {
		found = true
	} else if err == sql.ErrNoRows {
		err = nil
	}
	return found, err
}
