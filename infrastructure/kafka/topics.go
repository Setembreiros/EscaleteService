package kafka

import "escalateservice/internal/model/event"

func getTopics() []string {
	return []string{
		event.UserWasRegisteredEventName,
		event.PostWasCreatedEventName,
		event.ReviewWasCreatedEventName,
		event.UserLikedPostEventName,
	}
}
