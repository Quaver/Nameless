package processors

import (
	"github.com/Swan/Nameless/config"
	"testing"
)

func TestCompileQuaverTools(t *testing.T) {
	config.InitializeConfig("../")
	CompileQuaverTools()
}
