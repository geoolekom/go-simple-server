package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/geoolekom/go-simple-server/views"
	"github.com/geoolekom/go-simple-server/database"
	"github.com/geoolekom/go-simple-server/models"
	"github.com/geoolekom/go-simple-server/parser"
	"fmt"
)

func main() {
	db, err := database.InitDatabase("dbname=gosimpleserver user=go password=go port=5432")
	defer db.Close()
	fmt.Println("Connected to Db.")

	if err != nil {
		log.Fatal(err)
	}
	m := models.New(db)
	parser.LoadData(m)
	fmt.Println("Data is loading in goroutine.")

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(views.NotFoundHandler)

	router.GET("/locations/:id", views.GetLocationHandler(m))
	router.POST("/locations/:id", views.PostLocationHandler(m))

	router.GET("/users/:id", views.GetUserHandler(m))
	router.POST("/users/:id", views.PostUserHandler(m))

	router.GET("/visits/:id", views.GetVisitHandler(m))
	router.POST("/visits/:id", views.PostVisitHandler(m))

	fmt.Println("Now serving.")
	log.Fatal(http.ListenAndServe(":80", router))
}
