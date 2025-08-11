package model

type Post struct {
	PostId        string  `json:"post_id"`
	Username      string  `json:"username"`
	ReactionScore float64 `json:"reaction_score"`
	Score         float64 `json:"score"`
}
