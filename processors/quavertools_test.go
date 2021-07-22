package processors

import (
	config "github.com/Swan/Nameless/config"
	"testing"
)

func TestCompileQuaverTools(t *testing.T) {
	config.InitializeConfig("../")
	CompileQuaverTools()
}
