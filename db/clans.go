package db

type ClanScore struct {
	Id              int
	ClanId          int
	OverallRating   float64
	OverallAccuracy float64
}

func CalculateClanMapScore(clan int, md5 string) (ClanScore, error) {
	scores, err := GetClanScores(clan, md5, 10)

	if err != nil {
		return ClanScore{}, err
	}

	score := ClanScore{
		ClanId:          clan,
		OverallRating:   CalculateOverallRating(scores),
		OverallAccuracy: CalculateOverallAccuracy(scores),
	}

	return score, nil
}
