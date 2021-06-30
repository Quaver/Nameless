package processors

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type RatingProcessor struct {
	Version string  `json:"Version"`
	Rating  float64 `json:"Rating"`
}

// CalculatePerformanceRating Uses Quaver.Tools to calculate the performance rating of a score
func CalculatePerformanceRating(diff float64, acc float32) (RatingProcessor, error) {
	diffStr := fmt.Sprintf("%f", diff)
	accStr := fmt.Sprintf("%f", acc)
	output, err := exec.Command("dotnet", getQuaverToolsDllPath(), "-calcrating", diffStr, accStr).Output()

	if err != nil {
		return RatingProcessor{}, err
	}

	var r RatingProcessor

	err = json.Unmarshal(output, &r)

	if err != nil {
		return RatingProcessor{}, err
	}

	return r, nil
}
