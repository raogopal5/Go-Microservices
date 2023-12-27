package main

import (
	"Go-Microservices/pkg/controllers"
	"Go-Microservices/pkg/db"
	"Go-Microservices/pkg/repositories"
	"Go-Microservices/pkg/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	db := db.ConnectDB()
	defer db.Close()
	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	router := mux.NewRouter()
	fmt.Println("Server is running on Port 8081")

	router.HandleFunc("/api/v1/item", itemController.GetItemList).Methods("GET")
	router.HandleFunc("/api/v1/item", itemController.CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/item/{id}", itemController.UpdateItem).Methods("PUT")
	router.HandleFunc("/api/v1/item/{id}", itemController.DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/v1/item/{id}", itemController.GetItem).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))

}
