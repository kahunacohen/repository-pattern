package main

import (
	"fmt"
	"net/http"

	"github.com/kahunacohen/repo-pattern/controllers"
)

func main() {

	http.HandleFunc("/users/export", controllers.ExportUsers)

	// Start the server on port 8080
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
