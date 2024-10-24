package myviper

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	myViper := viper.New()
	myViper.SetConfigName("config")
	myViper.SetConfigType("json")
	myViper.AddConfigPath(".")
	err := myViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return myViper
}
