package onchain

import "github.com/spf13/viper"

const (
	// ecdsa keystore json, needed to send onchain tx eg. deploy contract
	kKsPath = "kspath"
	kKsPwd  = "kspwd"
)

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

	// uni v4 poolmgr
	PoolMgr string
	// vk hash
	VkHash string
}
