package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/blazeisclone/order-service/domain/organization"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Env vars not loaded from file")
	}

	router := http.NewServeMux()
	organization.Routes(router)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("server listening on port:", port)

	server.ListenAndServe()
}
