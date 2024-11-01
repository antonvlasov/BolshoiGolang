package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	queryCreateUsersTable = `CREATE TABLE IF NOT EXISTS users (
		id bigserial PRIMARY KEY,
		name text NOT NULL,
		password text NOT NULL
	)`

	queryGetUser    = `SELECT * FROM users WHERE id = $1`
	queryCreateUser = `INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id`
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	postgresURI := os.Getenv("POSTGRES")
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Fatal("open", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("ping", err)
	}

	_, err = db.Exec(queryCreateUsersTable)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	users := r.Group("/users")
	users.POST("", func(ctx *gin.Context) {
		var u user

		if err := ctx.Bind(&u); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := db.QueryRow(queryCreateUser, u.Name, u.Password).Scan(&u.ID)
		if err != nil {
			log.Print("create", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusCreated, u)
	})

	users.GET("/:id", func(ctx *gin.Context) {
		userID := ctx.Param("id")

		var u user

		err := db.QueryRow(queryGetUser, userID).Scan(
			&u.ID,
			&u.Name,
			&u.Password,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.Status(http.StatusNotFound)
				return
			}
			log.Print("get", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, u)
	})

	host := os.Getenv("HOST")
	r.Run(host)
}
