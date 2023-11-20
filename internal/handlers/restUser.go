package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todos-api/internal/models"
	"todos-api/internal/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	userRequest, err := RecoverJsonUserRequest(r)
	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 400)
		return
	}

	user := models.User{Username: userRequest.Username, Password: userRequest.Password}
	recoveredUser, err := models.GetByUserName(user.Username, user.Password)

	if err != nil {
		log.Printf("usuário não encontrado - %v", err)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	hashedpass, err := HashPassword(user.Password)

	ispassWrong := CheckPasswordHash(hashedpass, recoveredUser.Password)

	if ispassWrong {
		log.Printf("senha incorreta")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	token, _ := generateJWT(recoveredUser.Username, recoveredUser.ID)

	w.Write([]byte(token))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest, err := RecoverJsonUserRequest(r)
	if err != nil {
		utils.GenerateErrorResponse(w, fmt.Sprintf("%v", err), 400)
		return
	}

	encryptedpass, err := HashPassword(userRequest.Password)
	if err != nil {
		log.Printf("erro ao fazer encrypt da password %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user := models.User{Username: userRequest.Username, Password: encryptedpass}

	id, err := models.InsertUser(user)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": fmt.Sprintf("erro: %v", err),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": fmt.Sprintf("user inserido com sucesso - id: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(username string, id int64) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		Id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)

}
