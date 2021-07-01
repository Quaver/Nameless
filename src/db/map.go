package db

import (
	"database/sql"
	"fmt"
	"github.com/Swan/Nameless/src/common"
)

type Map struct {
	Id               int
	MapsetId         int
	MD5              string
	AlternativeMD5   sql.NullString
	CreatorId        int
	CreatorUsername  string
	GameMode         common.Mode
	RankedStatus     common.RankedStatus
	Length           int32
	DifficultyRating float64
}

// GetMapByMD5 Fetches a map in the database by its MD5 hash
func GetMapByMD5(md5 string) (Map, error) {
	query := "SELECT " +
		"id, mapset_id, md5, alternative_md5, creator_id, creator_username, game_mode, " +
		"ranked_status, length " +
		"FROM maps WHERE md5 = ? OR alternative_md5 = ? LIMIT 1"

	var m Map

	err := SQL.QueryRow(query, md5, md5).Scan(
		&m.Id, &m.MapsetId, &m.MD5, &m.AlternativeMD5, &m.CreatorId, &m.CreatorUsername, &m.GameMode,
		&m.RankedStatus, &m.Length)

	if err != nil {
		return Map{}, err
	}

	return m, nil
}

// GetMapById Fetches a map in the database by its id
func GetMapById(id int32) (Map, error) {
	query := "SELECT " +
		"id, mapset_id, md5, alternative_md5, creator_id, creator_username, game_mode, " +
		"ranked_status, length " +
		"FROM maps WHERE id = ? LIMIT 1"

	var m Map

	err := SQL.QueryRow(query, id).Scan(
		&m.Id, &m.MapsetId, &m.MD5, &m.AlternativeMD5, &m.CreatorId, &m.CreatorUsername, &m.GameMode,
		&m.RankedStatus, &m.Length)

	if err != nil {
		return Map{}, err
	}

	return m, nil
}

// IncrementMapPlayCount Increments the play count & fail count of the map in the db & ES
// TODO: Update ES
func IncrementMapPlayCount(id int, failed bool) error {
	failQueryStr := ""
	
	if failed {
		failQueryStr = ", fail_count = fail_count + 1"
	}
	
	query := fmt.Sprintf("UPDATE maps SET play_count = play_count + 1%v WHERE id = ?", failQueryStr)
	
	_, err := SQL.Exec(query, id)
	
	if err != nil {
		return err
	}
	
	return nil
}
