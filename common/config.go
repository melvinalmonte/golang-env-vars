package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Function to read an environment or return a default value
func GetConfigValue(key string, defaultValue string) string {

	viper.AddConfigPath("./configs/")
	viper.SetConfigName("dev")  // config file name without extension
	viper.SetConfigType("yaml") // config file extension

	//read configs and handle errors
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("An error has occurred when reading config file: ", err)
	}

	value := viper.Get(key).(string)
	if value != "" {
		return value
	}
	return defaultValue
}
