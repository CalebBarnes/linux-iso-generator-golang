package server

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/generate", generateIsoHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
