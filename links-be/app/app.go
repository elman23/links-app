package app

import (
	"fmt"
	"links-be/domain"
	"links-be/handler"
	"links-be/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Startup() {
	router := mux.NewRouter()

	userHandler := handler.NewUserHandler(
		service.NewLocalUserService(
			domain.NewUserRepoMongo()))
	linkHandler := handler.NewLinkHandler(
		service.NewLocalLinkService(
			domain.NewLinkRepoMongo()))

	router.HandleFunc("/create_user", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{user_id}", userHandler.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/delete/{user_id}", userHandler.DeleteUser).Methods(http.MethodGet)

	router.HandleFunc("/create_link", linkHandler.CreateLink).Methods(http.MethodPost)
	router.HandleFunc("/links/{link_id}", linkHandler.FindLink).Methods(http.MethodGet)
	router.HandleFunc("/links", linkHandler.FindLinks).Methods(http.MethodGet)
	router.HandleFunc("/links/delete/{link_id}", linkHandler.DeleteLink).Methods(http.MethodGet)

	fmt.Println("Running at localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		panic("Cannot start the server")
	}
}
