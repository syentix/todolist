package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Creds  string
	config *configStruct
)

type configStruct struct {
	Creds string `json:"credentials"`
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

	return nil
}
