package handlers

import (
	"fmt"
	"log"
	"net/http"
	"todos-api/internal/models"
	"todos-api/internal/utils"
)

func List(w http.ResponseWriter, r *http.Request) {
	userId, err := RecoverUserIdFromContext(r)
	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 500)
		return
	}

	todos, err := models.GetAllFromUser(*userId)

	if err != nil {
		errorMessage := ("Erro ao obter registros" + fmt.Sprintf("%v", err))
		log.Printf(errorMessage)
		utils.GenerateErrorResponse(w, errorMessage, 400)
		return
	}

	utils.GenerateResponse(w, todos, 200)
}

func Update(w http.ResponseWriter, r *http.Request) {
	userId, err := RecoverUserIdFromContext(r)
	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 500)
		return
	}

	id, err := RecoverTodoIdFromURL(r)

	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 500)
		return
	}

	todoRequest, err := RecoverJsonTodoRequest(r)
	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 400)
		return
	}

	todo := models.Todo{Title: todoRequest.Title, Description: todoRequest.Description, Done: *todoRequest.Done, User_id: *userId}

	rows, err := models.Update(int64(*id), todo)

	if err != nil {
		errorMessage := ("erro ao atualizar registro")
		log.Printf(errorMessage)
		utils.GenerateErrorResponse(w, errorMessage, 400)
		return
	}

	if rows > 1 {
		log.Printf("Erro: atualizou número %d de registros", rows)
	}

	utils.GenerateResponse(w, nil, 200)

}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := RecoverTodoIdFromURL(r)

	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 500)
		return
	}

	todo, err := models.Get(int64(*id))

	if err != nil {
		errorMessage := ("erro ao recuperar o registro " + fmt.Sprintf("%v", err))
		log.Printf(errorMessage)
		utils.GenerateErrorResponse(w, errorMessage, 400)
		return
	}

	utils.GenerateResponse(w, todo, 200)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := RecoverTodoIdFromURL(r)

	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 500)
		return
	}

	rows, err := models.Delete(int64(*id))

	if err != nil {
		errorMessage := ("erro ao deletar registro")
		log.Printf(errorMessage)
		utils.GenerateErrorResponse(w, errorMessage, 400)
		return
	}

	if rows > 1 {
		log.Printf("Erro: foram removidos %d de registros", rows)
	}

	utils.GenerateResponse(w, nil, 200)

}

func Create(w http.ResponseWriter, r *http.Request) {

	userId, err := RecoverUserIdFromContext(r)
	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 500)
		return
	}

	todoRequest, err := RecoverJsonTodoRequest(r)
	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 400)
		return
	}

	todo := models.Todo{Title: todoRequest.Title, Description: todoRequest.Description, Done: *todoRequest.Done, User_id: *userId}

	id, err := models.Insert(todo, int64(*userId))

	var resp map[string]any
	if err != nil {
		errorMessage := ("erro - " + fmt.Sprintf("%v", err))
		log.Printf(errorMessage)
		utils.GenerateErrorResponse(w, errorMessage, 400)
		return
	} else {
		resp = map[string]any{
			"id": id,
		}
		utils.GenerateResponse(w, resp, 200)
	}

}
