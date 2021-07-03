package processors

import (
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/config"
	"testing"
)

func TestCalculateDifficulty(t *testing.T) {
	config.InitializeConfig("../../")

	const qua string = ""

	// TODO: Needs test file
	if qua == "" {
		return
	}

	_, err := CalcDifficulty(qua, common.ModMirror|common.ModSpeed125X)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
