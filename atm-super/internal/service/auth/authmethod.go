package auth

import (
	"atm/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
)

func (a AuthService) CreateNewUser(user models.User) (map[string]string, error) {
	userData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(a.AuthURL+"/register", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		return nil, err
	}

	result := map[string]string{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}