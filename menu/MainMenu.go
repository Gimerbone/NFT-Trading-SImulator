package menu

import (
	"app/utils"
	"fmt"
)

func MainMenu() {
	fmt.Println("====================================================")
	fmt.Println("        Sagitarius Market - NFT Marketplace")
	fmt.Println("====================================================")
	fmt.Println(" 1. Browse Market")
	fmt.Println(" 2. View your portfolio")
	fmt.Println(" 3. Exit")

	var option int
	fmt.Print(" Choose Option: ")
	fmt.Scan(&option)
	for option < 1 || option > 3 {
		utils.ClearScreen()

		fmt.Println(" Option does not exist")
		fmt.Println()
		fmt.Println("====================================================")
		fmt.Println("        Sagitarius Market - NFT Marketplace")
		fmt.Println("====================================================")
		fmt.Println(" 1. Browse Market")
		fmt.Println(" 2. View your portfolio")
		fmt.Println(" 3. Exit")
		fmt.Print(" Choose Option (1-3): ")
		fmt.Scan(&option)
	}
}
