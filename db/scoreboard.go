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
	keys, err := Redis.Keys(RedisCtx, fmt.Sprintf("quaver:scores:%v_*", m.Id)).Result()

	if err != nil {
		return err
	}

	for _, key := range keys {
		str, err := Redis.Get(RedisCtx, key).Result()

		if err != nil {
			// The key does not exist anymore.
			if err == redis.Nil {
				return nil
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

	return nil
}
