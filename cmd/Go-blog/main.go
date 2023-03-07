package main

import (
	"log"

	"github.com/ishanshre/Go-blog/api/v1/db"
	"github.com/ishanshre/Go-blog/api/v1/handlerAndRouters"
)

func main() {
	store, err := db.NewPostgresStore() // connect to databse and returns if everything is ok
	if err != nil {
		log.Fatalln(err) // logs an error that occurs when connecting to database
	}
	server := handlerAndRouters.NewApiServer(":8000", store) // create server
	server.Run()                                             // run server
}
