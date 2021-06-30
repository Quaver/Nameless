package processors

import (
	"fmt"
	"testing"
)

func TestCalculateRating(t *testing.T) {
	err := CompileQuaverTools()
	
	if err != nil {
		t.Fatal(err.Error())
	}
	
	r, err := CalculatePerformanceRating(30.25, 100)
	
	if err != nil {
		t.Fatal(err.Error())
	}
	
	const expectedRating float64 = 34.14828715972686
	
	if r.Rating !=  expectedRating {
		t.Fatal(fmt.Sprintf("expected rating %v", expectedRating))
	}
}