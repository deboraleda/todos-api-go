package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todos-api/internal/dto"

	"github.com/go-chi/chi/v5"
)

func RecoverUserIdFromContext(r *http.Request) (*int64, error) {
	userId, ok := r.Context().Value("userId").(int64)

	if !ok {
		errorMessage := ("Erro ao obter o valor 'user id' do contexto")
		log.Printf(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return &userId, nil
}

func RecoverTodoIdFromURL(r *http.Request) (*int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		errorMessage := ("erro ao fazer parse do id" + fmt.Sprintf("%v", err))
		log.Printf(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return &id, nil
}

func RecoverJsonTodoRequest(r *http.Request) (*dto.TodoRequest, error) {
	todoRequest, err := dto.FromJsonCreateTodoRequest(r.Body)

	if err != nil {
		errorMessage := ("erro ao fazer decode do json - " + fmt.Sprintf("%v", err))
		log.Printf(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return todoRequest, nil
}
