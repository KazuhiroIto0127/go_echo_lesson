package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/insert", insert)

	e.Logger.Fatal(e.Start(":8080"))
}

func insert(c echo.Context) error {
	db, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO users (name, age, email) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec("Kazu", 33, "xxxxxx@omomuro.dev"); err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, "inserted.")
}

type User struct {
	ID int      `json:"id"`
	Name string `json:"name"`
	Age int     `json:"age"`
	Email sql.NullString `json:"email"`
}

func hello(c echo.Context) error {
	db, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cmd := "SELECT * FROM users"
	rows, err := db.Query(cmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "データの取得エラー"})
	}

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "データの取得エラー"})
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "データの読み取りエラー"})
	}

	return c.JSON(http.StatusOK, users)
}
