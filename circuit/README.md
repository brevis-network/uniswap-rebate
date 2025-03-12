# circuit
Per Uniswap 2025-03-10 decision, rebate gas for one tx is `n * (rebatePerSwap+rebatePerHook) + rebateFixed` where n is number of valid swaps in this tx. We then compare the result with tx's actual gas usage * 0.8  and choose the smaller one.

This requires circuit to be aware of which receipts belong to the same tx, so order receipts by poolid no longer works.

- receipts are from the same tx if both Receipt.MptKeyPath and BlockNum equal
- tx gas usage isn't available from sdk yet, so it's circuit struct field `TxGasCap` for now

TxGasCap must be of same length as receipts, and if one tx has multiple eligible swaps, it's set as follows:
|TxGasCap| 0 | Tx1 | Tx2 |
| --- | --- | --- | --- |
|Receipts| t1r1 | t1r2 | t2r1 |

If a tx has k receipts, first k-1 are 0 and last one is actual tx gas cap.

This design can be easily migrated to gas within tx once sdk adds support.