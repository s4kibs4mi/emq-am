package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
	"os"
)

var RootCmd = cobra.Command{}

func Execute() {
	viper.AddConfigPath("etc")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	readErr := viper.ReadInConfig()
	if readErr != nil {
		fmt.Println("Couldn't read config")
		os.Exit(-1)
	}

	RootCmd.AddCommand(&ServeCmd)
	RootCmd.Execute()
}
