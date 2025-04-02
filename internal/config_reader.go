package internal

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type OauthConfig struct {
	ClientID 		string 	`json:"clientID"`
	ClientSecret	string 	`json:"clientSecret"`
}

func ReadOauthConfig() *OauthConfig {
	basePath, _ := os.Getwd()
    configPath := filepath.Join(basePath, "configs", "oauth.config.json")
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	var config *OauthConfig = &OauthConfig{}
	err = json.Unmarshal(configFile, config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}