package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dataSourceName() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)
}

func Connect() *gorm.DB {
	conn, err := gorm.Open(
		postgres.Open(dataSourceName()),
		&gorm.Config{},
	)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(`database started`)
	}

	return conn
}
