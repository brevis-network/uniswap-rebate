package onchain

import "github.com/spf13/viper"

// config related defs and func
func GetMcc(key string) MultiChainConfig {
	var m MultiChainConfig
	viper.UnmarshalKey(key, &m)
	return m
}

type MultiChainConfig []*OneChainConfig

type OneChainConfig struct {
	ChainID       uint64
	Name, Gateway string
	// blk related for monitor
	BlkInterval, BlkDelay        uint64
	MaxBlkDelta, ForwardBlkDelay uint64

	GasPerSwap, GasPerTx uint32
	// uni v4 poolmgr
	PoolMgr string
	// vk hash
	VkHash string
}
