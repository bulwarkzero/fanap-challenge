package cmd

import (
	"bulwark/fanap-challenge/database"
	"bulwark/fanap-challenge/pkg/db"
	httpServer "bulwark/fanap-challenge/server/http"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	serveCmd.AddCommand(serveHTTPCmd)
	serveHTTPCmd.Flags().StringP("address", "a", "localhost", "HTTP Server address")
	serveHTTPCmd.Flags().IntP("port", "p", 8080, "HTTP Server port")
	// serveHTTPCmd.Flags().BoolP("debug", "d", false, "Debug mode")

	viper.BindPFlag("http.address", serveHTTPCmd.Flags().Lookup("address"))
	viper.BindPFlag("http.port", serveHTTPCmd.Flags().Lookup("port"))
	// viper.BindPFlag("debug", serveHTTPCmd.Flags().Lookup("debug"))
}

var serveHTTPCmd = &cobra.Command{
	Use:   "http",
	Short: "Start http server",
	Long:  `Start http server`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Init(db.Config{
			Engine:   viper.GetString("db.engine"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetInt("db.port"),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			DbName:   viper.GetString("db.db_name"),
			SslMode:  viper.GetString("db.ssl_mode"),
			Log:      viper.GetBool("db.debug"),
		})

		if err != nil {
			panic(fmt.Sprintf("Can't open database: %v", err))
		}

		httpServer.Serve(httpServer.Config{
			Address: viper.GetString("http.address"),
			Port:    viper.GetInt("http.port"),
			Debug:   viper.GetBool("debug"),
		})
	},
}
