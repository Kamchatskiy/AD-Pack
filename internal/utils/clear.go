package utils

import (
	"os"
	"os/exec"
)

func ClearSceen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	PrintLogo()
}
