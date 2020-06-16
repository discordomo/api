package main

import (
	"github.com/discordomo/api/router"
	"github.com/sirupsen/logrus"
)

func main() {

	router := router.Load()

	// set logrus to log in JSON format
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := router.Run(); err != nil {
		logrus.Fatal(err)
	}
}
