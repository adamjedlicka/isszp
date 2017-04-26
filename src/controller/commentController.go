package controller

import (
	"time"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"
)

func NewComment(userName, taskID, text string) error {
	u := model.NewUser()
	err := u.FillByUserName(userName)
	if err != nil {
		return err
	}

	t := model.NewTask()
	err = t.FillByID(taskID)
	if err != nil {
		return err
	}

	c := model.NewComment()
	c.SetTask(t)
	c.SetUser(u)
	c.SetText(text)
	c.SetPostDateTime(time.Now().Format(common.DateTimeFormat))

	return c.Save()
}
