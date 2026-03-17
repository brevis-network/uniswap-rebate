package dal

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/ethereum/go-ethereum/common"
)

const testDBTimeout = 2 * time.Second

func TestNewDALGetPoolKeys(t *testing.T) {
	d, err := NewDAL("localhost:26257")
	if err != nil {
		t.Fatalf("NewDAL err: %v", err)
	}
	defer d.Close()

	ctx, cancel := context.WithTimeout(context.Background(), testDBTimeout)
	defer cancel()
	if err := d.raw.PingContext(ctx); err != nil {
		t.Skipf("local cockroachdb unavailable at localhost:26257: %v", err)
	}

	pid := common.HexToHash("0x79e80f3ea7c2e18eb952a055f34ee2a15313adeee68804c64cdf66bfc0f62405")

	ctx, cancel = context.WithTimeout(context.Background(), testDBTimeout)
	defer cancel()
	want, err := d.PoolGet(ctx, PoolGetParams{
		Chid:   11155111,
		Poolid: pid.Hex(),
	})
	if err == sql.ErrNoRows {
		t.Fatalf("pool %s not found for chain 11155111", pid.Hex())
	}
	if err != nil {
		t.Fatalf("PoolGet err: %v", err)
	}

	got := d.GetPoolKeys(11155111, binding.PoolIdMap{
		pid: true,
	})
	if len(got) != 1 {
		t.Fatalf("GetPoolKeys returned %d rows, want 1", len(got))
	}

	assertPoolKeyEqual(t, got[0], want)
}

func assertPoolKeyEqual(t *testing.T, got, want binding.PoolKey) {
	t.Helper()

	if got.Currency0 != want.Currency0 {
		t.Fatalf("Currency0 = %s, want %s", got.Currency0.Hex(), want.Currency0.Hex())
	}
	if got.Currency1 != want.Currency1 {
		t.Fatalf("Currency1 = %s, want %s", got.Currency1.Hex(), want.Currency1.Hex())
	}
	if got.Fee == nil || want.Fee == nil || got.Fee.Cmp(want.Fee) != 0 {
		t.Fatalf("Fee = %v, want %v", got.Fee, want.Fee)
	}
	if got.TickSpacing == nil || want.TickSpacing == nil || got.TickSpacing.Cmp(want.TickSpacing) != 0 {
		t.Fatalf("TickSpacing = %v, want %v", got.TickSpacing, want.TickSpacing)
	}
	if got.Hooks != want.Hooks {
		t.Fatalf("Hooks = %s, want %s", got.Hooks.Hex(), want.Hooks.Hex())
	}
}
