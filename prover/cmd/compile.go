/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"context"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	outDir, srsDir string
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "sdk.Compile circuit",
	Run: func(cmd *cobra.Command, args []string) {
		var gc gwproto.GatewayClient
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		conn, err := grpc.NewClient(viper.GetString("service.gateway_url"), opts...)
		if err != nil {
			log.Fatal(err)
		}
		gc = gwproto.NewGatewayClient(conn)
		_, err = gc.GetCircuitDigest(context.Background(), &gwproto.CircuitDigestRequest{})
		if err != nil {
			log.Fatal(err)
		}
		appCircuit := circuit.DefaultCircuit()
		_, _, vk, vkhash, err := sdk.Compile(appCircuit, outDir, srsDir, sdk.NewBrevisAppWithDigestsSetOnlyFromRemote())
		if err != nil {
			log.Error("sdk.Compile err:", err)
		}
		log.Infof("vkhash: %x", vkhash)
		res, err := gc.SubmitVK(context.Background(), &gwproto.SubmitVKRequest{
			VkHash:        hexutil.Encode(vkhash),
			VkRaw:         hexutil.Encode(sdk.MustWriteToBytes(vk)),
			ChainId:       1,
			TargetChainId: 130,
			ApiKey:        BrvApiKey,
		})
		if err != nil {
			log.Fatal(err)
		}
		if res.Err != nil {
			panic(res.Err)
		}
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
	compileCmd.PersistentFlags().StringVar(&outDir, "outDir", "$HOME/circuitOut/unigasrebate", "folder for circuit compile output")
	compileCmd.PersistentFlags().StringVar(&srsDir, "srsDir", "$HOME/kzgsrs", "folder for kzg srs")
}
