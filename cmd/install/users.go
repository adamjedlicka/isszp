package install

import (
	"log"

	"gitlab.fit.cvut.cz/isszp/isszp/src/controller"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

// InstallUsers hashes password for default users. It usses UserNames as new passwords
func InstallUsers() {
	log.Println("Installing users...")

	users := model.QueryUsers()
	for _, v := range users {
		log.Println("Setting password for: ", v.GetUserName())
		controller.SetUserHashedPassword(v, v.GetUserName())
		err := v.Save()
		if err != nil {
			log.Fatal(err)
		}
	}
}
