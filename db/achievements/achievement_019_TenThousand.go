package achievements

import (
	"database/sql"
	db "github.com/Swan/Nameless/db"
)

type AchievementTenThousand Achievement

func NewAchievementTenThousand() AchievementTenThousand {
	return AchievementTenThousand{
		Id: 19,
		Name: "Ten Thousand",
		SteamAPIName: "TEN_THOUSAND",
	}
}

func (a AchievementTenThousand) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.MaxCombo >= 10000 && !score.Failed && !score.IsDonatorScore {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE max_combo >= 10000 AND is_donator_score = 0 AND failed = 0 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
