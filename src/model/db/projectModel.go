package db

import (
	"fmt"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewProject = NewProject
	model.QueryProjects = QueryProjects
}

type Project struct {
	ID string `gorm:"primary_key"`

	Name        string
	Code        string
	Description string
	StartDate   string
	PlanEndDate *string
	EndDate     *string

	MaintainerID string
	FirmID       string

	DeletedAt *string
}

func NewProject() model.Project {
	p := new(Project)

	return p
}

func (p *Project) FillByID(id string) error {
	return db.First(p, "id = ?", id).Error
}

func (p *Project) Save() error {
	return db.Save(p).Error
}

func (p *Project) Delete() error {
	return db.Delete(p).Error
}

func (p Project) GetID() string          { return p.ID }
func (p Project) GetName() string        { return p.Name }
func (p Project) GetCode() string        { return p.Code }
func (p Project) GetDescription() string { return p.Description }

func (p *Project) SetName(val string)        { p.Name = val }
func (p *Project) SetCode(val string)        { p.Code = val }
func (p *Project) SetDescription(val string) { p.Description = val }
func (p *Project) SetStartDate(val string)   { p.StartDate = val }

func (p Project) GetMaintainer() model.User {
	u := model.NewUser()
	u.FillByID(p.MaintainerID)
	return u
}

func (p Project) GetFirm() model.Firm {
	f := model.NewFirm()
	f.FillByID(p.FirmID)
	return f
}

func (p *Project) SetMaintainer(val model.User) { p.MaintainerID = val.GetID() }
func (p *Project) SetFirm(val model.Firm)       { p.FirmID = val.GetID() }

func (p Project) String() string {
	return fmt.Sprint("Project: ", p.Name)
}

// BeforeCreate is a GORM hook
func (*Project) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryProjects(args ...interface{}) []model.Project {
	projects := []*Project{}

	if len(args) > 0 {
		str, ok := args[0].(string)
		if ok {
			args[0] = common.CamelToSnake(str)
		}
	}

	db.Find(&projects, args...)

	ret := make([]model.Project, len(projects))
	for k, v := range projects {
		ret[k] = v
	}

	return ret
}
