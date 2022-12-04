package main

import (
    "fmt"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./configs/dev.yaml")
    viper.ReadInConfig()

    // add env variables as needed
    val1 := viper.Get("KEY1").(string)
    val2 := viper.Get("KEY2").(string)
	val3 := viper.Get("KEY3").(string)


    fmt.Println(val1, val2, val3)
}