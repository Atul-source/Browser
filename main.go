package main

import (
	"fmt"
	"log"
	"net/http"
)

// server side code
func main() {
	fmt.Println("Welcome to File Server")

	var serve_dir string
	fmt.Println("Enter the Path that you want to serve")
	fmt.Scanln(&serve_dir)

	// Creating File server using http package
	fs := http.FileServer(http.Dir(serve_dir))

	// Serving files or directory on localhost:9000

	err := http.ListenAndServe(":9000", fs)
	if err != nil {
		log.Fatalf("Server error :%v", err)
	}
}
