package processors

import (
	"fmt"
	"os/exec"
)

const QuaverAPIPath string = "C:/Users/Swan/go/src/Nameless/Quaver.API"

// CompileQuaverTools Compiles Quaver.Tools, so that it can be used for rating and difficulty calculations
// Requires .NET Core 3.1 installation
func CompileQuaverTools() error {
	fmt.Println("Compiling Quaver.Tools...")
	
	cmd := exec.Command("dotnet", "build", "--configuration", "Release", QuaverAPIPath)
	err := cmd.Run()
	
	if err != nil {
		return err
	}
	
	fmt.Println("Successfully compiled Quaver.Tools!")
	return nil
}

// Returns the expected path of the Quaver.Tools.dll file
func getDllPath() string {
	return fmt.Sprintf("%v/Quaver.Tools/bin/Release/netcoreapp3.1/Quaver.Tools.dll", QuaverAPIPath)
}