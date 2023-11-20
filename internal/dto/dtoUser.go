package dto

import (
	"encoding/json"
	"errors"
	"io"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FromJsonCreateUserRequest(body io.ReadCloser) (*UserRequest, error) {
	userRequest := UserRequest{}
	if err := json.NewDecoder(body).Decode(&userRequest); err != nil {
		return nil, err
	}

	if len(userRequest.Username) < 3 {
		return nil, errors.New("Username deve ter mais de 3 letras")
	}

	if len(userRequest.Password) < 5 {
		return nil, errors.New("Password deve ter mais de 5 letras")
	}

	return &userRequest, nil

}
