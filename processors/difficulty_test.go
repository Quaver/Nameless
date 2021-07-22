package processors

import (
	common "github.com/Swan/Nameless/common"
	config "github.com/Swan/Nameless/config"
	"testing"
)

func TestCalculateDifficulty(t *testing.T) {
	config.InitializeConfig("../")

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
