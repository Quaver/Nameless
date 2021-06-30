package processors

import "testing"

func TestCompileQuaverTools(t *testing.T) {
	err := CompileQuaverTools()

	if err != nil {
		t.Fatalf(err.Error())
	}
}
