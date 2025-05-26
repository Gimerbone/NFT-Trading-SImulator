package renderer

import "fmt"

func RenderMainMenu(option *int8) {
	renderLogo()
	fmt.Println("1. Browse Market")
	fmt.Println("2. View Your portfolio")
	fmt.Println("0. Exit App")

	fmt.Print("Choose Option: ")
	fmt.Scanf("%d\n", option)
}
