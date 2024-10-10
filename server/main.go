package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/brevis-network/uniswap-rebate/webapi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	fcfg = flag.String("c", "config.toml", "config toml file")
)

func main() {
	flag.Parse()
	viper.SetConfigFile(*fcfg)
	err := viper.ReadInConfig()
	chkErr(err, "viper ReadInConfig")

	// start grpc server, only listens on localhost
	grpcEndpoint := fmt.Sprintf("localhost:%d", viper.GetInt("grpcport"))
	lis, err := net.Listen("tcp", grpcEndpoint)
	chkErr(err, "listen")

	svr := new(Server)
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
