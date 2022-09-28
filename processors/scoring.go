package processors

import (
	"github.com/Swan/Nameless/common"
	"github.com/Swan/Nameless/db"
	"math"
)

// CalculateAccuracyFromJudgements Calculates the accuracy from the given score's judgements
func CalculateAccuracyFromJudgements(score db.Score) float64 {
	var acc float64 = 0
	
	// Since its being used in multiple places, Just to keep it shorter
	marvWeight := common.GetJudgementAccuracyWeight(common.JudgementMarv)
	
	acc += float64(score.CountMarv) * marvWeight
	acc += float64(score.CountPerf) * common.GetJudgementAccuracyWeight(common.JudgementPerf)
	acc += float64(score.CountGreat) * common.GetJudgementAccuracyWeight(common.JudgementGreat)
	acc += float64(score.CountGood) * common.GetJudgementAccuracyWeight(common.JudgementGood)
	acc += float64(score.CountOkay) * common.GetJudgementAccuracyWeight(common.JudgementOkay)
	acc += float64(score.CountMiss) * common.GetJudgementAccuracyWeight(common.JudgementMiss)
	
	totalCount := float64(score.CountMarv + score.CountPerf + score.CountGreat + score.CountGood + score.CountOkay + score.CountMiss)
	return math.Max(acc / (totalCount * marvWeight), 0) * marvWeight
}