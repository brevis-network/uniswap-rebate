/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/uniswap-rebate/onchain"
	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/constraint"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
)

// flags
var (
	outDir, brvGw string
	// accept req on this port
	lport int
)
var (
	compiledCircuit constraint.ConstraintSystem
	pk              plonk.ProvingKey
	vk              plonk.VerifyingKey
	vkHashStr       string // will set after read vk
)

// proveCmd represents the prove command
var proveCmd = &cobra.Command{
	Use:   "prove",
	Short: "accept prove req, send to Brevis",

	Run: func(cmd *cobra.Command, args []string) {
		// takes minutes to load
		var err error
		compiledCircuit, pk, vk, err = sdk.ReadSetupFrom(outDir)
		chkErr(err, "ReadSetupFrom "+outDir)
		vkhash, err := sdk.ComputeVkHash(vk)
		chkErr(err, "ComputeVkHash")
		vkHashStr = fmt.Sprintf("%s", vkhash) // common.Hash.Format adds 0x prefix
		// setup server
		router := httprouter.New()
		router.POST("/prove", HandleProve)
		log.Infoln("listen on port:", lport)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", lport), router))
	},
}

// handle StartProve req
func HandleProve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req []*onchain.OneProveReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	log.Infoln("provereqs:", len(req))
	// go ProveReq(&req)
	fmt.Fprint(w, "OK\n")
}

func init() {
	rootCmd.AddCommand(proveCmd)
	proveCmd.PersistentFlags().StringVar(&outDir, "outDir", "$HOME/circuitOut/viphookalg", "folder for circuit in/output")
	proveCmd.PersistentFlags().StringVar(&brvGw, "brvgw", "", "brevis gateway grpc endpoint")
	proveCmd.PersistentFlags().IntVar(&lport, "port", 8889, "listen port for prove request")
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
