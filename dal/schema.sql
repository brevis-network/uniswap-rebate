CREATE DATABASE unirebate;
CREATE USER unirebate;
GRANT ALL ON DATABASE unirebate TO unirebate;
SET DATABASE TO unirebate;

-- poolid, poolkey non-mutable once add. only add if poolkey.hooks isn't zero
CREATE TABLE IF NOT EXISTS pools (
    chid BIGINT NOT NULL, -- chainid
    poolid TEXT NOT NULL,  -- 0x...
    poolkey JSONB NOT NULL, -- PoolKey struct that initialized this pool. poolid is hash of PoolKey
    UNIQUE (chid, poolid) -- ensure no duplicated pool on same chain, audo-create index
);

CREATE TABLE IF NOT EXISTS reqs (
    id BIGINT PRIMARY KEY, -- epoch seconds when requested
    step INT NOT NULL DEFAULT 0, -- 0: fetching tx receipts done 1: started app circuit proof 2: has app proof, sent to Brevis gw for final proof 3: have data ready to submit
    proofreq JSONB NOT NULL, -- webapi.NewProofReq
    calldata JSONB -- only not null if step is 3 (received ready to send onchain data from Brevis gw)
);

-- persist block num/index to resume when restart, key is chid-addr
CREATE TABLE IF NOT EXISTS monitor (
    key TEXT PRIMARY KEY NOT NULL,
    blknum BIGINT NOT NULL,
    blkidx INT NOT NULL -- could be -1 when fast forward with no log received
);