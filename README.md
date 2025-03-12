# uniswap-rebate
This is NOT a uniswap v4 hook, but to give fee rebate to pools w/ non-zero hooks address.

See https://github.com/uniswapfoundation/router-rebates/ for more details. Onchain claim code is integrated into Uniswap's contract.

Over the time, scope/requirements have changed multiple times. Latest as of 2025-03-10:
1. swaps may happen on 12 univ4 chains. claim only happens on unichain
2. rebate gas cost in Eth
3. each router impl `rebateClaimer() view` to indicate which address is allowed to call claim onchain

## High level design
1. Due to uniswap contract doesn't save PoolKey (a struct about the pool including hooks address), we need to filter non-eligible pools ourselves. Onchain contract keeps a map of poolid->bool, there is a public api addPool(PoolKey) can be called by anyone, it checks poolkey.hooks isn’t all 0 and saves key.toId() to map. Compare to check in zk circuit, this has higher onchain cost per pool but minimal eng complexity (as keccak to compute poolid is not easy using sdk)
2. rebate contract maintains addr->poolid->blocknum map to avoid replay
3. Circuit output: addr,[poolid,firstblk,lastblk,u128(uni)] contract will check poolid is in eligible map and firstblk > saved blocknum, sum uni across all pools, then send total uni to specified recipient.

## Flow
1. projects submit list of tx hashes to brevis unirebate server and get a request id
2. server will fetch tx receipts, organize events by poolids and start proving process (ineligible pools are ignored)
3. server submits to Brevis gateway, polling till onchain calldata available and save to db
4. projects polling server for calldata and send onchain tx to their own router contract

## Sepolia setup
- this repo, deploy uni/eth price oracle, UniEthTestnet has ratio in constructor
- brevis-network/uniswap-v4-core repo mysetup branch, deploy myswaprouter and two erc20s. Note smaller addr token will be currency0.
- brevis-network/uniswap-v4-template repo my branch, set script/base/Constants.sol POOLMANAGER and posm from [uniswap doc](https://docs.uniswap.org/contracts/v4/deployments), Config.sol token0 and token1 from previous step
- run 00_Counter to deploy counter hook by CREATE2, then add hook address to Config.sol. CREATE2 due to uniswap requires hook addr to have certain bits set, only possible by mining salt with create2.
- run 01_CreatePoolAndMintLiquidity, initialize pool w/ our erc20 tokens and add liquidity
- 03_Swap: set myswaprouter address and modify swap amount, then run script

now we have swap event and slot available at same block. next to setup zk related, prove and verify onchain

onchain:
- this repo, deploy zkrebate and transfer test uni to zkrebate contract
- run add_pool, make sure PoolKey matches 01_CreatePoolAndMintLiquidity, query map to check poolid matches
- add zkrebate addr to allowed provers. call BrevisProof.addProvers (can use etherscan web, must be owner)

offchain:
- server must listen to pool Initialize event from 01_CreatePoolAndMintLiquidity blocknumber and populate db
- send swap tx to server which starts the process and sent to prover
- prover generates app proof and submits to Brevis Gateway

when proof is ready, run claim.sol submit onchain to myswaprouter:

| Contract  | Address |
| ------------- | ------------- |
| Oracle | 0xc8caf3c1d76d264b6d5d511d883b58b3874eafa0 |
| MySwapRouter | 0xd45f72e31e19fcf75bF47d217e500745Cb36263b |
| Hook | 0xaA5f84D2F6A2423E50b3e598F934f6626e12CAc0 |
| ZkRebate | 0xA50e5203ECa1685B4e01A8A569dd12150A8b419D |
