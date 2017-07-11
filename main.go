package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/geoolekom/go-simple-server/views"
	"github.com/geoolekom/go-simple-server/database"
	"github.com/geoolekom/go-simple-server/models"
)

func main() {

	db, err := database.InitDatabase("dbname=gosimpleserver user=postgres password=postgres port=5432")
	if err != nil {
		log.Fatal(err)
	}
	m := models.New(db)

	router := httprouter.New()
	router.GET("/locations/:id", views.GetLocationHandler(m))
	router.GET("/users/:id", views.GetUserHandler(m))

	log.Fatal(http.ListenAndServe(":9000", router))
}
