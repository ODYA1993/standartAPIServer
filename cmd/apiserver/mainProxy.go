package main

import (
	"bytes"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/app/api"
	"io"
	"log"
	"net/http"
)

const (
	proxyAddr = "localhost:9006"
)

var (
	counter   = true
	firstURL  = "http://localhost:8085"
	secondURL = "http://localhost:8086"
)

func main() {
	log.Println("Started Proxy...")
	http.HandleFunc("/", handleProxy)
	log.Fatal(http.ListenAndServe(proxyAddr, nil))
}

func handleProxy(rw http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.RequestURI()
	baseURL := firstURL
	method := r.Method

	textByte, err := io.ReadAll(r.Body)

	if counter {
		baseURL = secondURL
	}

	req, err := http.NewRequest(method, baseURL+urlPath, bytes.NewBuffer(textByte))
	if err != nil {
		log.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(method, baseURL+urlPath)

	textByte, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	defer func() {
		counter = !counter
	}()

	api.Write(rw, textByte)
	return
}
