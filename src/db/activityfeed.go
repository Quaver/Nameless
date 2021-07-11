package db

import "time"

type ActivityFeed int 

const (
	ActivityFeedRegistered ActivityFeed = iota
	ActivityFeedUploadedMapset
	ActivityFeedUpdatedMapset
	ActivityFeedRankedMapset
	ActivityFeedDeniedMapset
	ActivityFeedAchievedFirstPlace
	ActivityFeedLostFirstPlace
	ActivityFeedUnlockedAchievement
	ActivityFeedDonated
	ActivityFeedReceivedDonatorGift
)

// InsertActivityFeed Adds an activity feed log to the database
func InsertActivityFeed(userId int, feed ActivityFeed, value string, mapsetId int) error {
	query := "INSERT INTO activity_feed (user_id, type, timestamp, value, mapset_id) VALUES (?, ?, ?, ?, ?)"
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	
	_, err := SQL.Exec(query, userId, feed, timestamp, value, mapsetId)
	
	if err != nil {
		return err
	}
	
	return nil
}
