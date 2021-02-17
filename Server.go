package main

import (
	"Golang/config"
	"Golang/conn"
	"Golang/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":3000"
	rout := chi.NewRouter()
	userHandler := handler.NewUserHandler()
	//fmt.Println("HERE")
	rout.Route("/users", userHandler.Handle)
	//fmt.Println("here2")
	articleHandler := handler.NewArticleHandler()
	rout.Route("/articles", articleHandler.Handler)
	fmt.Println("Serving on port ", port)
	//upload file handler
	//http.HandleFunc("/upload", handler.UploadHandler)
	rout.Route("/upload", handler.NewUploadsHandler().UploadHandler)
	config2 := config.NewDBConfig()
	connection := conn.ConnectDB(config2)
	connection.Migration()
	defer connection.Close()
	http.ListenAndServe(port, rout)
}
