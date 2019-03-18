package result

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetResult() int {
	return http.StatusOK
}

func GetResultCreated() int {
	return http.StatusCreated
}

func GetResultOK() int {
	return http.StatusOK
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	msgbytes, _ := json.Marshal(err)
	WriteJsonResponse(w, msgbytes, http.StatusBadRequest)
}

func WriteJsonResponse(w http.ResponseWriter, bytes []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err := w.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}
