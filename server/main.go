package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/brevis-network/uniswap-rebate/webapi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
)

var (
	fcfg = flag.String("c", "config.toml", "config toml file")
	// chainid -> onechain
	chainMap map[uint64]*onchain.OneChain
)

func main() {
	flag.Parse()
	viper.SetConfigFile(*fcfg)
	err := viper.ReadInConfig()
	chkErr(err, "viper ReadInConfig")

	// setup db
	db, err := dal.NewDAL(viper.GetString("db"))
	chkErr(err, "new dal")

	chainMap = make(map[uint64]*onchain.OneChain)
	cfgs := onchain.GetMcc("multichain")
	for _, cfg := range cfgs {
		onec, err := onchain.NewOneChain(cfg, db)
		chkErr(err, "NewOneChain"+cfg.Name)
		chainMap[cfg.ChainID] = onec
		onec.MonPoolInit()
		onec.MonClaimer()
	}

	// grcp-gateway for http apis
	mux := runtime.NewServeMux()
	err = webapi.RegisterUniRebateHandlerServer(context.Background(), mux, &Server{
		db: db,
	})
	chkErr(err, "gw register")
	// blocking and proxy http to grpc server
	http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("httpport")), mux)
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
