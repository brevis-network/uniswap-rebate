# uniswap-rebate
This is NOT a uniswap v4 hook, but to give fee rebate to pools w/ non-zero hooks address.

## v2 supports multiple pools in one app proof and use single proof onchain

## Use case
Other projects deploy router contracts and interact with uniswap poolmanager. Uniswap will rebate gas fees in UNI if swap pool has hooks. Projects need to initiate claim flow and it's up to them how to distribute to actual users. Uniswap may have an internal implementation based on centralized signature. Brevis will provide a zk-based solution

## High level design
1. Due to uniswap contract doesn't save PoolKey (a struct about the pool including hooks address), we need to filter non-eligible pools ourselves. Onchain contract keeps a map of poolid->bool, there is a public api addPool(PoolKey) can be called by anyone, it checks poolkey.hooks isn’t all 0 and saves key.toId() to map. Compare to check in zk circuit, this has higher onchain cost per pool and minimal eng effort (as keccak to compute poolid is not easy using sdk)
2. Onchain also maintains addr->poolid->blocknum map to avoid replay, and a map addr->claimable eth
3. Circuit output: addr,[poolid,firstblk,lastblk,u128(eth)] contract will check poolid is in eligible map and firstblk > saved blocknum

## Flow
1. projects submit list of tx hashes to brevis unirebate server and get a request id
2. server will fetch tx receipts, organize events by poolids and start proving process (possible filtering of ineligible pools here)
3. server submits to Brevis gateway and needs a way to get ready to submit onchain data and save to db
4. projects polling data and send onchain tx

## Sepolia complete flow
- this repo, deploy uni/eth price oracle, test contract has ratio in constructor
- brevis-network/uniswap-v4-core repo mysetup branch, deploy myswaprouter and two erc20s. Note smaller addr token will be currency0.
- brevis-network/uniswap-v4-template repo my branch, set script/base/Constants.sol POOLMANAGER and posm from [uniswap doc](https://docs.uniswap.org/contracts/v4/deployments), Config.sol token0 and token1 from previous step
- run 00_Counter to deploy counter hook by CREATE2, then add hook address to Config.sol. CREATE2 due to uniswap requires hook addr to have certain bits set, only possible by mining salt with create2.
- run 01_CreatePoolAndMintLiquidity
- set myswaprouter address and modify swap amount in 03_Swap and run script

now we have swap event and slot at same block. next to prove/setup zk related

onchain:
- this repo, deploy zkrebate and transfer test uni to zkrebate contract
- run add_pool, make sure PoolKey matches 01_CreatePoolAndMintLiquidity, can query map to check poolid matches
- add zkrebate to allowed provers. call BrevisProof.addProvers (can use etherscan web, must be owner)

offchain:
- server must listen to pool Initialize event from 01_CreatePoolAndMintLiquidity blocknumber and populate db
- send swap tx to server which starts the process and sent to prover
- prover generates app proof and submits to Brevis Gateway

when proof is ready, submit onchain to myswaprouter:
- run claim.sol

| Contract  | Address |
| ------------- | ------------- |
| Oracle | 0xc8caf3c1d76d264b6d5d511d883b58b3874eafa0 |
| MySwapRouter | 0xd45f72e31e19fcf75bF47d217e500745Cb36263b |
| Hook | 0xaA5f84D2F6A2423E50b3e598F934f6626e12CAc0 |
| ZkRebate | 0xA50e5203ECa1685B4e01A8A569dd12150A8b419D |
