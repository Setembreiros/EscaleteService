package model

type Post struct {
	PostId        string `json:"post_id"`
	Username      string `json:"username"`
	ReactionScore int    `json:"reaction_score"`
}
