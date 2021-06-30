package processors

import (
	"github.com/Swan/Nameless/src/common"
	"testing"
)

func TestCalculateDifficulty(t *testing.T) {
	err := CompileQuaverTools()

	if err != nil {
		return
	}

	const qua string = ""

	// TODO: Needs test file
	if qua == "" {
		return
	}
	
	_, err = CalculateDifficulty(qua, common.ModMirror | common.ModSpeed125X)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
