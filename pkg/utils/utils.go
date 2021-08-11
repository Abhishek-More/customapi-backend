package utils

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"encoding/json"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func CheckFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error, w http.ResponseWriter, code int) {
	if err != nil {
		w.WriteHeader(code)
		return
	}
}

func RouteNotImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not Implemented Yet!")
}

func GetError(err error, w http.ResponseWriter) {

	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

func GetPort() string{
	port := os.Getenv("PORT")
	if port == "" {
		return "8000"
	}
	return port
}