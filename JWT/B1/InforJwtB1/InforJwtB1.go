package InforJwtB1

import (
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	//	gorm.Model
	Macc         string `json:"ma_cc" gorm:"unique"`
	Username     string `json:"user_name" gorm:"unique"`
	Userpassword string `json:"user_password"`
	Address      string `json:"address" gorm:"unique"`
	Phonenumber  string `json:"phonenumber" gorm:"unique"`
}

func (user *Users) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Userpassword = string(bytes)
	return nil
}
func (user *Users) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Userpassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
