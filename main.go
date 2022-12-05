package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// viper.SetConfigFile("./configs/dev.yaml")
	// viper.ReadInConfig()

	viper.SetConfigName("dev")  // config file name without extension
	viper.SetConfigType("yaml") // config file extension
	viper.AddConfigPath("./configs/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("An error has occurred when reading config file: ", err)
	}

	// Different ways to retrieve config values:

	// Using .GetString
	env := viper.GetString("app.env")
	fmt.Println(env)

	// Using Get, then declaring the type
	token := viper.Get("app.token").(string)
	fmt.Println(token)

	secret_data := viper.Get("app.secret_data").(string)
	fmt.Println(secret_data)

}
