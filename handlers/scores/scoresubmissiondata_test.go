package scores

import (
	common2 "github.com/Swan/Nameless/common"
	"testing"
)

// Returns sample score submission data used for testing
func makeSampleScoreSubmissionData() scoreSubmissionData {
	sampleData := scoreSubmissionData{
		ReplayData:           "dGVzdA==",
		ReplayMD5:            "b2dbeb695fa205804b1e5e72650ad2bb",
		GameId:               -1,
		ExecutingAssemblyMD5: "c1a961ee686488d541cd848d2ec1bb51",
		EntryAssemblyMD5:     "7cf500b5cff114ef31f8b88c24dff07d",
		MapMD5:               "1d78dc8ed51214e518b5114fe24490ae",
		MapMD5Replay:         "1d78dc8ed51214e518b5114fe24490ae",
		ReplayVersion:        "0.0.1",
		TimePlayEnded:        1624895393,
		AudioPlaybackRate:    1.0,
		ScrollSpeed:          300,
		GameMode:             common2.ModeKeys4,
		Mods:                 0,
		Failed:               false,
		TotalScore:           900000,
		Accuracy:             96.44,
		MaxCombo:             1510,
		CountMarv:            1279,
		CountPerf:            583,
		CountGreat:           98,
		CountGood:            12,
		CountOkay:            4,
		CountMiss:            6,
		ReplayFrameCount:     9000,
		PauseCount:           0,
		Username:             "TestUser123",
		ComboAtEnd:           17,
		HealthAtEnd:          51,
		TimePlayStart:        1624894000,
	}

	return sampleData
}

// Tests if the user has passed, but did not provide any replay data
func TestPassWithNoReplayData(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.Failed = false
	data.ReplayData = ""

	fails, ok := data.validate()

	if !ok && len(fails) == 1 && fails[0] == detectPassNoReplayData {
		return
	}

	t.Fatalf("expected failure for passing score w/ no replay data")
}

// Tests if the user failed, but still provided replay data
func TestFailWithReplayData(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.Failed = true
	data.ReplayData = "dGVzdA=="
	data.HealthAtEnd = 0

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectFailWithReplayData {
		return
	}

	t.Fatalf("expected detection for non-passing score w/ provided replay data")
}

// Tests if the user provided invalid base64 for replay data
func TestReplayDecodeFailure(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.Failed = false
	data.ReplayData = "XXXXXaGVsbG8="

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectReplayDecodeError {
		return
	}

	t.Fatalf("expected detection for replay decoding error")
}

// Tests if the user provided an invalid replay md5
func TestInvalidMD5Hashes(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.ReplayMD5 = "invalid-md5"
	data.ExecutingAssemblyMD5 = "invalid-md5-2"
	data.EntryAssemblyMD5 = "invalid-md5-3"
	data.MapMD5 = "invalid-md5-4"
	data.MapMD5Replay = "invalid-md5-5"

	detections, ok := data.validate()

	const expectedFailCount int = 5

	if !ok && len(detections) == expectedFailCount {
		return
	}

	t.Fatalf("expected %v expected md5 validation failure detections", expectedFailCount)
}

// Tests if the user provided an invalid score value
func TestInvalidTotalScore(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.TotalScore = 999999999

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectInvalidTotalScore {
		return
	}

	t.Fatalf("expected detection for invalid total score")
}

// Tests if the user provided a max score value while failing
func TestMaxScoreWithFailure(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.ReplayData = ""
	data.HealthAtEnd = 0
	data.TotalScore = 1000000
	data.Failed = true

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectMaxTotalScoreWithFailure {
		return
	}

	t.Fatalf("expected detection for max score + failure")
}

// Tests if the user provided an invalid audio playback rate
func TestInvalidAudioPlaybackRate(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.AudioPlaybackRate = 3.0

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectInvalidAudioPlaybackRate {
		return
	}

	t.Fatalf("expected detection for invalid playback rate")
}

// Tests if the user provided an invalid max combo for the amount of non-miss judgements they have
func TestInvalidMaxComboForJudgements(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	// This is an impossible combination of judgements + max combo
	data.MaxCombo = 9999999
	data.CountMarv = 100
	data.CountPerf = 15
	data.CountGreat = 10
	data.CountGood = 5
	data.CountOkay = 3
	data.CountMiss = 1
	data.ComboAtEnd = 0

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectInvalidMaxComboForJudgements {
		return
	}

	t.Fatalf("expected detection for invalid max combo for non-miss judgements")
}

// Tests if the user provided an invalid max combo vs. the combo they had at the end of the map
func TestInvalidMaxComboForComboAtEnd(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.CountMarv = 100
	data.CountPerf = 15
	data.CountGreat = 10
	data.CountGood = 5
	data.CountOkay = 3
	data.CountMiss = 1
	data.MaxCombo = 20
	data.ComboAtEnd = 21 // Combo at the end is higher than max combo, which is impossible

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectMaxComboAndEndMismatch {
		return
	}

	t.Fatalf("expected detection for invalid combo at end vs max combo")
}

// Tests if the user failed, but their health is not zero.
func TestFailureWithNonZeroHealth(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.Failed = true
	data.ReplayData = ""
	data.HealthAtEnd = 50

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectFailWithNonZeroHealth {
		return
	}

	t.Fatalf("expected detection for non-passing score w/ non-zero health")
}

// Tests if the user passed, but their health is zero.
func TestPassWithZeroHealth(t *testing.T) {
	data := makeSampleScoreSubmissionData()

	data.Failed = false
	data.HealthAtEnd = 0

	detections, ok := data.validate()

	if !ok && len(detections) == 1 && detections[0] == detectPassWithZeroHealth {
		return
	}

	t.Fatalf("expected detection for passing score w/ zero health")
}
