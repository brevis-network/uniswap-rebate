package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/brevis-network/uniswap-rebate/proofmgr"
	"github.com/spf13/viper"
)

var (
	fcfg = flag.String("c", "config.toml", "config toml file")
	// chainid -> onechain
	chainMap map[uint64]*onchain.OneChain
	prMgr    *proofmgr.ProofMgr
)

func main() {
	flag.Parse()
	viper.SetConfigFile(*fcfg)
	err := viper.ReadInConfig()
	chkErr(err, "viper ReadInConfig")

	// setup db
	db, err := dal.NewDAL(viper.GetString("db"))
	chkErr(err, "new dal")

	prMgr = proofmgr.NewProofMgr(viper.GetString("brvgw"), viper.GetStringSlice("prover"), viper.GetUint64("dstchid"), db)

	chainMap = make(map[uint64]*onchain.OneChain)
	cfgs := onchain.GetMcc("multichain")
	for _, cfg := range cfgs {
		onec, err := onchain.NewOneChain(cfg, db)
		chkErr(err, "NewOneChain"+cfg.Name)
		chainMap[cfg.ChainID] = onec
		onec.MonPoolInit()
	}

	registryAddr := viper.GetString("beneficiary_registry")
	if registryAddr == "" {
		log.Fatalln("missing beneficiary_registry config")
	}
	dstChid := viper.GetUint64("dstchid")
	dstChain, ok := chainMap[dstChid]
	if !ok {
		log.Fatalf("dst chain %d must be configured in multichain for registry monitor", dstChid)
	}
	dstChain.MonBeneficiarySet(onchain.Hex2addr(registryAddr))

	srv := &Server{db: db}
	intv := viper.GetDuration("prove_interval")
	if intv <= 0 {
		intv = 5 * time.Minute
	}
	srv.RunScheduledProver(context.Background(), intv)
	/*
		// grcp-gateway for http apis
		mux := runtime.NewServeMux()
		err = webapi.RegisterUniRebateHandlerServer(context.Background(), mux, srv)
		chkErr(err, "gw register")
		// blocking and proxy http to grpc server
		http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("httpport")), mux)
	*/
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
