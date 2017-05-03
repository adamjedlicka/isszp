package model

type Model interface {
	GetID() string
	FillByID(string) error
	Save() error
	Delete() error
	String() string
}

type Task interface {
	Model

	GetName() string
	SetName(string)
	GetDescription() string
	SetDescription(string)
	GetState() TaskState
	SetState(TaskState)
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
	NewTask    func() Task
	QueryTasks func(...interface{}) []Task
)

type TaskState string

const (
	TaskStateFree      TaskState = "free"
	TaskStateActive    TaskState = "active"
	TaskStateRevission TaskState = "revision"
	TaskStateSuccess   TaskState = "success"
	TaskStateFail      TaskState = "fail"
)

type Project interface {
	Model

	GetName() string
	SetName(string)
	GetCode() string
	SetCode(string)
	GetDescription() string
	SetDescription(string)

	GetMaintainer() User
	SetMaintainer(User)

	GetFirm() Firm
	SetFirm(Firm)
}

var (
	NewProject    func() Project
	QueryProjects func(...interface{}) []Project
)

type User interface {
	Model

	FillByUserName(string) error

	GetUserName() string
	SetUserName(string)
	SetPassword(string)
	CheckPassword(string) bool
	GetFirstName() string
	SetFirstName(string)
	GetLastName() string
	SetLastName(string)

	GetPermission() Permission
	SetPermission(Permission)
}

var (
	NewUser    func() User
	QueryUsers func(...interface{}) []User
)

type Permission interface {
	Model

	FillByName(string) error

	GetName() string
	SetName(string)
	GetIsAdmin() bool
	SetIsAdmin(bool)
}

var (
	NewPermission    func() Permission
	QueryPermissions func(...interface{}) []Permission
)

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
	NewFirm    func() Firm
	QueryFirms func(...interface{}) []Firm
)

type TimeRecord interface {
	Model

	GetDescription() string
	SetDescription(string)
	GetDate() string
	SetDate(string)
	GetStart() string
	SetStart(string)
}

var (
	NewTimeRecord    func() TimeRecord
	QueryTimeRecords func(...interface{}) []TimeRecord
)

type File interface {
	Model

	GetName() string
	SetName(string)
	GetUploadDateTime() string
	SetUploadDateTime(string)
	GetData() string
	SetData(string)
}

var (
	NewFile    func() File
	QueryFiles func(...interface{}) []File
)

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
	NewComment    func() Comment
	QueryComments func(...interface{}) []Comment
)
