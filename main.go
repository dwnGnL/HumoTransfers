package main

// countries

// todo: primary key, autoIncrement - for tables

// todo при пагинации должно возвращаться также страница и количество строк

import (
	"Humo/db"
	"Humo/handlers"
	"Humo/repository"
	"Humo/routs"
	"github.com/gin-gonic/gin"
	"log"
)

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
