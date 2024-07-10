package processors

import (
	"github.com/Swan/Nameless/config"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

// CompileQuaverTools Compiles Quaver.Tools, so that it can be used for rating and difficulty calculations
// Requires .NET Core 3.1 installation
func CompileQuaverTools() {
	log.Info("Compiling Quaver.Tools...")

	cmd := exec.Command("dotnet", "build", "--configuration", "Release", config.Data.QuaverToolsPath)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	log.Info("Successfully compiled Quaver.Tools!")
}
