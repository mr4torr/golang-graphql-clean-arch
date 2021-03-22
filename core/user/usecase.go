package user

type GetUserResponse struct {
	response int
}

type UseCase interface {
	GetUsers() error
	CreateUser() error
}
