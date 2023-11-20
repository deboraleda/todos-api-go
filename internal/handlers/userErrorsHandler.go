package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"todos-api/internal/dto"
)

func RecoverJsonUserRequest(r *http.Request) (*dto.UserRequest, error) {
	userRequest, err := dto.FromJsonCreateUserRequest(r.Body)

	if err != nil {
		errorMessage := ("erro ao fazer decode do json - " + fmt.Sprintf("%v", err))
		log.Printf(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return userRequest, nil
}
