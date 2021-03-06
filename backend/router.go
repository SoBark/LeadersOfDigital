package main

import (
	"LeadersOfDigital/backend/controllers/auth"
	"LeadersOfDigital/backend/controllers/common"
	"LeadersOfDigital/backend/controllers/employee"
	"LeadersOfDigital/backend/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Base API router
	api   := router.PathPrefix("/api").Subrouter()

	// Secondary routers
	clientAPI := api.PathPrefix("/clientAPI").Subrouter()
	employeeAPI := api.PathPrefix("/employeeAPI").Subrouter()
	authAPI := api.PathPrefix("/auth").Subrouter()

	authAPI.HandleFunc("/register", auth.Registration).Methods(http.MethodPost, http.MethodOptions)
	authAPI.HandleFunc("/login", auth.Login).Methods(http.MethodPost, http.MethodOptions)

	clientAPI.HandleFunc("/me", nil).Methods(http.MethodGet, http.MethodOptions)

	clientAPI.HandleFunc("/applications", nil).Methods(http.MethodGet, http.MethodOptions)
	clientAPI.HandleFunc("/application", nil).Methods(http.MethodPost, http.MethodOptions)
	clientAPI.HandleFunc("/application/{id}", nil).Methods(http.MethodGet, http.MethodOptions)
	clientAPI.HandleFunc("/application/{id}", nil).Methods(http.MethodDelete, http.MethodOptions)
	clientAPI.HandleFunc("/application/{id}/changelog", nil).Methods(http.MethodGet, http.MethodOptions)

	clientAPI.HandleFunc("/service_types", common.GetAllServiceTypes).Methods(http.MethodGet, http.MethodOptions)

	clientAPI.HandleFunc("/application/{id}/add_document", nil).Methods(http.MethodPost, http.MethodOptions)
	clientAPI.HandleFunc("/document/{id}", nil).Methods(http.MethodGet, http.MethodOptions)
	clientAPI.HandleFunc("/document/{id}", nil).Methods(http.MethodDelete, http.MethodOptions)

	employeeAPI.HandleFunc("/me", nil).Methods(http.MethodGet, http.MethodOptions)

	employeeAPI.HandleFunc("/free_applications", nil).Methods(http.MethodGet, http.MethodOptions)
	employeeAPI.HandleFunc("/applications", nil).Methods(http.MethodGet, http.MethodOptions)
	employeeAPI.HandleFunc("/application/{id}", nil).Methods(http.MethodGet, http.MethodOptions)
	employeeAPI.HandleFunc("/application/{id}", nil).Methods(http.MethodPut, http.MethodOptions)
	employeeAPI.HandleFunc("/application/{id}/changelog", nil).Methods(http.MethodGet, http.MethodOptions)

	employeeAPI.HandleFunc("/application_statuses", employee.GetAllApplicationStatuses).Methods(http.MethodGet, http.MethodOptions)

	// middleware usage
	// do NOT modify the order
	api.Use(middleware.CORS)    // enable CORS headers
	api.Use(middleware.LogPath) // log HTTP request URI and method
	api.Use(middleware.LogBody) // log HTTP request body

	clientAPI.Use(middleware.JwtAuthentication)   // check JWT token
	employeeAPI.Use(middleware.JwtAuthentication) // check JWT token

	return router
}
