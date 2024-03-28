package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var message string

func init() {
	message = "Hello world!"
}

func Start(address string, port string) {
	http.HandleFunc("/", Handle)
	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", address, port),
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Listening error")
	}
}

func Stop() {
	os.Exit(0)
}

func Handle(w http.ResponseWriter, req *http.Request) {
	PrintInfo(*req)
	switch req.Method {
	case "GET":
		fmt.Fprint(w, message)
	case "POST":
		defer req.Body.Close()
		resp, err := UpdateMessage(req.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprint(w, resp)
	}
}

func UpdateMessage(body io.ReadCloser) (string, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}
	message = string(bodyBytes)
	return "Message is updated", nil
}

func PrintInfo(req http.Request) {
	log.Printf("Request method: %s; User Agent: %s; Remote address: %s;",
		req.Method,
		req.UserAgent(),
		req.RemoteAddr)
}
