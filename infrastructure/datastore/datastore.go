package datastore

import (
	"fmt"
	"os"

	"github.com/mr4torr/go-lang-graphql/infrastructure/datastore/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	if os.Getenv("DB_DRIVER") == "postgres" {
		return postgres.Connect()
	}

	panic(fmt.Errorf("System only allow Postgres database driver"))
}
