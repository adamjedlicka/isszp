package controller

type Notification struct {
	ToUserID string
	Header   string
	Message  string
	Address  string
}

type Option func(*Notification)

var notifications map[string][]Notification

func init() {
	notifications = make(map[string][]Notification)
}

func SendNotification(userID, message string, opts ...Option) {
	notify := Notification{
		ToUserID: userID,
		Header:   "ISSZP",
		Message:  message,
		Address:  "",
	}

	for _, opt := range opts {
		opt(&notify)
	}

	notifications[userID] = append(notifications[userID], notify)
}

func GetNotifications(userID string) []Notification {
	notify, ok := notifications[userID]
	if !ok {
		notify = []Notification{}
	}

	notifications[userID] = []Notification{}

	return notify
}

func NotifyWithAddress(address string) Option {
	return func(n *Notification) {
		n.Address = address
	}
}

func NotifyWithHeader(header string) Option {
	return func(n *Notification) {
		n.Header = header
	}
}
