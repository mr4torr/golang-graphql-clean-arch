package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//ID entity ID
type ID = uuid.UUID

//NewID create a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

//StringToID convert a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
