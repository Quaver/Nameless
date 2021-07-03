package processors

import (
	"github.com/Swan/Nameless/src/config"
	"testing"
)

func TestCompileQuaverTools(t *testing.T) {
	config.InitializeConfig("../../")
	CompileQuaverTools()
}
