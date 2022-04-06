package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/janpavtel/site/internal/drivers"
	"bitbucket.org/janpavtel/site/internal/handlers"
	"bitbucket.org/janpavtel/site/internal/routes"
)

const portNumber = ":8080"

func main() {

	db, err := drivers.ConnectSQL("host=localhost port=5432 dbname=site user=postgres password=admin")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying ...")
	}
	defer db.SQL.Close()

	log.Println("Connected to database")

	fmt.Printf("Starting application on port %s", portNumber)

	view := handlers.NewView(db)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes.CreateRoutes(view),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
