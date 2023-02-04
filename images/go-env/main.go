package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	var (
		APPLICATION = os.Getenv("APPLICATION")
		VERSION     = os.Getenv("VERSION")
	)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"APPLICATION": APPLICATION,
		"VERSION":     VERSION,
		"ENVIRON":     os.Environ(),
	})
}
