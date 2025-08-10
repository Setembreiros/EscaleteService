package event

type UserAFollowedUserBEvent struct {
	Follower string `json:"followerId"`
	Followee string `json:"followeeId"`
}

var UserAFollowedUserBEventName = "UserAFollowedUserBEvent"
