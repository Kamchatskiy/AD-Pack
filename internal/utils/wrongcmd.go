package utils

import (
	"fmt"

	"../colors"
)

func WrongCMD() {
	fmt.Println(colors.Colors.Warning, "Incorrect command")
	fmt.Println(colors.Colors.Secondary)
}
