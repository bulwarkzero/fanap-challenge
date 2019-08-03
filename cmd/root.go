package cmd

import (
	"bulwark/fanap-challenge/env"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/config.json)")
}

func initConfig() {
	env.Init(cfgFile)
}

var rootCmd = &cobra.Command{
	Use:   "fanap-challenge",
	Short: "Fanap challenge application",
	Long:  `Fanap challenge`,
	Run: func(cmd *cobra.Command, args []string) {
		serveCmd.Run(cmd, args)
	},
}

// Execute excutes root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
