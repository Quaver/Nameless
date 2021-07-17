package achievements

import (
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/config"
	"github.com/Swan/Nameless/src/db"
	"testing"
)

func TestInitializeAchievement(t *testing.T) {
	config.InitializeConfig("../../../")
	db.InitializeSQL()
	db.InitializeRedis()
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

func Test7500Deep(t *testing.T) {
	user, stats, err := getUser(608, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievement7500Deep()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestTenThousand(t *testing.T) {
	user, stats, err := getUser(608, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementTenThousand()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestBeginnersLuck(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementBeginnersLuck()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestItsGettingHarder(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementItsGettingHarder()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestGoingInsane(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementGoingInsane()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestYoureAnExpert(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementYoureAnExpert()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestPieceOfCake(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementPieceOfCake()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestFailureIsAnOption(t *testing.T) {
	user, stats, err := getUser(170, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementFailureIsAnOption()
	ok, err := a.Check(&user, &db.Score{}, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestApproachingTheBlueZenith(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementApproachingTheBlueZenith()
	ok, err := a.Check(&user, &db.Score{ Mode: common.ModeKeys4 }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestClickTheArrows(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementClickTheArrows()
	ok, err := a.Check(&user, &db.Score{ Mode: common.ModeKeys4 }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestFingerBreaker(t *testing.T) {
	user, stats, err := getUser(5, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementFingerBreaker()
	ok, err := a.Check(&user, &db.Score{ Mode: common.ModeKeys4 }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestSlowlyButSurely(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementSlowlyButSurely()
	ok, err := a.Check(&user, &db.Score{ Mode: common.ModeKeys4 }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestHeWasNumberOne(t *testing.T) {
	user, stats, err := getUser(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementHeWasNumberOne()
	ok, err := a.Check(&user, &db.Score{ Mode: common.ModeKeys4 }, &stats)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !ok {
		t.Fatalf("achievement not unlocked")
	}
}

func TestStarvelous(t *testing.T) {
	user, stats, err := getUser(608, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	a := NewAchievementStarvelous()
	ok, err := a.Check(&user, &db.Score{ Mode: common.ModeKeys4, CountMarv: 1 }, &stats)

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