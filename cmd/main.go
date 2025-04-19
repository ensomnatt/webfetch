package main

import (
	"github.com/ensomnatt/webfetch/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	app := server.NewServer(":6969")
	logrus.Info("server started")
	err := app.Start()
	if err != nil {
		logrus.Errorf("error with ListenAndServe: %v", err)
	}
}
