package scores

import (
	"encoding/base64"
	"fmt"
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/db"
	"github.com/Swan/Nameless/src/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type scoreSubmissionData struct {
	ReplayData           string      `json:"replay_data"` // Base64 encoded replay data
	RawReplayData        []byte      // Raw & decoded replay data
	ReplayMD5            string      `json:"replay_md5"`
	GameId               int         `json:"game_id"`
	ExecutingAssemblyMD5 string      `json:"executing_assembly"`
	EntryAssemblyMD5     string      `json:"entry_assembly"`
	MapMD5               string      `json:"map_md5"`
	MapMD5Replay         string      `json:"map_md5_replay"`
	ReplayVersion        string      `json:"replay_version"`
	TimePlayEnded        int64       `json:"time_play_ended"`
	AudioPlaybackRate    float32     `json:"audio_playback_rate"`
	ScrollSpeed          int16       `json:"scroll_speed"`
	GameMode             common.Mode `json:"game_mode"`
	Mods                 common.Mods `json:"mods"`
	Failed               bool        `json:"failed"`
	TotalScore           int32       `json:"total_score"`
	Accuracy             float32     `json:"accuracy"`
	MaxCombo             int32       `json:"max_combo"`
	CountMarv            int32       `json:"count_marv"`
	CountPerf            int32       `json:"count_perf"`
	CountGreat           int32       `json:"count_great"`
	CountGood            int32       `json:"count_good"`
	CountOkay            int32       `json:"count_okay"`
	CountMiss            int32       `json:"count_miss"`
	ReplayFrameCount     int32       `json:"replay_frame_count"`
	PauseCount           int32       `json:"pause_count"`
	Username             string      `json:"username"`
	ComboAtEnd           int32       `json:"combo_at_end"`
	HealthAtEnd          float32     `json:"health_at_end"`
	TimePlayStart        int64       `json:"time_play_start"`
}

// Handles the parsing of incoming score submission scoreData.
func parseScoreSubmissionData(user *db.User, c *gin.Context) (scoreSubmissionData, error) {
	data := scoreSubmissionData{}

	err := c.BindJSON(&data)

	if err != nil {
		log.Errorf("Failed to deserialize score submission data - %v\n", err.Error())
		return scoreSubmissionData{}, err
	}

	detections, ok := data.validate()
	
	if !ok {
		dString := detectionListToString(detections)
		err = utils.SendAnticheatWebhook(user, nil, 0, false, dString)
		
		if err != nil {
			log.Errorf("Error sending anti-cheat log to discord - %v", err)
		}
		
		return scoreSubmissionData{}, fmt.Errorf("\n%v", dString)
	}

	return data, nil
}

// Validates incoming score submission data
func (data *scoreSubmissionData) validate() ([]string, bool) {
	detections := make([]string, 0)
	detections = data.validateReplayData(detections)
	detections = data.validateMD5Values(detections)
	detections = data.validateScoreData(detections)

	if len(detections) > 0 {
		return detections, false
	}

	return nil, true
}

// Makes sure we're getting replay data with only passing scores &
// makes sure the replay data passed is valid data
func (data *scoreSubmissionData) validateReplayData(d []string) []string {
	// Player stated that they passed, but did not provide replay data.
	if !data.Failed && data.ReplayData == "" {
		d = append(d, "Player passed but did not provide replay data")
	}

	// Player stated that they failed, but gave us replay data
	if data.Failed && data.ReplayData != "" {
		d = append(d, "Player failed but provided replay data")
	}

	var err error

	data.RawReplayData, err = base64.StdEncoding.DecodeString(data.ReplayData)

	if err != nil {
		d = append(d, "Failed to decode replay data")
	}

	return d
}

