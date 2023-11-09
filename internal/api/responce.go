package api

import (
	"encoding/json"
	"net/http"
)

func initHeaders(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
}

func Write(rw http.ResponseWriter, b []byte) {
	initHeaders(rw)
	rw.Write(b)
}

func Json(rw http.ResponseWriter, any interface{}, statCode int) {
	initHeaders(rw)
	rw.WriteHeader(statCode)
	json.NewEncoder(rw).Encode(any)

}
