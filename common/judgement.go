package common

type Judgement int

const (
	JudgementMarv = iota
	JudgementPerf
	JudgementGreat
	JudgementGood
	JudgementOkay
	JudgementMiss
)

// GetJudgementAccuracyWeight Returns the accuracy weighting for a given judgement
func GetJudgementAccuracyWeight(j Judgement) float64 {
	switch j {
	case JudgementMarv:
		return 100
	case JudgementPerf:
		return 98.25
	case JudgementGreat:
		return 65
	case JudgementGood:
		return 25
	case JudgementOkay:
		return -100
	case JudgementMiss:
		return -50
	default:
		return 0
	}
}

