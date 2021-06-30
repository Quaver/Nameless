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

	// TODO: Needs testing file
	const qua string = ""

	_, err = CalculateDifficulty(qua, common.ModMirror)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
