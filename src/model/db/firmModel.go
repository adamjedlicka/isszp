package db

import (
	"fmt"

	"isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewFirm = NewFirm
	model.QueryFirms = QueryFirms
}

type Firm struct {
	ID string `gorm:"primary_key"`

	Name        string
	Description string
	Email       string
	TelNumber   string

	DeletedAt *string
}

func NewFirm() model.Firm {
	f := new(Firm)

	return f
}

func (f *Firm) FillByID(id string) error {
	return db.First(f, "id = ?", id).Error
}

func (f *Firm) Save() error {
	return db.Save(f).Error
}

func (f *Firm) Delete() error {
	return db.Delete(f).Error
}

func (f Firm) GetID() string          { return f.ID }
func (f Firm) GetName() string        { return f.Name }
func (f Firm) GetDescription() string { return f.Description }
func (f Firm) GetEmail() string       { return f.Email }
func (f Firm) GetTelNumber() string   { return f.TelNumber }

func (f *Firm) SetName(val string)        { f.Name = val }
func (f *Firm) SetDescription(val string) { f.Description = val }
func (f *Firm) SetEmail(val string)       { f.Email = val }
func (f *Firm) SetTelNumber(val string)   { f.TelNumber = val }

func (f Firm) String() string {
	return fmt.Sprint("Firm: ", f.Name)
}

// BeforeCreate is a GORM hook
func (*Firm) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryFirms(args ...interface{}) []model.Firm {
	firms := []*Firm{}

	db.Find(&firms)

	ret := make([]model.Firm, len(firms))
	for k, v := range firms {
		ret[k] = v
	}

	return ret
}