// Makes sure that values where an MD5 hash are expected are valid
func (data *scoreSubmissionData) validateMD5Values(d []string) []string {
	if !utils.IsValidMD5(data.ReplayMD5) {
		d = append(d, fmt.Sprintf("Replay MD5 was not a valid hash - %v", data.ReplayMD5))
	}

	if !utils.IsValidMD5(data.ExecutingAssemblyMD5) {
		d = append(d, fmt.Sprintf("Executing Assembly MD5 was not a valid hash - %v", data.ExecutingAssemblyMD5))
	}

	if !utils.IsValidMD5(data.EntryAssemblyMD5) {
		d = append(d, fmt.Sprintf("Entry assembly MD5 was not a valid hash - %v", data.EntryAssemblyMD5))
	}

	/*
	//  StepMania doesn't use MD5 hashes for their charts, but instead a "Chart Key", so skip this check.
	if !utils.IsValidMD5(data.MapMD5) {
		d = append(d, fmt.Sprintf("Map MD5 was not a valid hash - %v", data.MapMD5))
	}

	if !utils.IsValidMD5(data.MapMD5Replay) {
		d = append(d, fmt.Sprintf("Map Replay MD5 was not a valid hash - %v", data.MapMD5Replay))
	}
	*/

	return d
}

// Validates score-related data to make sure there are no discrepancies.
func (data *scoreSubmissionData) validateScoreData(d []string) []string {
	const maxScore int32 = 1000000

	if data.TotalScore > maxScore || data.TotalScore < 0 {
		d = append(d, fmt.Sprintf("Invalid total score provided - %v", data.TotalScore))
	}

	if data.TotalScore >= maxScore && data.Failed {
		d = append(d, "Max total score with a failing score provided")
	}

	if data.AudioPlaybackRate < 0.5 || data.AudioPlaybackRate > 2.0 {
		d = append(d, fmt.Sprintf("Invalid audio playback rate provided - %v", data.AudioPlaybackRate))
	}

	nonMiss := data.CountMarv + data.CountPerf + data.CountGreat + data.CountGood + data.CountOkay

	if data.MaxCombo > nonMiss {
		d = append(d, fmt.Sprintf("Invalid Max Combo for non-miss judgements: %v vs. %v", data.MaxCombo, nonMiss))
	}

	if data.ComboAtEnd > data.MaxCombo {
		d = append(d, fmt.Sprintf("Combo @ End > than Max Combo - %v vs %v", data.ComboAtEnd, data.MaxCombo))
	}

	if data.Failed && data.HealthAtEnd != 0 {
		d = append(d, "Player provided a failing score without zero health")
	}

	if !data.Failed && data.HealthAtEnd == 0 {
		d = append(d, "Player provided a passing score with zero health")
	}

	return d
}

// Returns if the game mode the user provided matches what is in the database
func (data *scoreSubmissionData) validateGameMode(m *db.Map) error {
	if data.GameMode != m.GameMode {
		return fmt.Errorf("provided game mode does not match DB: %v vs %v", data.GameMode, m.GameMode)
	}

	return nil
}

// Returns if the score has valid total score
func (data *scoreSubmissionData) isValidTotalScore() bool {
	return !(data.Failed && data.TotalScore == 0)
}

// Checks the score for anything suspicious about the score and returns it to Discord.
// Returns if the score is clean or not.
func (data *scoreSubmissionData) checkSuspiciousScore(h *Handler) bool {
	// Disregard failed scores
	if h.scoreData.Failed {
		return true
	}
	
	var detections []string
	
	// Detect extremely high ratio (potential autoplay)
	var ratio = data.getMARatio()
	
	if h.difficulty.Result.OverallDifficulty >= 10 && ratio >= 100 {
		d := fmt.Sprintf("Abnormally high ratio on score achieved: **%v** (Autoplay)", ratio)
		detections = append(detections, d)
	}
	
	// Nothing suspicious has been detected
	if len(detections) == 0 {
		return true
	}
	
	// Send webhook to discord
	err := utils.SendAnticheatWebhook(&h.user, &h.mapData, int(h.newScoreId), h.isPersonalBestScore(), 
		detectionListToString(detections))
	
	if err != nil {
		log.Errorf("Failed to send anticheat webhook to Discord - %v", err)
	}
	
	return false
}

// Returns the Marvelous:Perfect ratio
func (data *scoreSubmissionData) getMARatio() float32 {
	perfects := data.CountPerf

	if perfects == 0 {
		perfects = 1
	}
	
	return float32(data.CountMarv / perfects)
}

// Converts a list of detections to a readable string
func detectionListToString(d []string) string {
	str := ""
	
	for _, detection := range d {
		str += fmt.Sprintf("• %v\n", detection)
	}
	
	return str
}