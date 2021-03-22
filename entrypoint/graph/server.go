package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/mr4torr/go-lang-graphql/entrypoint/graph/generated"
)

func ServerStart(resolver Resolver) echo.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}),
	)

	return echo.WrapHandler(h)
	// return func(c echo.Context) error {
		// return h.ServeHTTP(
		// 	// c.Response().Writer,
		// 	c.Echo().AcquireContext().Response().Writer,
		// 	c.Request(),
		// )
		// return c.String(http.StatusOK, h)
	// }

}

// func ServerStart(resolver Resolver) *handler.Server {
// 	return handler.NewDefaultServer(
// 		generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}),
// 	)
// }

func Playground() echo.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/api")
	return echo.WrapHandler(h)
}
