package renderer

import (
	"app/data"
	"fmt"
)

func RenderMainMenu(option *int8) {
	fmt.Println("====================================================")
	fmt.Println("        Sagitarius Market - NFT Marketplace")
	fmt.Println("====================================================")
	fmt.Println(" 1. Browse Market")
	fmt.Println(" 2. List New NFT")
	fmt.Println(" 3. View Your portfolio")
	fmt.Println(" 4. Exit")

	fmt.Print(" Choose Option: ")
	fmt.Scan(option)
	for *option < 1 || *option > 4 {
		ClearScreen()

		fmt.Println(" Option does not exist")
		fmt.Println()
		fmt.Println("====================================================")
		fmt.Println("        Sagitarius Market - NFT Marketplace")
		fmt.Println("====================================================")
		fmt.Println(" 1. Browse Market")
		fmt.Println(" 2. List New NFT")
		fmt.Println(" 3. View your portfolio")
		fmt.Println(" 4. Exit")
		fmt.Print(" Choose Option (1-4): ")
		fmt.Scan(option)
	}
}

func RenderMarket() {
	fmt.Println("===========================================================================================================")
	fmt.Printf(" | %-4s | %-16s | %-12s | %-12s | %-10s | %-9s | %-10s | %-7s |\n",
		"ID", "Name", "Creator", "Owner", "Blockchain", "Price ETH", "Created At", "Royalty",
	)
	fmt.Println("===========================================================================================================")

	data.InitiateMarket()
	for i := 0; i < 5; i++ {
		fmt.Printf("%+v\n", data.NftList[i])
	}
}
