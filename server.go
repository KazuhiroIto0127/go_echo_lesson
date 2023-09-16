package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	databasePath = "./example.sql"
	serverAddr = ":8080"
)

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := initDB(databasePath)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	e.GET("/", func(c echo.Context) error { return index(c, db) })
	e.GET("/insert", func(c echo.Context) error { return insert(c, db) })

	e.Logger.Fatal(e.Start(":8080"))
}

func initDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func insert(c echo.Context, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO users (name, age, email) VALUES (?, ?, ?)")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "データベースエラー"})
	}
	defer stmt.Close()

	if _, err = stmt.Exec("Kazu", 33, "xxxxxx@omomuro.dev"); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "データ挿入エラー"})
	}

	return c.String(http.StatusOK, "inserted.")
}

type User struct {
	ID int      `json:"id"`
	Name string `json:"name"`
	Age int     `json:"age"`
	Email sql.NullString `json:"email"`
}

func index(c echo.Context, db *sql.DB) error {
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
