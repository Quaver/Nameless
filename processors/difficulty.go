package processors

import (
	"encoding/json"
	common2 "github.com/Swan/Nameless/common"
	"os/exec"
	"strconv"
)

type DifficultyProcessor struct {
	Metadata DifficultyProcessorMetadata `json:"Metadata"`
	Result   DifficultyProcessorResult   `json:"Difficulty"`
}

type DifficultyProcessorMetadata struct {
	Artist         string       `json:"Artist"`
	Title          string       `json:"Title"`
	DifficultyName string       `json:"DifficultyName"`
	Creator        string       `json:"Creator"`
	Mode           common2.Mode `json:"Mode"`
	Length         int          `json:"Length"`
	MapId          int          `json:"MapId"`
	MapSetId       int          `json:"MapSetId"`
	ObjectCount    int          `json:"ObjectCount"`
}

type DifficultyProcessorResult struct {
	OverallDifficulty float64 `json:"OverallDifficulty"`
	Version           string  `json:"Version"`
}

// CalcDifficulty Calculates the difficulty rating of a local .qua file
func CalcDifficulty(path string, mods common2.Mods) (DifficultyProcessor, error) {
	modsStr := strconv.Itoa(int(mods))
	output, err := exec.Command("dotnet", getQuaverToolsDllPath(), "-calcdiff", path, modsStr).Output()

	if err != nil {
		return DifficultyProcessor{}, err
	}

	var d DifficultyProcessor

	err = json.Unmarshal(output, &d)

	if err != nil {
		return DifficultyProcessor{}, err
	}

	return d, nil
}
