package main

import (
	"github.com/mr4torr/go-lang-graphql/entity"
	"github.com/mr4torr/go-lang-graphql/infrastructure/datastore"
)

func main() {
	db := datastore.Init()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	defer sqlDB.Close()

	db.AutoMigrate(
		&entity.Context{},
		&entity.User{},
	)
}
