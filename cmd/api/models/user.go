package models

import (
	"LaburAR/cmd/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string `json:"email"`
	Password        string `json:"password"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	DNI             int    `json:"dni"`
	City            string `json:"city"`
	Province        string `json:"province"`
	AddressStreet   string `json:"address_street"`
	AddressNumber   string `json:"address_number"`
	Apartment       string `json:"apartment"`
	AccountVerified bool   `json:"account_verified"`
	Role            string `json:"role"`
}

func (user *User) CreateUserRecord() error {
	result := database.PostgresDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
