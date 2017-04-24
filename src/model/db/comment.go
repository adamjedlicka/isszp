package db

import (
	"fmt"

	"isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewComment = NewComment
	model.QueryComments = QueryComments
}

type Comment struct {
	ID           string
	Text         string
	PostDateTime string

	UserID string
	User   User
	TaskID string
	Task   Task
}

func NewComment() model.Comment {
	c := new(Comment)

	return c
}

func (c *Comment) FillByID(id string) error {
	return db.First(c, "id = ?", id).Error
}

func (c *Comment) Save() error {
	return db.Save(c).Error
}

func (c *Comment) Delete() error {
	return db.Delete(c).Error
}

func (c Comment) GetID() string           { return c.ID }
func (c Comment) GetText() string         { return c.Text }
func (c Comment) GetPostDateTime() string { return c.PostDateTime }

func (c *Comment) SetText(val string)         { c.Text = val }
func (c *Comment) SetPostDateTime(val string) { c.PostDateTime = val }

func (c *Comment) GetUser() model.User {
	db.Model(c).Related(&c.User)
	return &c.User
}

func (c *Comment) SetUser(val model.User) { c.UserID = val.GetID() }

func (c *Comment) GetTask() model.Task {
	db.Model(c).Related(&c.Task)
	return &c.Task
}

func (c *Comment) SetTask(val model.Task) { c.TaskID = val.GetID() }

func (c Comment) String() string {
	return fmt.Sprint("Comment: ", c.Text)
}

// BeforeCreate is a GORM hook
func (*Comment) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryComments(args ...interface{}) []model.Comment {
	comments := []*Comment{}

	db.Find(&comments, args...)

	ret := make([]model.Comment, len(comments))
	for k, v := range comments {
		ret[k] = v
	}

	return ret
}
