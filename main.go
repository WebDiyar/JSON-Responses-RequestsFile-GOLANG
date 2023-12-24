package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to match the expected JSON data
type RequestData struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", postHandler) // Set up the handler
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Message == "" {
		http.Error(w, `{"status": "400", "message": "Invalid JSON message"}`, http.StatusBadRequest)
		return
	}

	fmt.Println("Received message:", data.Message)

	// Sending JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "success",
		Message: "Data successfully received",
	})
}
