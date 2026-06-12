package main

import (
	"fmt"

	"easygioui/sdk/easygio"
)

func main() {
	app := easygio.New()
	ui := app.Window("Go Only", app.VBox(
		app.Text("titleText", "Hello from Go-only DSL"),
		app.Button("loginBtn", "Login", "App.Login"),
	))

	tree := ui.BuildTree()
	if v, ok := tree.Get("title"); ok {
		fmt.Println("window title:", v)
	}
	fmt.Println("easygio example prepared")
}
