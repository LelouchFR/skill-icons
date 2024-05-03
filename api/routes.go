package api

import (
    "github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
    router.HandleFunc("/icons", GetSkillIcons).Methods("GET")
}
