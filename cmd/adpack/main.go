package main

import (
	"../../internal/setup"
	"../../internal/utils"
)

func main() {
	ctx, _ := setup.ContextSetup()

	utils.PrintLogo()
	utils.ClearSceen()
	if ctx.Value("interface") == "cli" {
		for {
			CLI(ctx)
		}
	} else if ctx.Value("interface") == "ui" {
		for {
			UI(ctx)
		}
	}

}
