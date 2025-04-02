package types

import "encoding/json"

type UserData struct {
	Email     string `json:"email"`
	UserName  string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

func (u *UserData) FromJSON(data []byte) (*UserData, error) {
	var err = json.Unmarshal(data, u)
	return u, err
}
