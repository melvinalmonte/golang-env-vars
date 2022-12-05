package main

import (
	"fmt"
	config "golang-configs/common"
)

func main() {

	env := config.GetConfigValue("app.env", "")
	fmt.Println(env)

	// Using Get, then declaring the type
	token := config.GetConfigValue("app.token", "")
	fmt.Println(token)

	secret_data := config.GetConfigValue("app.secret_data", "")
	fmt.Println(secret_data)

}
