package db

import (
	"log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/harisaginting/ginting/pkg/utils/helper"
)

func Connection() *gorm.DB {
	var (
		dbUser  = helper.MustGetEnv("DB_USER")
		dbPass  = helper.MustGetEnv("DB_PASSWORD")
		dbHost  = helper.MustGetEnv("DB_HOST")
		dbName  = helper.MustGetEnv("DB_NAME")
		dbPort  = helper.MustGetEnv("DB_PORT")
		TZ      = helper.MustGetEnv("TIMEZONE")
		sslMode = helper.MustGetEnv("SSL_MODE")
	)
	// dsn
	dsn := fmt.Sprintf(`
		host=%s user=%s password=%s database=%s port=%s sslmode=%s TimeZone=%s`,
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
		sslMode,
		TZ,
	)

	log.Print("db:", dsn)
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		// set without default transaction
		// will call manually per-case query
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println("Connected to database Failed:", err)
	}
	log.Println("Connected to database")
	return connect
}