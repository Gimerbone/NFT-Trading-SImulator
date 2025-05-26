package renderer

import (
	"app/data"
	"app/utils"
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func renderLogo() {
	fmt.Println("==============================================================")
	fmt.Println("             Sagitarius Market - NFT Marketplace")
	fmt.Println("  GitHub: https://github.com/Gimerbone/NFT-Trading-SImulator")
	fmt.Println("==============================================================")
}

func RenderOptionNotExist() {
	fmt.Println("Option does not exist.")
}

func RenderQuitMessage() {
	renderLogo()
	fmt.Println("Thanks for using our app!")
	fmt.Println()
}

func RenderUsernameInput(username *string) {
	renderLogo()
	fmt.Print("Enter your username (1-10 characters): ")
	utils.ScanSentence(username)
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
