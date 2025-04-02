package types

import (
	"encoding/json"
	"errors"
)

type TokenExchangeRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

func (t *TokenExchangeRequest) CreateRequestJSON() ([]byte, error) {
	var err error
	if (t.ClientID == "" || t.ClientSecret == "" || t.Code == "") {
		err = errors.New("ClientID, ClientSecret and Code  are required")
		return nil, err
	}
	json, _ := json.Marshal(t);
   	return json, err
}