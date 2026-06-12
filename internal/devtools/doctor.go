package devtools

import (
	"fmt"
	"os/exec"
)

func Doctor() error {
	goVersion, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("Go: %s", goVersion)
	fmt.Println("easygio: doctor passed basic checks")
	return nil
}
