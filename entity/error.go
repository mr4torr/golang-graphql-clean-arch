package entity

import "errors"

var (
	//ErrNotFound not found
	ErrNotFound = errors.New("Not found")

	//ErrInvalidEntity invalid entity
	ErrInvalidEntity = errors.New("Invalid entity")

	//ErrCannotBeDeleted cannot be deleted
	ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

	//ErrNotEnoughBooks cannot borrow
	ErrNotEnoughBooks = errors.New("Not enough books")

	//ErrBookAlreadyBorrowed cannot borrow
	ErrBookAlreadyBorrowed = errors.New("Book already borrowed")

	//ErrBookNotBorrowed cannot return
	ErrBookNotBorrowed = errors.New("Book not borrowed")
)