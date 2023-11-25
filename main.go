package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	appPath := "app/app.go"

	cmd := exec.Command("sh", "-c", "find . | grep '\\.go' | entr -r go run "+appPath)
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:", err)
	}
}
