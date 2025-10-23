package model

type Review struct {
	ReviewId uint64 `json:"reviewId"`
	PostId   string `json:"postId"`
	Reviewer string `json:"reviewer"`
	Rating   int    `json:"rating"`
}
