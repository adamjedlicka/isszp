package db

import (
	"fmt"

	"isszp/src/model"

	"github.com/jinzhu/gorm"
)

func init() {
	model.NewUser = NewUser
	model.QueryUsers = QueryUsers
}

type User struct {
	ID string `gorm:"primary_key"`

	UserName  string
	Password  string
	FirstName string
	LastName  string

	PermissionID string
	Permission   Permission

	DeletedAt *string
}

func NewUser() model.User {
	u := new(User)

	return u
}

func (u *User) FillByID(id string) error {
	return db.First(u, "id = ?", id).Error
}

func (u *User) FillByUserName(userName string) error {
	return db.First(u, "user_name = ?", userName).Error
}

func (u *User) Save() error {
	return db.Save(u).Error
}

func (u *User) Delete() error {
	return db.Delete(u).Error
}

func (u User) GetID() string        { return u.ID }
func (u User) GetUserName() string  { return u.UserName }
func (u User) GetFirstName() string { return u.FirstName }
func (u User) GetLastName() string  { return u.LastName }

func (u *User) SetUserName(val string)  { u.UserName = val }
func (u *User) SetPassword(val string)  { u.Password = val }
func (u *User) SetFirstName(val string) { u.FirstName = val }
func (u *User) SetLastName(val string)  { u.LastName = val }

func (u *User) CheckPassword(p string) bool {
	return p == u.Password
}

func (u *User) GetPermission() model.Permission {
	db.Model(u).Related(&u.Permission)
	return &u.Permission
}

func (u *User) SetPermission(p model.Permission) { u.PermissionID = p.GetID() }

func (u User) String() string {
	return fmt.Sprint("User: ", u.UserName)
}

// BeforeCreate is a GORM hook
func (*User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", NewUUID())
}

func QueryUsers(args ...interface{}) []model.User {
	users := []*User{}

	db.Find(&users)

	ret := make([]model.User, len(users))
	for k, v := range users {
		ret[k] = v
	}

	return ret
}
