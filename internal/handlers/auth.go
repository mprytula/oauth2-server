package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"qweinke/oauth/api/types"
	"qweinke/oauth/internal"
	"qweinke/oauth/internal/tools"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var code = r.URL.Query().Get("code")
		fmt.Println("[ LOG ] Auth code: " + code)
		var config = internal.ReadOauthConfig()
		json, err := (&types.TokenExchangeRequest{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			Code:         code,
		}).CreateRequestJSON()
		if err != nil {
			slog.Error(err.Error())
		}

		var token, _ = exchangeToken("https://github.com/login/oauth/access_token", &json)
		fmt.Println("[ LOG ] Token: " + token)

		userData, err := tools.FetchUserData(token)
		if err != nil {
			slog.Error(err.Error())
		}
		fmt.Println(userData)

		if next != nil {
			next.ServeHTTP(w, r)
		}

	}
}

func exchangeToken(exchangeURL string, tokenJSON *[]byte) (string, error) {
	var reader = bytes.NewReader(*tokenJSON)
	var resp, err = http.Post(exchangeURL, "application/json", reader)
	if resp.StatusCode != 200 {
		err = errors.New("exchange request not succeed")
		return "", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	dataMap, err := url.ParseQuery(string(data))
	if err != nil {
		return "", err
	}
	var token = dataMap["access_token"][0]
	return token, err
}
