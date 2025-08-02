package main

import (
	"log"
	"net/http"
	"github.com/lelouchfr/skill-icons/api/icons"
)

func main() {
	http.HandleFunc("/api/icons", icons.Handler)

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}