package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)

type ErrNotFound struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var Response = ErrNotFound{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}
	message, _ := json.Marshal(Response)
	w.WriteHeader(Response.StatusCode)
	w.Write(message)
	fmt.Println(message)
}