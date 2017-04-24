package db

import (
	"fmt"

	"isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewTimeRecord = NewTimeRecord
}

type TimeRecord struct {
	ID string

	Description string
	Date        string
	Start       string
	Stop        string

	UserID string
	User   User
	TaskID string
	Task   Task

	DeletedAt string
}

func NewTimeRecord() model.TimeRecord {
	t := new(TimeRecord)

	return t
}

func (t *TimeRecord) FillByID(id string) error {
	return db.First(t, "id = ?", id).Error
}

func (t *TimeRecord) Save() error {
	return db.Save(t).Error
}

func (t *TimeRecord) Delete() error {
	return db.Delete(t).Error
}
func (t TimeRecord) GetID() string          { return t.ID }
func (t TimeRecord) GetDescription() string { return t.Description }
func (t TimeRecord) GetDate() string        { return t.Date }
func (t TimeRecord) GetStart() string       { return t.Start }
func (t TimeRecord) GetStop() string        { return t.Stop }

func (t *TimeRecord) SetDescription(val string) { t.Description = val }
func (t *TimeRecord) SetDate(val string)        { t.Date = val }
func (t *TimeRecord) SetStart(val string)       { t.Start = val }
func (t *TimeRecord) SetStop(val string)        { t.Stop = val }

func (t *TimeRecord) GetUser() model.User {
	db.Model(t).Related(&t.User)
	return &t.User
}

func (t *TimeRecord) SetUser(val model.User) { t.UserID = val.GetID() }

func (t *TimeRecord) GetTask() model.Task {
	db.Model(t).Related(&t.Task)
	return &t.Task
}

func (t *TimeRecord) SetTask(val model.Task) { t.TaskID = val.GetID() }

func (t TimeRecord) String() string {
	return fmt.Sprint("TimeRecord: ", t.Description)
}

// BeforeCreate is a GORM hook
func (*TimeRecord) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}
