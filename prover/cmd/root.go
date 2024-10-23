/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prover",
	Short: "accept prove request, generate app proof, send to brevis gateway",
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "prover.toml", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		log.Fatal("config flag is empty")
	}
	viper.SetConfigFile(cfgFile)
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read config", cfgFile, "err:", err)
	}
}
