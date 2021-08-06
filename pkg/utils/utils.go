package utils

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
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