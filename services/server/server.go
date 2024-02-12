package server

import (
	"log"
	"net/http"

	"github.com/calebbarnes/linux-iso-generator-golang/services/generator"
)

func StartServer() {
	generator.EnsureUbuntuIsoExists()

	http.HandleFunc("/generate", generateIsoHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
