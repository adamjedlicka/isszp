package db

import (
	"fmt"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
	"gitlab.fit.cvut.cz/isszp/isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewFile = NewFile
	model.QueryFiles = QueryFiles
}

type File struct {
	ID             string
	Name           string
	UploadDateTime string
	Data           string
}

func NewFile() model.File {
	f := new(File)

	return f
}

func (f *File) FillByID(id string) error {
	return db.First(f, "id = ?", id).Error
}

func (f *File) Save() error {
	return db.Save(f).Error
}

func (f *File) Delete() error {
	return db.Delete(f).Error
}

func (f File) GetID() string             { return f.ID }
func (f File) GetName() string           { return f.Name }
func (f File) GetUploadDateTime() string { return f.UploadDateTime }
func (f File) GetData() string           { return f.Data }

func (f *File) SetName(val string)           { f.Name = val }
func (f *File) SetUploadDateTime(val string) { f.UploadDateTime = val }
func (f *File) SetData(val string)           { f.Data = val }

func (f File) String() string {
	return fmt.Sprint("File: ", f.Name)
}

// BeforeCreate is a GORM hook
func (*File) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryFiles(args ...interface{}) []model.File {
	files := []*File{}

	if len(args) > 0 {
		str, ok := args[0].(string)
		if ok {
			args[0] = common.CamelToSnake(str)
		}
	}

	db.Find(&files, args...)

	ret := make([]model.File, len(files))
	for k, v := range files {
		ret[k] = v
	}

	return ret
}
