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

func TestCombolicious(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementCombolicious()
	ok, err := a.Check(&user, &db.Score{ Failed: true }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestPerfectionist(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementPerfectionist()
	ok, err := a.Check(&user, &db.Score{ Failed: true }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestKeptYouPlayingHuh(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementKeptYouPlayingHuh()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestHumbleBeginnings(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementHumbleBeginnings()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestSteppingUpTheLadder(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementSteppingUpTheLadder()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestWideningYourHorizons(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementWideningYourHorizons()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestReachingNewHeights(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementReachingNewHeights()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestOutOfThisWorld(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementOutOfThisWorld()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestArea51(t *testing.T) {
	user, stats, err := getUser(608, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementArea51()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestAlien(t *testing.T) {
	user, stats, err := getUser(608, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementAlien()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestExtraterrestrial(t *testing.T) {
	user, stats, err := getUser(36960, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementAExtraterrestrial()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestET(t *testing.T) {
	user, stats, err := getUser(36960, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementET()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestQuombo(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementQuombo()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestOneTwoMayweather(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementOneTwoMayweather()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestItsOver5000(t *testing.T) {
	user, stats, err := getUser(608, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementItsOver5000()
	ok, err := a.Check(&user, &db.Score{}, &stats)

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