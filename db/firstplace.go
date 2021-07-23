package db

type FirstPlaceScore struct {
	MD5               string
	UserId            int
	ScoreId           int
	PerformanceRating float64
}

// GetFirstPlaceScore Retrieves a first place score for a given map
func GetFirstPlaceScore(md5 string) (FirstPlaceScore, error) {
	query := "SELECT * FROM scores_first_place WHERE md5 = ? LIMIT 1"

	var score FirstPlaceScore

	err := SQL.QueryRow(query, md5).Scan(
		&score.MD5, &score.UserId, &score.ScoreId, &score.PerformanceRating)

	if err != nil {
		return FirstPlaceScore{}, err
	}

	return score, nil
}

// Insert Inserts this first place score into the database
func (s *FirstPlaceScore) Insert() error {
	query := "INSERT INTO scores_first_place VALUES (?, ?, ?, ?)"

	_, err := SQL.Exec(query, s.MD5, s.UserId, s.ScoreId, s.PerformanceRating)

	if err != nil {
		return err
	}

	return nil
}

// Update Updates the first place score in the database
func (s *FirstPlaceScore) Update() error {
	query := "UPDATE scores_first_place SET user_id = ?, performance_rating = ?, score_id = ? WHERE md5 = ?"

	_, err := SQL.Exec(query, s.UserId, s.PerformanceRating, s.ScoreId, s.MD5)

	if err != nil {
		return err
	}

	return nil
}

func NewFirstPlaceScore(md5 string, userId int, scoreId int, rating float64) FirstPlaceScore {
	return FirstPlaceScore{
		MD5:               md5,
		UserId:            userId,
		ScoreId:           scoreId,
		PerformanceRating: rating,
	}
}
