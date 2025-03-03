package main

import (
	"fmt"
	"log"

	"github.com/ensomnatt/webfetch/server"
)

func main() {
	app := server.NewServer(":8080")
	log.Println("server is up")
	err := app.Start()
	if err != nil {
		_ = fmt.Errorf("error with ListenAndServe: %v", err)
	}
}
