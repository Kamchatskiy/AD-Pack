package main

import (
	"bufio"
	"fmt"
	"os"

	"../../internal/colors"
	"../../internal/logo"
	"../../internal/setup"
)

func main() {
	ctx, err := setup.Context()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logo.Print(ctx)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		logo.Print(ctx)
		command := scanner.Text()

		switch command {
		case "sploit", "spl":

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Incorrect command", colors.Colors.Warning)
		}

	}

}
