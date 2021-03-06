package db

import (
	"fmt"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"

	"time"

	"errors"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewTimeRecord = NewTimeRecord
	model.QueryTimeRecords = QueryTimeRecords
}

type TimeRecord struct {
	ID string

	Description string
	Date        string
	Start       string
	End         *string
	TimeInMs    string

	UserID string
	TaskID string

	DeletedAt *string
}

func NewTimeRecord() model.TimeRecord {
	t := new(TimeRecord)

	return t
}

func (t *TimeRecord) FillByID(id string) error {
	return db.First(t, "id = ?", id).Error
}

func (t *TimeRecord) Save() error {
	if t.End != nil {
		start, err := time.Parse("15:04:05", t.Start)
		if err != nil {
			return err
		}

		end, err := time.Parse("15:04:05", *t.End)
		if err != nil {
			return err
		}

		if end.Before(start) {
			return errors.New("Cannot set 'Stop time' before 'Start time")

		}
	}

	return db.Save(t).Error
}

func (t *TimeRecord) Delete() error {
	return db.Delete(t).Error
}
func (t TimeRecord) GetID() string          { return t.ID }
func (t TimeRecord) GetDescription() string { return t.Description }
func (t TimeRecord) GetDate() string        { return t.Date }
func (t TimeRecord) GetStart() string       { return t.Start }
func (t TimeRecord) GetStop() *string       { return t.End }
func (t TimeRecord) GetTimeInMs() string    { return t.TimeInMs }

func (t *TimeRecord) SetDescription(val string) { t.Description = val }
func (t *TimeRecord) SetDate(val string)        { t.Date = val }
func (t *TimeRecord) SetStart(val string)       { t.Start = val }
func (t *TimeRecord) SetStop(val *string)       { t.End = val }
func (t *TimeRecord) SetTimeInMs(val string)    { t.TimeInMs = val }

func (t TimeRecord) GetUser() model.User {
	u := model.NewUser()
	u.FillByID(t.UserID)
	return u
}

func (t *TimeRecord) SetUser(val model.User) { t.UserID = val.GetID() }

func (t TimeRecord) GetTask() model.Task {
	ta := model.NewTask()
	ta.FillByID(t.TaskID)
	return ta
}

func (t *TimeRecord) SetTask(val model.Task) { t.TaskID = val.GetID() }

func (t TimeRecord) String() string {
	return fmt.Sprint("TimeRecord: ", t.Description)
}

func (t TimeRecord) InProgress() bool {
	return t.End == nil
}

// BeforeCreate is a GORM hook
func (*TimeRecord) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryTimeRecords(args ...interface{}) []model.TimeRecord {
	records := []*TimeRecord{}

	if len(args) > 0 {
		str, ok := args[0].(string)
		if ok {
			args[0] = common.CamelToSnake(str)
		}
	}

	db.Order("Date DESC, Start DESC").Find(&records, args...)

	ret := make([]model.TimeRecord, len(records))
	for k, v := range records {
		ret[k] = v
	}

	return ret
}
