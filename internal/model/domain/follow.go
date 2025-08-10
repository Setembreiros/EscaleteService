package model

type Follow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}
