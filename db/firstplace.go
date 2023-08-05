package db

import (
	"encoding/json"
)

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

func PublishFirstPlaceScoreRedis(username, artist, title, difficultyName string) error {
	type redisFirstPlaceScore struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		Map struct {
			Artist         string `json:"artist"`
			Title          string `json:"title"`
			DifficultyName string `json:"difficulty_name"`
		} `json:"map"`
	}

	jsonData, err := json.Marshal(redisFirstPlaceScore{
		User: struct {
			Username string `json:"username"`
		}{
			Username: username,
		},
		Map: struct {
			Artist         string `json:"artist"`
			Title          string `json:"title"`
			DifficultyName string `json:"difficulty_name"`
		}{
			Artist:         artist,
			Title:          title,
			DifficultyName: difficultyName,
		},
	})

	if err != nil {
		return err
	}

	_, err = Redis.Publish(RedisCtx, "quaver:first_place_scores", string(jsonData)).Result()

	if err != nil {
		return err
	}

	return nil
}
