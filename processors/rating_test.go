package processors

import (
	"fmt"
	config2 "github.com/Swan/Nameless/config"
	"testing"
)

func TestCalculateRating(t *testing.T) {
	config2.InitializeConfig("../")
	r, err := CalcPerformance(30.25, 100, false)

	if err != nil {
		t.Fatal(err.Error())
	}

	const expectedRating float64 = 34.14828715972686

	if r.Rating != expectedRating {
		t.Fatal(fmt.Sprintf("expected rating %v", expectedRating))
	}
}
