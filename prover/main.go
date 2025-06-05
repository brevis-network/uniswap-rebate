package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/sdk/prover"
	"github.com/brevis-network/uniswap-rebate/circuit"
	"github.com/celer-network/goutils/log"
	"github.com/spf13/viper"
)

var (
	flagConfig = flag.String("c", "", "config toml file")
)

func getServiceConfig(key string) (*prover.ServiceConfig, error) {
	var s prover.ServiceConfig
	err := viper.UnmarshalKey(key, &s)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalKey err: %w", err)
	}
	return &s, nil
}

func getSourceChainConfigs(key string) (prover.SourceChainConfigs, error) {
	var s prover.SourceChainConfigs
	err := viper.UnmarshalKey(key, &s)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalKey err: %w", err)
	}
	return s, nil
}

func main() {
	flag.Parse()
	viper.SetConfigFile(".env") // get circuit name, and can add special config in future
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("ReadInConfig err: %s", err.Error())
		os.Exit(1)
	}
	circuitName := viper.GetString("circuit_name")
	log.Infof("this prover circuit name: %s", circuitName)

	viper.SetConfigFile(*flagConfig) // common config such as chain config and rpc
	err = viper.ReadInConfig()
	if err != nil {
		log.Errorf("ReadInConfig err: %s", err.Error())
		os.Exit(1)
	}
	grpcPort := viper.GetUint("grpc_port")
	restPort := viper.GetUint("rest_port")

	serviceConfig, err := getServiceConfig("service")
	if err != nil {
		log.Errorf("getServiceConfig err: %s", err.Error())
		os.Exit(1)
	}

	sourceChainConfigs, err := getSourceChainConfigs("source_chain")
	if err != nil {
		log.Errorf("getSourceChainConfigs err: %s", err.Error())
		os.Exit(1)
	}

	log.Infof("sourceChainConfigs: %+v", sourceChainConfigs)

	err = startService(*serviceConfig, sourceChainConfigs, circuitName, grpcPort, restPort)
	if err != nil {
		log.Errorf("startService err: %s", err)
		os.Exit(1)
	}
}

func startService(serviceConfig prover.ServiceConfig, sourceChainConfigs prover.SourceChainConfigs, circuitName string, grpcPort, restPort uint) error {
	log.Infof("start service for circuit: %s", circuitName)
	var appCircuit sdk.AppCircuit
	switch circuitName {
	case "unigas":
		appCircuit = &circuit.GasCircuit{}
	default:
		return fmt.Errorf("unsupported circuit %s", circuitName)
	}

	serviceConfig.SetupDir = GetCircuitOutDir(serviceConfig.SetupDir, circuitName)
	log.Infof("serviceConfig: %+v", serviceConfig)
	proverService, err := prover.NewService(appCircuit, serviceConfig, sourceChainConfigs)
	if err != nil {
		return fmt.Errorf("NewService err: %w", err)
	}

	err = proverService.Serve("", grpcPort, restPort)
	if err != nil {
		//lint:ignore ST1005 Nit
		return fmt.Errorf("Serve err: %w", err)
	}
	return nil
}

func GetCircuitOutDir(outDir string, circuitName string) string {
	return fmt.Sprint(outDir, "/", circuitName)
}
