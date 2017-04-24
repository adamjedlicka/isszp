package db

import (
	"fmt"

	"isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewTask = NewTask
	model.QueryTasks = QueryTasks
}

type Task struct {
	ID string `gorm:"primary_key"`

	Name        string
	Description string
	StartDate   string
	PlanEndDate *string
	EndDate     *string

	MaintainerID string
	Maintainer   User
	WorkerID     *string
	Worker       User
	ProjectID    string
	Project      Project

	DeletedAt *string
}

func NewTask() model.Task {
	t := new(Task)

	return t
}

func (t *Task) FillByID(id string) error {
	return db.First(t, "id = ?", id).Error
}

func (t *Task) Save() error {
	return db.Save(t).Error
}

func (t *Task) Delete() error {
	return db.Delete(t).Error
}

func (t Task) GetID() string           { return t.ID }
func (t Task) GetName() string         { return t.Name }
func (t Task) GetDescription() string  { return t.Description }
func (t Task) GetStartDate() string    { return t.StartDate }
func (t Task) GetPlanEndDate() *string { return t.PlanEndDate }
func (t Task) GetEndDate() *string     { return t.EndDate }

func (t *Task) SetName(val string)         { t.Name = val }
func (t *Task) SetDescription(val string)  { t.Description = val }
func (t *Task) SetStartDate(val string)    { t.StartDate = val }
func (t *Task) SetPlanEndDate(val *string) { t.PlanEndDate = val }
func (t *Task) SetEndDate(val *string)     { t.EndDate = val }

func (t *Task) GetMaintainer() model.User {
	db.Model(t).Related(&t.Maintainer, "Maintainer")
	return &t.Maintainer
}

func (t *Task) SetMaintainer(val model.User) { t.MaintainerID = val.GetID() }

func (t *Task) GetWorker() model.User {
	db.Model(t).Related(&t.Worker, "Worker")
	return &t.Worker
}

func (t *Task) SetWorker(val model.User) {
	if val == nil {
		t.WorkerID = nil
	} else {
		id := val.GetID()
		t.WorkerID = &id
	}
}

func (t *Task) GetProject() model.Project {
	db.Model(t).Related(&t.Project)
	return &t.Project
}

func (t *Task) SetProject(val model.Project) { t.ProjectID = val.GetID() }

func (t Task) String() string {
	return fmt.Sprint("Project: ", t.Name, ", MaintainerID: ", t.MaintainerID)
}

// BeforeCreate is a GORM hook
func (*Task) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryTasks(args ...interface{}) []model.Task {
	tasks := []*Task{}

	db.Find(&tasks)

	ret := make([]model.Task, len(tasks))
	for k, v := range tasks {
		ret[k] = v
	}

	return ret
}
