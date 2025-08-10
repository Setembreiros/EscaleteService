package model

type Post struct {
	PostId   string `json:"post_id"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}
