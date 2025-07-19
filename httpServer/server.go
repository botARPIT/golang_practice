package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the go server")
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "This is a json message, from a go server",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening from the port from this computer")
	fmt.Println("This is for testing")
	http.ListenAndServe(":8080", nil)
}
