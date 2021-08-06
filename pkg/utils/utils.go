package utils

import (
	"net/http"
	"log"
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RouteNotImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not Implemented Yet!")
}