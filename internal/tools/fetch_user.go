package tools

import (
	"io"
	"net/http"
	"qweinke/oauth/api/types"
)

const FETCH_FROM = "https://api.github.com/user"

func FetchUserData(authToken string) (*types.UserData, error) {
	var outcomingRequest, _ = http.NewRequest("GET", FETCH_FROM, nil)
	outcomingRequest.Header.Set("Authorization", "Bearer "+authToken)
	outcomingRequest.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	outcomingRequest.Header.Set("Accept", "application/vnd.github+json")
	var client = &http.Client{}

	response, err := client.Do(outcomingRequest)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	user, err := (&types.UserData{}).FromJSON(data)
	return user, err
}
