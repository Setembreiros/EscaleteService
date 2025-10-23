package model

const (
	LikeScore        = 1
	SuperLikeScore   = 5
	Review1StarScore = -200
	Review2StarScore = -100
	Review3StarScore = 100
	Review4StarScore = 200
	Review5StarScore = 300
)

// ReactionTypes mapea os nomes das reaccións cos seus valores
var reactionTypes = map[string]int{
	"like":        LikeScore,
	"superlike":   SuperLikeScore,
	"review1star": Review1StarScore,
	"review2star": Review2StarScore,
	"review3star": Review3StarScore,
	"review4star": Review4StarScore,
	"review5star": Review5StarScore,
}

func GetScore(reactionType string) int {
	return reactionTypes[reactionType]
}
