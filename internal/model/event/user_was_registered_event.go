package event

type UserWasRegisteredEvent struct {
	Username string `json:"username"`
}

var UserWasRegisteredEventName = "UserWasRegisteredEvent"
