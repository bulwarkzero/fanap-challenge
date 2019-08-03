package env

import (
	"bulwark/fanap-challenge/utils"

	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Init initialize environment values
func Init(cfg string) {
	viper.SetConfigType("json")

	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		cwd := utils.GetCurrentPath()
		defaultConfigName := "config"
	
		viper.SetConfigName(defaultConfigName)
		viper.AddConfigPath(cwd)
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	log.Printf("Using config file %s\n", viper.ConfigFileUsed())
}
