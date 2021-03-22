package registry

import (
	"github.com/mr4torr/go-lang-graphql/core/user"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

// type Context interface {
// 	JSON(code int, i interface{}) error
// 	Bind(i interface{}) error
// }

type UseCase struct {
	User user.UseCase
}

type Registry interface {
	UseCase() UseCase
}
