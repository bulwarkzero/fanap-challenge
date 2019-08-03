package cmd

import (
	"bulwark/fanap-challenge/database/migrations"
	"bulwark/fanap-challenge/pkg/db"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrateCmd command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Long:  `Migrate all changes to database`,
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Init(db.Config{
			Engine:   viper.GetString("db.engine"),
			DbName:   viper.GetString("db.db_name"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetInt("db.port"),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			Log:      viper.GetBool("db.log"),
			SslMode:  viper.GetString("db.ssl_mode"),
		})

		migrations.Migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().String("db_address", "127.0.0.1", "Database full address")
	migrateCmd.Flags().Int("db_port", 5432, "Database instance port")
	migrateCmd.Flags().String("db_username", "postgres", "Database username")
	migrateCmd.Flags().String("db_password", "", "Database password")
	migrateCmd.Flags().String("db_dbname", "fanap-challenge", "Default database name")

	viper.BindPFlag("db.host", migrateCmd.Flags().Lookup("db_address"))
	viper.BindPFlag("db.port", migrateCmd.Flags().Lookup("db_port"))
	viper.BindPFlag("db.username", migrateCmd.Flags().Lookup("db_username"))
	viper.BindPFlag("db.password", migrateCmd.Flags().Lookup("db_password"))
	viper.BindPFlag("db.db_name", migrateCmd.Flags().Lookup("db_dbname"))
}
