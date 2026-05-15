package main

import (
	"os"
	"os/exec"
)

func Color(input, text string) string {
	return text + input + "\033[0m"
}

func ClearScreen() {
	command := exec.Command("clear")
	command.Stdout = os.Stdout
	command.Run()
}
