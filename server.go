package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mr4torr/go-lang-graphql/entrypoint/api"
	"github.com/mr4torr/go-lang-graphql/entrypoint/graph"
	"github.com/mr4torr/go-lang-graphql/infrastructure/datastore"
	"github.com/mr4torr/go-lang-graphql/registry"
	"gorm.io/gorm/logger"
)

// import "github.com/mahmudinashar/go/api"

// directive @hasRole(role: Role!) on FIELD_DEFINITION
// type DirectiveRoot struct {
// 	HasRole func(ctx context.Context, obj interface{}, next graphql.Resolver, role Role) (res interface{}, err error)
// }

// directive @hasRole(role: Role!) on FIELD_DEFINITION

// enum Role {
//     ADMIN
//     USER
// }

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
 	e.Use(middleware.Recover())

	db := datastore.Init()
	sqlDB, err := db.DB()
	logFatal(err)
	db.Logger.LogMode(logger.Info)
	defer sqlDB.Close()

	// _graph := graph.ServerStart(graph.Resolver{ DB: db })
	_registry := registry.Start(db)

	// _registry.NewAppController()

	e.GET("/", api.Welcome())
	e.POST("/api", graph.ServerStart(graph.Resolver{ UseCase: _registry.UseCase()  }))
	e.GET("/playground", graph.Playground())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
	err = e.Start(":" + port)
	logFatal(err)
}


func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}