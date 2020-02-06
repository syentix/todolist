package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Creds       string
	TelegramAPI string
	config      *configStruct
)

type configStruct struct {
	Creds       string `json:"credentials"`
	TelegramAPI string `json:"telegram-http-api"`
}

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Creds = config.Creds
	TelegramAPI = config.TelegramAPI

	return nil
}
