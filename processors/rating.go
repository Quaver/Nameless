package processors

import (
	"encoding/json"
	"fmt"
	"github.com/Swan/Nameless/config"
	"os/exec"
)

type RatingProcessor struct {
	Version string  `json:"Version"`
	Rating  float64 `json:"Rating"`
}

// CalcPerformance Uses Quaver.Tools to calculate the performance rating of a score
func CalcPerformance(diff float64, acc float32, failed bool) (RatingProcessor, error) {
	diffStr := fmt.Sprintf("%f", diff)
	accStr := fmt.Sprintf("%f", acc)
	output, err := exec.Command(config.Data.QuaverToolsPath, "-calcrating", diffStr, accStr).Output()

	if err != nil {
		return RatingProcessor{}, err
	}

	var r RatingProcessor

	err = json.Unmarshal(output, &r)

	if err != nil {
		return RatingProcessor{}, err
	}

	// Failures will always result in a zero rating
	if failed {
		r.Rating = 0
	}

	return r, nil
}
