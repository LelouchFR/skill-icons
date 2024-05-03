package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/lelouchfr/skill-icons/api"
)

func main() {
    router := mux.NewRouter()

    api.SetupRoutes(router)
    http.ListenAndServe(":8080", router)
}
