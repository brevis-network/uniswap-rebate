package main

import (
	"context"
	"fmt"
	"os"

	"github.com/brevis-network/uniswap-rebate/binding"
	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	dbAddr  string
	rpcURL  string
	txHash  string
	poolMgr string
	chId    uint64
)

var rootCmd = &cobra.Command{
	Use:   "poolinitdb",
	Short: "load a Uniswap v4 Pool Initialize event from a transaction receipt into the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := loadConfig(); err != nil {
			return err
		}
		dbAddr = viper.GetString("db")
		cfgs := onchain.GetMcc("multichain")
		for _, cfg := range cfgs {
			if cfg.ChainID == chId {
				rpcURL = cfg.Gateway
				poolMgr = cfg.PoolMgr
				break
			}
		}

		ctx := context.Background()
		ec, err := ethclient.DialContext(ctx, rpcURL)
		if err != nil {
			return fmt.Errorf("dial rpc: %w", err)
		}
		defer ec.Close()

		chid, err := ec.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("get chain id: %w", err)
		}
		poolMgrAddr := common.HexToAddress(poolMgr)
		filter, err := binding.NewPoolMgrFilterer(poolMgrAddr, ec)
		if err != nil {
			return fmt.Errorf("create pool manager filterer: %w", err)
		}

		receipt, err := ec.TransactionReceipt(ctx, onchain.Hex2hash(txHash))
		if err != nil {
			return fmt.Errorf("fetch receipt: %w", err)
		}

		lg, err := onchain.FindPoolInitLog(receipt, &poolMgrAddr, nil)
		if err != nil {
			return err
		}

		db, err := dal.NewDAL(dbAddr)
		if err != nil {
			return fmt.Errorf("connect db: %w", err)
		}
		defer db.Close()

		rec, err := onchain.AddPoolInitFromLog(ctx, db, chid.Uint64(), filter, *lg)
		if err != nil {
			return err
		}

		fmt.Printf("added pool chid=%d poolid=%s tx=%s log_index=%d hooks=%s\n",
			chid.Uint64(), rec.PoolID, txHash, lg.Index, rec.PoolKey.Hooks.Hex())
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&cfgFile, "config", "config.toml", "config file used for db and multichain pool manager defaults")
	rootCmd.Flags().StringVar(&txHash, "tx", "", "transaction hash containing the Initialize event")
	rootCmd.Flags().Uint64Var(&chId, "chid", 1, "Chain ID of the chain")
}

func loadConfig() error {
	if cfgFile == "" {
		return nil
	}
	if _, err := os.Stat(cfgFile); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	viper.SetConfigFile(cfgFile)
	return viper.ReadInConfig()
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
