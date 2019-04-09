package config

import (
	"fmt"
	"io/ioutil"
	"bytes"
    "github.com/spf13/viper"
)


func InitConfig(path string) *viper.Viper{
    if  path == ""{
        panic("Invalid config file path")
    }   
    bArray, err := ioutil.ReadFile(path)
    if err!= nil{
        panic(fmt.Errorf("Error parsing config file %s", err))
    }
    return parseConfig(bArray)
}

func parseConfig(byteArray []byte) *viper.Viper{
    var configViper = viper.New()
    configViper.SetConfigType("yaml")

    err := configViper.ReadConfig(bytes.NewBuffer(byteArray))
    if err != nil{
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    return configViper
}