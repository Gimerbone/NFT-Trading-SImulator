package renderer

import (
	"app/data"
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func RenderLogo() {
	fmt.Println("==============================================================")
	fmt.Println("             Sagitarius Market - NFT Marketplace")
	fmt.Println("  GitHub: https://github.com/Gimerbone/NFT-Trading-SImulator")
	fmt.Println("==============================================================")
}

func RenderBalance() {
	fmt.Printf("%s%.2f%s\n",
		"Balance: ",
		data.User.BalanceETH, " ETH",
	)
}

func RenderOptionNotExist() {
	fmt.Println("Option does not exist.")
}

func RenderQuitMessage() {
	RenderLogo()
	fmt.Println("Thanks for using our app!")
	fmt.Println()
}

func RenderInvalidName() {
	fmt.Println("That name is not permitted.")
}

func RenderWelcome() {
	fmt.Printf("Welcome, %s!\n", data.User.Name)
}

func RenderNotFound(notFoundTarget string) {
	fmt.Printf("%s not found.\n--------------------------\n", notFoundTarget)
}

func RenderOperationFailed() {
	fmt.Println("Operation failed.")
}

func RenderNotEnoughBalance() {
	fmt.Println("Balance not enough.")
}
