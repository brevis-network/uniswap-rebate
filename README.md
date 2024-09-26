# uniswap-rebate
This is NOT a uniswap v4 hook, but to give fee rebate to pools w/ non-zero hooks address.

## Use case
Other projects deploy router contracts and interact with uniswap poolmanager. Uniswap will rebate gas fees in UNI if swap pool has hooks. Projects need to initiate claim flow and it's up to them how to distribute to actual users. Uniswap may have an internal implementation based on centralized signature. Brevis will provide a zk-based solution

## High level design
1. Due to uniswap contract doesn't save PoolKey (a struct about the pool including hooks address), we need to filter non-eligible pools ourselves. Onchain contract keeps a map of poolid->bool, there is a public api addPool(PoolKey) can be called by anyone, it checks poolkey.hooks isn’t all 0 and saves key.toId() to map. Compare to check in zk circuit, this has higher onchain cost per pool and minimal eng effort (as keccak to compute poolid is not easy using sdk)
2. Onchain also maintains addr->poolid->blocknum map to avoid replay, and a map addr->claimable eth
3. Circuit output: addr,[poolid,firstblk,lastblk,u128(eth)] contract will check poolid is in eligible map and firstblk > saved blocknum, then add eth to claimable map
