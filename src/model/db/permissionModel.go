package db

import (
	"fmt"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewPermission = NewPermission
	model.QueryPermissions = QueryPermissions
}

type Permission struct {
	ID string `gorm:"primary_key"`

	Name              string
	IsAdmin           bool
	CanManageUsers    bool
	CanManageProjects bool

	DeletedAt *string
}

func NewPermission() model.Permission {
	p := new(Permission)

	return p
}

func (p *Permission) FillByID(id string) error {
	return db.First(p, "id = ?", id).Error
}

func (p *Permission) FillByName(name string) error {
	return db.First(p, "name = ?", name).Error
}

func (p *Permission) Save() error {
	return db.Save(p).Error
}

func (p *Permission) Delete() error {
	return db.Delete(p).Error
}

func (p Permission) GetID() string    { return p.ID }
func (p Permission) GetName() string  { return p.Name }
func (p Permission) GetIsAdmin() bool { return p.IsAdmin }

func (p *Permission) SetName(val string)  { p.Name = val }
func (p *Permission) SetIsAdmin(val bool) { p.IsAdmin = val }

func (p Permission) String() string {
	return fmt.Sprintf("Permission: %v", p.Name)
}

// BeforeCreate is a GORM hook
func (*Permission) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryPermissions(args ...interface{}) []model.Permission {
	permissions := []*Permission{}

	if len(args) > 0 {
		str, ok := args[0].(string)
		if ok {
			args[0] = common.CamelToSnake(str)
		}
	}

	db.Find(&permissions, args...)

	ret := make([]model.Permission, len(permissions))
	for k, v := range permissions {
		ret[k] = v
	}

	return ret
}
