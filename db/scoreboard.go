package db

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type scoreboardScore struct {
	PerformanceRating float64 `json:"performance_rating"`
}

// UpdateScoreboardCache Updates the redis cache for a particular score
func UpdateScoreboardCache(s *Score, m *Map) error {
	patterns := []string{
		fmt.Sprintf("quaver:scores:%v_*", m.Id),
		fmt.Sprintf("quaver:scoreboard:%v:*", m.MD5),
	}

	for _, pattern := range patterns {
		iter := Redis.Scan(RedisCtx, 0, pattern, 10000).Iterator()
		for iter.Next(RedisCtx) {
			key := iter.Val()
			str, err := Redis.Get(RedisCtx, key).Result()

			if err != nil {
				// The key does not exist anymore.
				if err == redis.Nil {
					continue
				}

				return err
			}

			var scores []scoreboardScore
			err = json.Unmarshal([]byte(str), &scores)

			if err != nil {
				return err
			}

			if len(scores) < 50 {
				err = Redis.Del(RedisCtx, key).Err()

				if err != nil {
					return err
				}

				continue
			}

			for _, score := range scores {
				if s.PerformanceRating > score.PerformanceRating {
					err = Redis.Del(RedisCtx, key).Err()

					if err != nil {
						return err
					}

					break
				}
			}
		}

		if err := iter.Err(); err != nil {
			return err
		}
	}

	return nil
}
