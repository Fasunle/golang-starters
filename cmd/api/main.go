package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Fasunle/golang-starters.git/cmd/routes"
)

type Config struct{}

func main() {
	app := Config{}
	//
	app.serve()
}

// start and serve http server
func (app *Config) serve() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", getPort()),
		Handler: routes.New(),
	}

	defer server.Close()

	server.RegisterOnShutdown(func() {
		fmt.Println("Server is shutting down...")
	})

	fmt.Println("Server started on port", getPort())

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// use the set port in env or default to 80
func getPort() string {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "80"
	}

	return PORT
}
