package utils

import (
	"fmt"

	"../colors"
)

func PrintLogo() {
	primaryColor := colors.Colors.Primary
	fmt.Println("                                              ", primaryColor)
	fmt.Println("          _____         _____           _    ", primaryColor)
	fmt.Println("    /\\   |  __ \\       |  __ \\         | |   ", primaryColor)
	fmt.Println("   /  \\  | |  | |______| |__) |_ _  ___| | __", primaryColor)
	fmt.Println("  / /\\ \\ | |  | |______|  ___/ _` |/ __| |/ /", primaryColor)
	fmt.Println(" / ____ \\| |__| |      | |  | (_| | (__|   < ", primaryColor)
	fmt.Println("/_/    \\_\\_____/       |_|   \\__,_|\\___|_|\\_", primaryColor)
	fmt.Println("                                              ")
}
