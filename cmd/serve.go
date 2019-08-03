package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve application",
	Long:  `Serve application`,
	Run: func(cmd *cobra.Command, args []string) {
		serveHTTPCmd.Run(cmd, args)
	},
}
