package common

type RankedStatus int32

const (
	StatusNotSubmitted = iota
	StatusUnranked
	StatusRanked
	StatusDanCourse
)
