package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"easygioui/internal/core"
	"easygioui/internal/devtools"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "run":
		fs := flag.NewFlagSet("run", flag.ExitOnError)
		appDir := fs.String("app", ".", "app directory")
		_ = fs.Parse(os.Args[2:])
		must(core.Run(*appDir, false))
	case "dev":
		fs := flag.NewFlagSet("dev", flag.ExitOnError)
		appDir := fs.String("app", ".", "app directory")
		_ = fs.Parse(os.Args[2:])
		must(core.Run(*appDir, true))
	case "build":
		fs := flag.NewFlagSet("build", flag.ExitOnError)
		appDir := fs.String("app", ".", "app directory")
		out := fs.String("o", "easygio-app", "output binary")
		_ = fs.Parse(os.Args[2:])
		must(core.Build(*appDir, *out))
	case "create":
		fs := flag.NewFlagSet("create", flag.ExitOnError)
		name := fs.String("name", "myapp", "application name")
		_ = fs.Parse(os.Args[2:])
		must(createSkeleton(*name))
	case "doctor":
		must(devtools.Doctor())
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("easygio commands: run, dev, build, create, doctor")
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func createSkeleton(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if err := os.MkdirAll(filepath.Join(name, "app"), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(name, "app", "ui.easy"), []byte(`Window {
    title: "EasyGioUI App"
    VBox {
        Text {
            id: titleText
            value: "Hello EasyGioUI"
        }
        Button {
            id: clickBtn
            text: "Click"
            onClick: App.OnClick
        }
    }
}
`), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(name, "main.go"), []byte(`package main

import "fmt"

func main() {
    fmt.Println("Run with: easygio dev -app .")
}
`), 0o644); err != nil {
		return err
	}
	fmt.Println("created", name)
	return nil
}

func _runGoBuild(target, output string) error {
	cmd := exec.Command("go", "build", "-o", output, target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
