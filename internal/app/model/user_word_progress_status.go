package model

type WordProgressStatus string

const (
	WordProgressNew      WordProgressStatus = "new"
	WordProgressLearning WordProgressStatus = "learning"
	WordProgressReview   WordProgressStatus = "review"
	WordProgressMastered WordProgressStatus = "mastered"
)
