package setup

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"../colors"
	"../utils"
)

type CtxStr string

func ContextSetup() (context.Context, error) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, CtxStr("scanner"), bufio.NewScanner(os.Stdin))

	ctx = InterfaceSetup(ctx)

	return ctx, nil
}

func InterfaceSetup(ctx context.Context) context.Context {
	fmt.Print(colors.Colors.Error, "CLI ")
	fmt.Print(colors.Colors.Secondary, " / ")
	fmt.Print(colors.Colors.Warning, "UI ")
	fmt.Print(colors.Colors.Secondary, "mode?")

	scanner := ctx.Value("scanner").(*bufio.Scanner)
	for scanner.Scan() {
		choose := scanner.Text()
		choose = strings.ToLower(choose)
		utils.ClearSceen()
		switch choose {
		case "c", "cl", "cli":
			return context.WithValue(ctx, CtxStr("interface"), "cli")
		case "u", "ui":
			return context.WithValue(ctx, CtxStr("interface"), "ui")
		default:
			utils.WrongCMD()
		}
	}
	return ctx
}
