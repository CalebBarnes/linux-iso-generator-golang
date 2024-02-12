package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/calebbarnes/linux-iso-generator-golang/services/generator"
	"github.com/calebbarnes/linux-iso-generator-golang/services/templates"
)

var isDebug = false

type requestData struct {
	Hostname string   `json:"hostname"`
	User     string   `json:"user"`
	Password string   `json:"password"`
	SSHKeys  []string `json:"ssh_keys"`
}

type responseData struct {
	Status string `json:"status"`
}

func generateIsoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req requestData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received request to generate ISO for host %s", req.Hostname)

	userDataString := templates.GetUserDataConfig(templates.UserData{
		Hostname: req.Hostname,
		Username: req.User,
		Password: req.Password,
		SSHKeys:  req.SSHKeys,
	})

	if isDebug {
		log.Println("UserData:")
		log.Println(userDataString)
	}

	err := generator.GenerateIso(userDataString)
	if err != nil {
		log.Printf("Failed to generate ISO: %v", err)
		http.Error(w, "Failed to generate ISO", http.StatusInternalServerError)
		return
	}
	log.Println("ISO generation initiated.")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData{Status: "pending"})
}
