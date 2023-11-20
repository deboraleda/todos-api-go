package dto

import (
	"encoding/json"
	"errors"
	"io"
)

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"done"`
}

func FromJsonCreateTodoRequest(body io.ReadCloser) (*TodoRequest, error) {
	todoRequest := TodoRequest{}
	if err := json.NewDecoder(body).Decode(&todoRequest); err != nil {
		return nil, err
	}

	if len(todoRequest.Title) < 3 {
		return nil, errors.New("titulo deve ter mais de 3 letras")
	}

	if len(todoRequest.Description) < 3 {
		return nil, errors.New("Description deve ter mais de 3 letras")
	}

	if todoRequest.Done == nil {
		return nil, errors.New("Deve passar a variavel done")
	}

	return &todoRequest, nil

}
