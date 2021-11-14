package main

import (
	"database/sql"
	"fmt"
	"github.com/Swan/Nameless/common"
	"github.com/Swan/Nameless/config"
	"github.com/Swan/Nameless/db"
	"github.com/Swan/Nameless/processors"
	"github.com/Swan/Nameless/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

const diffCalcVersion string = "0.0.5"
const ratingVersion string = "0.0.2"

var maps = map[string]db.Map {}

type recalcScore struct {
	id int
	mods common.Mods
	accuracy float32
	mapMd5 string
	diffCalcVersion sql.NullString
	ratingVersion sql.NullString
}

func main() {
	log.SetFormatter(&log.TextFormatter{})
	cwd, _ := os.Getwd()
	config.InitializeConfig(cwd)
	processors.CompileQuaverTools()
	db.InitializeSQL()
	utils.InitializeAzure()
	
	users := fetchUsers()
	recalculateScores(users)
}

// Fetches a list of users in the database
func fetchUsers() []int {
	rows, err := db.SQL.Query("SELECT id FROM users")
	
	if err != nil {
		panic(err)
	}
	
	defer rows.Close()
	
	var users []int
	
	for rows.Next() {
		var id int = -1
		err := rows.Scan(&id)
		
		if err != nil {
			panic(err)
		}

		users = append(users, id)
	}
	
	return users
}

// Recalculate all scores for a list of users
func recalculateScores(users []int) {
	log.Infof("Recalculating scores for %v users", len(users))
	
	for index, val := range users {
		err, scores := fetchUserScores(val)

		if err != nil {
			log.Error(err)
			continue
		}

		var wg sync.WaitGroup
		
		for i := 0; i < len(scores); i++ {
			wg.Add(1)
			
			score := scores[i]

			var m db.Map
			
			if val, ok := maps[score.mapMd5]; ok {
				m = val
			} else {
				m, err = db.GetMapByMD5(score.mapMd5)

				if err != nil {
					wg.Done()
					continue
				}

				maps[score.mapMd5] = m
			}
			
			go func(score *recalcScore) {
				defer wg.Done()
				
				err := recalculateScore(index + 1, len(users), score, m)

				if err != nil {
					log.Errorf("%v %v", score.mapMd5, err)
				}
			}(&score)
		}
		
		wg.Wait()
	}

	log.Info("Finished recalculating scores!")
}

// Recalculates an individual score
func recalculateScore(user int, totalUsers int, score *recalcScore, m db.Map) error {
	path, err := utils.CacheQuaFile(m)
	
	if err != nil {
		log.Error(err)
		return err
	}
	
	d, err := processors.CalcDifficulty(path, score.mods)

	if err != nil {
		return err
	}

	p, err := processors.CalcPerformance(d.Result.OverallDifficulty, score.accuracy, false)
	
	if err != nil {
		return err
	}
	
	fmt.Printf("[%v/%v] [#%v] Recalculated Rating -> %v\n", user, totalUsers, score.id, p.Rating)
	return nil
}

// Fetches a given user's outdated scores
func fetchUserScores(id int) (error, []recalcScore) {
	query := "SELECT id, mods, accuracy, map_md5, difficulty_processor_version, performance_processor_version " + 
		"FROM scores WHERE user_id = ?"
	
	rows, err := db.SQL.Query(query, id)
	defer rows.Close()
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, []recalcScore{}
		}
		
		log.Error(err)
		return err, []recalcScore{}
	}
	
	var scores []recalcScore
	
	for rows.Next() {
		score := recalcScore{}
		err = rows.Scan(&score.id, &score.mods, &score.accuracy, &score.mapMd5, &score.diffCalcVersion, &score.ratingVersion)
		
		if err != nil {
			log.Error(err)
			continue
		}

		if score.diffCalcVersion.String == diffCalcVersion && score.ratingVersion.String == ratingVersion {
			continue
		}
		
		scores = append(scores, score)
	}
	
	return nil, scores
}
