package main

import (
	"log"

	"github.com/ishanshre/Go-blog/api/v1/db"
	routers "github.com/ishanshre/Go-blog/api/v1/handlerAndRouters/routers"
)

func main() {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatalln(err)
	}
	server := routers.NewApiServer(":8000", store)
	server.Run()

}
