package achievements

import (
	common2 "github.com/Swan/Nameless/common"
	config2 "github.com/Swan/Nameless/config"
	db2 "github.com/Swan/Nameless/db"
	"testing"
)

func TestInitializeAchievement(t *testing.T) {
	config2.InitializeConfig("../../")
	db2.InitializeSQL()
	db2.InitializeRedis()
}

func TestGetUserUnlockedAchievements(t *testing.T) {
	achievements, err := GetUserUnlockedAchievements(1)

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if len(achievements) != 23 {
		t.Fatalf("expected 23 achievement count")
	}
}

func TestGetUserLockedAchievements(t *testing.T) {
	achievements, err := GetUserLockedAchievements(1)

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if len(achievements) != 8 {
		t.Fatalf("expected 8 achievement count")
	}
}

func TestCheckAchievementWithNewScore(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	user.CheckedPreviousAchievements = true
	achievements, err := CheckAchievementsWithNewScore(&user, &db2.Score{ Failed: true }, &stats) 
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if len(achievements) != 23 {
		t.Fatalf("expected 23 achievements unlocked")
	}
}

func TestBabySteps(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	a := NewAchievementBabySteps()
	ok, err := a.Check(&user, &db2.Score{ Failed: true }, &stats)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestAbsolutelyMarvelous(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementAbsolutelyMarvelous()
	ok, err := a.Check(&user, &db2.Score{ Failed: true }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestCombolicious(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementCombolicious()
	ok, err := a.Check(&user, &db2.Score{ Failed: true }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestPerfectionist(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementPerfectionist()
	ok, err := a.Check(&user, &db2.Score{ Failed: true }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestKeptYouPlayingHuh(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementKeptYouPlayingHuh()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestHumbleBeginnings(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementHumbleBeginnings()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestSteppingUpTheLadder(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementSteppingUpTheLadder()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestWideningYourHorizons(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementWideningYourHorizons()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestReachingNewHeights(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementReachingNewHeights()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestOutOfThisWorld(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementOutOfThisWorld()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestArea51(t *testing.T) {
	user, stats, err := getUser(608, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementArea51()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestAlien(t *testing.T) {
	user, stats, err := getUser(608, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementAlien()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestExtraterrestrial(t *testing.T) {
	user, stats, err := getUser(36960, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementAExtraterrestrial()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestET(t *testing.T) {
	user, stats, err := getUser(36960, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementET()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestQuombo(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementQuombo()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestOneTwoMayweather(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementOneTwoMayweather()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestItsOver5000(t *testing.T) {
	user, stats, err := getUser(608, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementItsOver5000()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func Test7500Deep(t *testing.T) {
	user, stats, err := getUser(608, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievement7500Deep()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestTenThousand(t *testing.T) {
	user, stats, err := getUser(608, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementTenThousand()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestBeginnersLuck(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementBeginnersLuck()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestItsGettingHarder(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementItsGettingHarder()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestGoingInsane(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementGoingInsane()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestYoureAnExpert(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementYoureAnExpert()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestPieceOfCake(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementPieceOfCake()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestFailureIsAnOption(t *testing.T) {
	user, stats, err := getUser(170, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementFailureIsAnOption()
	ok, err := a.Check(&user, &db2.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestApproachingTheBlueZenith(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementApproachingTheBlueZenith()
	ok, err := a.Check(&user, &db2.Score{ Mode: common2.ModeKeys4}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestClickTheArrows(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementClickTheArrows()
	ok, err := a.Check(&user, &db2.Score{ Mode: common2.ModeKeys4}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestFingerBreaker(t *testing.T) {
	user, stats, err := getUser(5, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementFingerBreaker()
	ok, err := a.Check(&user, &db2.Score{ Mode: common2.ModeKeys4}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestSlowlyButSurely(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementSlowlyButSurely()
	ok, err := a.Check(&user, &db2.Score{ Mode: common2.ModeKeys4}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestHeWasNumberOne(t *testing.T) {
	user, stats, err := getUser(1, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementHeWasNumberOne()
	ok, err := a.Check(&user, &db2.Score{ Mode: common2.ModeKeys4}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestStarvelous(t *testing.T) {
	user, stats, err := getUser(608, common2.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementStarvelous()
	ok, err := a.Check(&user, &db2.Score{ Mode: common2.ModeKeys4, CountMarv: 1 }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestCloseAchievement(t *testing.T) {
	db2.CloseSQLConnection()
}

func getUser(id int, mode common2.Mode) (db2.User, db2.UserStats, error) {
	user, err := db2.GetUserById(id)
	
	if err != nil {
		return db2.User{}, db2.UserStats{}, err
	}
	
	user.CheckedPreviousAchievements = false
	stats, err := db2.GetUserStats(id, mode)

	if err != nil {
		return db2.User{}, db2.UserStats{}, err
	}
	
	return user, stats, nil
}