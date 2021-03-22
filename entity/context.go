package entity

import (
	"gorm.io/gorm"
)

type Context struct {
	gorm.Model
	ID		uint32    	`gorm:"primary_key;auto_increment"`
	UUID	string 		`gorm:"type:uuid" json:"id"`

	Name	string    	`gorm:"size:50;not null" json:"name"`
	Active	bool    	`gorm:"size:60;default:false" json:"email"`
	Users	[]User
}

func (u *Context) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = NewID().String()

  return
}
