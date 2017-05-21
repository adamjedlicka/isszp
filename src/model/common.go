// Package model describes common interface for communicating with model layer
// this metods allows for easily implementing new independent models that use their own sepcific method of storing data
package model

import "math"

// Model is common interface of every model. It must be composited by every other model interface
type Model interface {
	// GetID returns unique UUID of the model
	GetID() string

	// FillByID fils model with data based on supplied ID. If no record with this ID is found returns an error
	FillByID(string) error

	// Save saves the model. If ID == "" creates new record, othervise update existing one
	Save() error

	// Delete deletes the model. If possible hide the data instead of deleting
	Delete() error

	// String returns string representation of the models
	String() string
}

// Task describes interface for communicating with task model
type Task interface {
	Model

	GetName() string
	SetName(string)
	GetDescription() string
	SetDescription(string)
	GetState() string
	SetState(string)
	GetStartDate() string
	SetStartDate(string)
	GetPlanEndDate() *string
	SetPlanEndDate(*string)
	GetEndDate() *string
	SetEndDate(*string)

	GetMaintainer() User
	SetMaintainer(User)
	GetWorker() User
	SetWorker(User)
	GetProject() Project
	SetProject(Project)
}

var (
	// NewTask returns new empty task, that can be filled with .FillByID()
	NewTask func() Task

	// QueryTasks allows for querying tasks using Queryable interface
	QueryTasks func(...interface{}) []Task
)

// Project describes interface for communicating with project model
type Project interface {
	Model

	GetName() string
	SetName(string)
	GetCode() string
	SetCode(string)
	GetDescription() string
	SetDescription(string)
	GetStartDate() string
	SetStartDate(string)
	GetPlanEndDate() *string
	SetPlanEndDate(*string)
	GetEndDate() *string
	SetEndDate(*string)

	GetMaintainer() User
	SetMaintainer(User)

	GetFirm() Firm
	SetFirm(Firm)
}

var (
	// NewProject returns new empty project, that can be filled with .FillByID()
	NewProject func() Project

	// QueryProjects allows for querying projects using Queryable interface
	QueryProjects func(...interface{}) []Project
)

// User describes interface for communicating with user model
type User interface {
	Model

	FillByUserName(string) error

	GetUserName() string
	SetUserName(string)
	SetPassword(string)
	GetPassword() string
	CheckPassword(string) bool
	GetFirstName() string
	SetFirstName(string)
	GetLastName() string
	SetLastName(string)

	GetPermission() Permission
	SetPermission(Permission)
	AddPermission(Permission)
}

var (
	// NewUser returns new empty project, that can be filled with .FillByID()
	NewUser func() User

	// QueryUsers allows for querying users using Queryable interface
	QueryUsers func(...interface{}) []User
)

// Permission is data type used to store User permission
// Every permission is different bit in 64-bit integer which can be obtained with masks defined below
// example:
// p := 0b00101 - can manage projects and cam manage users => (p | CanManageProjects == CanManageProjects) : true
// p := 0b11111 - can do everything - is an admin => (p | IsAdmin == IsAdmin) : true
type Permission uint64

const (
	IsAdmin           Permission = math.MaxUint64 >> 1 // 2^63 - 9223372036854775808
	CanManageProjects Permission = 1 << iota
	CanManageTasks
	CanManageUsers
)

// Firm describes interface for communicating with firm model
type Firm interface {
	Model

	GetName() string
	SetName(string)
	GetDescription() string
	SetDescription(string)
	GetEmail() string
	SetEmail(string)
	GetTelNumber() string
	SetTelNumber(string)
}

var (
	// NewFirm returns new empty project, that can be filled with .FillByID()
	NewFirm func() Firm

	// QueryFirm allows for querying users using Queryable interface
	QueryFirms func(...interface{}) []Firm
)

// TimeRecord describes interface for communicating with firm model
type TimeRecord interface {
	Model

	GetDescription() string
	SetDescription(string)
	GetDate() string
	SetDate(string)
	GetStart() string
	SetStart(string)
	InProgress() bool

	SetStop(*string)
	SetTimeInMs(string)
	GetTimeInMs() string

	GetTask() Task
	SetTask(Task)
	GetUser() User
	SetUser(User)
}

var (
	// NewTimeRecord returns new empty project, that can be filled with .FillByID()
	NewTimeRecord func() TimeRecord

	// QueryTimeRecords allows for querying users using Queryable interface
	QueryTimeRecords func(...interface{}) []TimeRecord
)

// Comment describes interface for communicating with firm model
type Comment interface {
	Model

	GetText() string
	SetText(string)
	GetPostDateTime() string
	SetPostDateTime(string)

	GetUser() User
	SetUser(User)
	GetTask() Task
	SetTask(Task)
}

var (
	// NewComment returns new empty project, that can be filled with .FillByID()
	NewComment func() Comment

	// QueryComments allows for querying users using Queryable interface
	QueryComments func(...interface{}) []Comment
)
