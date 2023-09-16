package controllers

import (
	"myapp/models"
	"net/http"

	"database/sql"

	"github.com/labstack/echo/v4"
)

func Index(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		cmd := "SELECT * FROM users"
		rows, err := db.Query(cmd)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "データの取得エラー"})
		}

		var users []models.User
		for rows.Next() {
			var user models.User
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
}

func Insert(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
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
}
