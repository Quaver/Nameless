package processors

import (
	common2 "github.com/Swan/Nameless/common"
	config2 "github.com/Swan/Nameless/config"
	"testing"
)

func TestCalculateDifficulty(t *testing.T) {
	config2.InitializeConfig("../")

	const qua string = ""

	// TODO: Needs test file
	if qua == "" {
		return
	}

	_, err := CalcDifficulty(qua, common2.ModMirror|common2.ModSpeed125X)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
