package usecase

import "github.com/mr4torr/go-lang-graphql/core/user"

func getUsers() *user.GetUserResponse {
	value := user.GetUserResponse{response: 1}

	return &value
}