package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"

	"myapp/api/controllers"
	"myapp/config"
)

const serverAddr = ":8080"

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := config.InitDB(config.DatabasePath)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	e.GET("/", controllers.Index(db))
	e.GET("/insert", controllers.Insert(db) )

	e.Logger.Fatal(e.Start(serverAddr))
}
