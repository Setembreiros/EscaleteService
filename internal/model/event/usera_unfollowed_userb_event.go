package event

type UserAUnfollowedUserBEvent struct {
	Follower string `json:"followerId"`
	Followee string `json:"followeeId"`
}

var UserAUnfollowedUserBEventName = "UserAUnfollowedUserBEvent"
