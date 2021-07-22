package processors

import (
	config2 "github.com/Swan/Nameless/config"
	"testing"
)

func TestCompileQuaverTools(t *testing.T) {
	config2.InitializeConfig("../")
	CompileQuaverTools()
}
