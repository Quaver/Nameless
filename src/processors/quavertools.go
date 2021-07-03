package processors

import (
	"fmt"
	"github.com/Swan/Nameless/src/config"
	"os/exec"
)

// CompileQuaverTools Compiles Quaver.Tools, so that it can be used for rating and difficulty calculations
// Requires .NET Core 3.1 installation
func CompileQuaverTools() {
	cmd := exec.Command("dotnet", "build", "--configuration", "Release", config.Data.QuaverAPIPath)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
	
	fmt.Println("Compiled Quaver.Tools!")
}

// Returns the expected path of the Quaver.Tools.dll file
func getQuaverToolsDllPath() string {
	return fmt.Sprintf("%v/Quaver.Tools/bin/Release/netcoreapp3.1/Quaver.Tools.dll", config.Data.QuaverAPIPath)
}
