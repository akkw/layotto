package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mosn.io/layotto/spec/proto/runtime/v1"
	"os"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(runtime.RuntimeClientCommand)
}

var RootCmd = &cobra.Command{
	Use:   "Layotto",
	Short: "123",
	Long:  ``,
}

func main() {
	Execute()
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
