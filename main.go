package main

import (
	"Humo/db"
	"Humo/handlers"
	"Humo/repository"
	"Humo/routs"
	"github.com/gin-gonic/gin"
	"log"
)

// countries

//  primary key, autoIncrement - for tables

//  при пагинации должно возвращаться также страница и количество строк

func main() {

	route := gin.Default()
	DB := db.SetupGorm()

	newRepository := repository.NewRepository(DB)
	newServer := handlers.NewHandler(newRepository)
	newRouts := routs.NewHandler(route, newServer)
	newRouts.Init()

	err := route.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
