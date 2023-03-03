package main

import (
	"log"

	"github.com/ishanshre/Go-blog/api/v1/db"
	"github.com/ishanshre/Go-blog/api/v1/handlerAndRouters"
)

func main() {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatalln(err)
	}
	server := handlerAndRouters.NewApiServer(":8000", store)
	server.Run()

}
