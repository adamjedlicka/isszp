package controller

import (
	"errors"
	"log"

	"isszp/src/model"
)

func CreateUser(userName, firstName, lastName string) error {
	user := model.NewUser()
	user.SetUserName(userName)
	user.SetFirstName(firstName)
	user.SetLastName(lastName)

	err := user.Save()
	if err != nil {
		log.Println("Failed to create new Useer: ", err)
		return errors.New("failed to create new User")
	}

	return nil
}
