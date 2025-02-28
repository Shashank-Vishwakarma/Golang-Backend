package routes

import (
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/handlers"
	"github.com/gorilla/mux"
)

func ProjectRoutes(router *mux.Router) {
	router.HandleFunc("", handlers.CreateProjectHandler).Methods(http.MethodPost)
	router.HandleFunc("", handlers.GetAllProjectsHandler).Methods(http.MethodGet)
	router.HandleFunc("/{project_id}", handlers.GetProjectDetails).Methods(http.MethodGet)
	router.HandleFunc("/{project_id}", handlers.UpdateProjectHandler).Methods(http.MethodPut)
	router.HandleFunc("/{project_id}", handlers.DeleteProjectHandler).Methods(http.MethodDelete)
	router.HandleFunc("/{project_id}/members", handlers.GetAllMembersOnAProject).Methods(http.MethodGet)
	router.HandleFunc("/{project_id}/members", handlers.AddMemberToProject).Methods(http.MethodPost)
	router.HandleFunc("/{project_id}/members/{user_id}", handlers.RemoveMemberFromProject).Methods(http.MethodDelete)
}
