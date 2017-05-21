package controller

// Notification holds data that are neccessary to send notification between two clients
// each notification when sent can be received only once
type Notification struct {
	ToUserID string
	Header   string
	Message  string
	Address  string
}

// Option is data types used during creatiuon of new notification
type Option func(*Notification)

var notifications map[string][]Notification

func init() {
	notifications = make(map[string][]Notification)
}

// SendNotification sends notification to user defined by userID with text message
// it can be configured with number of options
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

// GetNotifications returns an array of all notification sent to user defined by userID
// and marks all these notifications as sent which prevents them from being received a second time
func GetNotifications(userID string) []Notification {
	notify, ok := notifications[userID]
	if !ok {
		notify = []Notification{}
	}

	notifications[userID] = []Notification{}

	return notify
}

// NotifyWithAddress is an option that sets an URL to be opened when clicked on the notification
func NotifyWithAddress(address string) Option {
	return func(n *Notification) {
		n.Address = address
	}
}

// NotifyWithHeader is an option that sets title of the notification
func NotifyWithHeader(header string) Option {
	return func(n *Notification) {
		n.Header = header
	}
}
