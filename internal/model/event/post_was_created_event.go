package event

type Metadata struct {
	Username string `json:"username"`
}

type PostWasCreatedEvent struct {
	PostId   string   `json:"post_id"`
	Metadata Metadata `json:"metadata"`
}

var PostWasCreatedEventName = "PostWasCreatedEvent"
