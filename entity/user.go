package entity

import (
	"gorm.io/gorm"
)

// type NewUser struct {
// 	Name   		string `json:"text"`
// 	Password 	string `json:"password"`
// 	IsAdmin 	bool `json:"is_admin"`
// 	IsActivated bool `json:"is_activated"`
// }

type User struct {
	gorm.Model
	ID					uint32    	`gorm:"primary_key;auto_increment"`
	UUID    			string 		`gorm:"type:uuid" json:"id"`

	Name  				string    	`gorm:"size:50;not null" json:"name"`
	Email     			string    	`gorm:"size:50;not null;unique" json:"email"`
	Password  			string    	`gorm:"size:100;not null;" json:"password"`
	RememberMeToken		string    	`gorm:"size:100;" json:"remember_me_token"`
	ContextID 			uint32
}

// func NewUser(name string, password string, isAdmin int) (*User, error) {
// 	r := &User{
// 		ID:        NewID(),
// 		Title:     title,
// 		Author:    author,
// 		Pages:     pages,
// 		Quantity:  quantity,
// 		CreatedAt: time.Now(),
// 	}
// 	err := r.Validate()
// 	if err != nil {
// 		return nil, ErrInvalidEntity
// 	}
// 	return b, nil
// }

// func (b *User) Validate() error {
// 	if b.Title == "" || b.Author == "" || b.Pages <= 0 || b.Quantity <= 0 {
// 		return ErrInvalidEntity
// 	}
// 	return nil
// }

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = NewID().String()
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

  return
}


// func (u *User) BeforeSave() error {
// 	hashedPassword, err := Hash(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }

// func (u *User) Prepare() {
// 	u.ID = 0
// 	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
// 	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// }

// func (u *User) Validate(action string) error {
// 	switch strings.ToLower(action) {
// 	case "update":
// 		if u.Nickname == "" {
// 			return errors.New("Required Nickname")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}

// 		return nil
// 	case "login":
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil

// 	default:
// 		if u.Nickname == "" {
// 			return errors.New("Required Nickname")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil
// 	}
// }

// func (u *User) SaveUser(db *gorm.DB) (*User, error) {

// 	var err error
// 	err = db.Create(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }

// func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
// 	var err error
// 	users := []User{}
// 	err = db.Model(&User{}).Limit(100).Find(&users).Error
// 	if err != nil {
// 		return &[]User{}, err
// 	}
// 	return &users, err
// }

// func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
// 	var err error
// 	err = db.Model(User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	// if gorm.IsRecordNotFoundError(err) {
// 	// 	return &User{}, errors.New("User Not Found")
// 	// }
// 	return u, err
// }

// func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

// 	// To hash the password
// 	err := u.BeforeSave()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = db.Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
// 		map[string]interface{}{
// 			"password":   u.Password,
// 			"nickname":   u.Nickname,
// 			"email":      u.Email,
// 			"updated_at": time.Now(),
// 		},
// 	)
// 	if db.Error != nil {
// 		return &User{}, db.Error
// 	}
// 	// This is the display the updated user
// 	err = db.Model(&User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }

// func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

// 	db = db.Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

// 	if db.Error != nil {
// 		return 0, db.Error
// 	}
// 	return db.RowsAffected, nil
// }