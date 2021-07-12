package achievements

import (
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/config"
	"github.com/Swan/Nameless/src/db"
	"testing"
)

func TestInitializeAchievement(t *testing.T) {
	config.InitializeConfig("../../")
	db.InitializeSQL()
}

func TestBabySteps(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	a := NewAchievementBabySteps()
	ok, err := a.Check(&user, &db.Score{ Failed: true }, &stats)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestAbsolutelyMarvelous(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementAbsolutelyMarvelous()
	ok, err := a.Check(&user, &db.Score{ Failed: true }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestCloseAchievement(t *testing.T) {
	db.CloseSQLConnection()
}

func getUser(id int, mode common.Mode) (db.User, db.UserStats, error) {
	user, err := db.GetUserById(id)
	
	if err != nil {
		return db.User{}, db.UserStats{}, err
	}
	
	user.CheckedPreviousAchievements = false
	stats, err := db.GetUserStats(id, mode)

	if err != nil {
		return db.User{}, db.UserStats{}, err
	}
	
	return user, stats, nil
}