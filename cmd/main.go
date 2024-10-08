package main

import (
	"fmt"
	"net/http"

	"github.com/kahunacohen/repo-pattern/internal/controllers"
)

func main() {
	http.HandleFunc("/bl/emergency/import", controllers.ImportEmergencyDetails)

	// Start the server on port 8080
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
