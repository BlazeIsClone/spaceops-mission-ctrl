package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/blazeisclone/spaceops-mission-ctrl/domain/mission"
	"github.com/blazeisclone/spaceops-mission-ctrl/instrumenting"
	mysqlDB "github.com/blazeisclone/spaceops-mission-ctrl/internal/database/mysql"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Env vars not loaded from file")
	}

	router := http.NewServeMux()

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("server listening on port:", port)

	db, err := mysqlDB.Init()
	if err != nil {
		fmt.Println("database init", err)
	}

	defer func() {
		db.Close()
		fmt.Println("db.Closed")
	}()

	instrumenting.Routes(router)
	mission.Routes(router, db)

	server.ListenAndServe()
}
