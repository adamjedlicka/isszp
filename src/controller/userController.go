package controller

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

func SetUserHashedPassword(user model.User, password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.SetPassword(string(hashedPassword))
}

func CheckUserHashedPassword(user model.User, password string) bool {
	if user.GetPassword() == "" {
		return true
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(password))
	if err != nil {
		log.Println(err)
	}

	return err == nil
}
