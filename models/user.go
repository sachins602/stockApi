package models

import (
	// "errors"
	"html"
	// "log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Fullname string `gorm:"size:255;not null;" json:"fullname"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	// fmt.Println(u.Password)
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) SaveUser() (*User, error) {

	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
// 	var err error
// 	users := []User{}
// 	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
// 	if err != nil {
// 		return &[]User{}, err
// 	}
// 	return &users, err
// }

// func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
// 	var err error
// 	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	if gorm.IsRecordNotFoundError(err) {
// 		return &User{}, errors.New("User Not Found")
// 	}
// 	return u, err
// }

// func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

// 	// To hash the password
// 	err := u.BeforeSave()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
// 		map[string]interface{}{
// 			"password":  u.Password,
// 			"nickname":  u.Nickname,
// 			"email":     u.Email,
// 			"update_at": time.Now(),
// 		},
// 	)
// 	if db.Error != nil {
// 		return &User{}, db.Error
// 	}
// 	// This is the display the updated user
// 	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }

// func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

// 	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

// 	if db.Error != nil {
// 		return 0, db.Error
// 	}
// 	return db.RowsAffected, nil
// }
