package processors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Swan/Nameless/common"
	"github.com/Swan/Nameless/config"
	"os/exec"
	"strconv"
)

type DifficultyProcessor struct {
	Metadata DifficultyProcessorMetadata `json:"Metadata"`
	Result   DifficultyProcessorResult   `json:"Difficulty"`
}

type DifficultyProcessorMetadata struct {
	Artist         string      `json:"Artist"`
	Title          string      `json:"Title"`
	DifficultyName string      `json:"DifficultyName"`
	Creator        string      `json:"Creator"`
	Mode           common.Mode `json:"Mode"`
	Length         int         `json:"Length"`
	MapId          int         `json:"MapId"`
	MapSetId       int         `json:"MapSetId"`
	ObjectCount    int         `json:"ObjectCount"`
}

type DifficultyProcessorResult struct {
	OverallDifficulty float64 `json:"OverallDifficulty"`
	Version           string  `json:"Version"`
}

// CalcDifficulty Calculates the difficulty rating of a local .qua file
func CalcDifficulty(path string, mods common.Mods) (DifficultyProcessor, error) {
	modsStr := strconv.FormatInt(int64(mods), 10)
	cmd := exec.Command(config.Data.QuaverToolsPath, "-calcdiff", path, modsStr)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return DifficultyProcessor{}, fmt.Errorf("%v\n\n```%v```", err, stderr.String())
	}

	var d DifficultyProcessor

	err = json.Unmarshal(out.Bytes(), &d)

	if err != nil {
		return DifficultyProcessor{}, err
	}

	return d, nil
}
