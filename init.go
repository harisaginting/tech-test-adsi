package main

import (
	"os"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// init before exec main
func init() {
	//Log as JSON instead of the default ASCII formatter.
  	log.SetFormatter(&log.JSONFormatter{})
	
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		log.Error("Cannot load file .env: ", err)
		panic(err)
	}
	log.Info("ENV:", os.Getenv("MODE"))

}