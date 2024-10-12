package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/brevis-network/uniswap-rebate/webapi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	fcfg = flag.String("c", "config.toml", "config toml file")
	fdir = flag.String("d", "data", "directory path for data files")
	// chainid -> onechain
	chainMap map[uint64]*onchain.OneChain
)

func main() {
	flag.Parse()
	viper.SetConfigFile(*fcfg)
	err := viper.ReadInConfig()
	chkErr(err, "viper ReadInConfig")

	// check fdir is writable
	f, err := os.CreateTemp(*fdir, "tmp")
	chkErr(err, "fail to write to dir: "+*fdir)
	f.Close()
	os.Remove(f.Name())

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
	}

	// start grpc server, only listens on localhost
	grpcEndpoint := fmt.Sprintf("localhost:%d", viper.GetInt("grpcport"))
	lis, err := net.Listen("tcp", grpcEndpoint)
	chkErr(err, "listen")

	svr := &Server{
		db: db,
	}
	gs := grpc.NewServer()
	webapi.RegisterUniRebateServer(gs, svr)
	go func() {
		if err = gs.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// grcp-gateway for http apis
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = webapi.RegisterUniRebateHandlerFromEndpoint(context.Background(), mux, grpcEndpoint, opts)
	chkErr(err, "gw register")
	// blocking and proxy http to grpc server
	http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("httpport")), mux)
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
