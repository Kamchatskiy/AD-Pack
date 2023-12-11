package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"../../../internal/colors"
	"../../../internal/utils"
)

func CLI(ctx context.Context) {
	utils.PrintLogo()
	fmt.Println(colors.Colors.Secondary)

	scanner := ctx.Value("scanner").(*bufio.Scanner)
	command := scanner.Text()
	command = strings.ToLower(command)

	switch command {
	case "sploit", "spl":

	case "clear":
		utils.ClearSceen()
	case "exit", "leave":
		os.Exit(0)
	default:
		utils.WrongCMD()
	}
}
