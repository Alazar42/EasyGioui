package core

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"easygioui"
)

func Run(appDir string, devMode bool) error {
	// Load the .easy UI file from the app directory
	uiPath := filepath.Join(appDir, "app", "ui", "app.easy")
	ui := easygioui.Load(uiPath)
	if ui == nil {
		return fmt.Errorf("failed to load UI from %s", uiPath)
	}

	// TODO: Implement proper Gio window setup and event loop
	// This would require integrating with Gio's window.Window
	// and running the application in Gio's event loop
	fmt.Println("EasyGioUI app loaded (window integration pending)")
	return nil
}

func Build(appDir, output string) error {
	cmd := exec.Command("go", "build", "-o", output, filepath.Join(appDir, "cmd", "easygio"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build failed: %w: %s", err, string(out))
	}
	return nil
}

