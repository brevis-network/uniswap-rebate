/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/celer-network/goutils/log"
	"github.com/spf13/cobra"
)

var (
	outDir, srsDir string
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "sdk.Compile circuit",
	Run: func(cmd *cobra.Command, args []string) {
		app, _ := sdk.NewBrevisApp(11155111, "https://sepolia.drpc.org", outDir)
		appCircuit := circuit.DefaultCircuit()

		_, _, _, vkhash, err := sdk.Compile(appCircuit, outDir, srsDir, app)
		if err != nil {
			log.Error("sdk.Compile err:", err)
		}
		log.Infof("vkhash: %x", vkhash)
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
	compileCmd.PersistentFlags().StringVar(&outDir, "outDir", "$HOME/circuitOut/unigasrebate", "folder for circuit compile output")
	compileCmd.PersistentFlags().StringVar(&srsDir, "srsDir", "$HOME/kzgsrs", "folder for kzg srs")
}
