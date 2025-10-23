package event

type ReviewWasCreatedEvent struct {
	ReviewId uint64 `json:"reviewId"`
	PostId   string `json:"postId"`
	Username string `json:"username"`
	Rating   int    `json:"rating"`
}

var ReviewWasCreatedEventName = "ReviewWasCreatedEvent"
