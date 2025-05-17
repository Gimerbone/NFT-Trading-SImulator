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
	fmt.Println(" 2. View Your portfolio")
	fmt.Println(" 0. Exit App")

	fmt.Print(" Choose Option: ")
	fmt.Scan(option)
	for *option < 0 || *option > 2 {
		ClearScreen()

		fmt.Println(" Option does not exist")
		fmt.Println()
		fmt.Println("====================================================")
		fmt.Println("        Sagitarius Market - NFT Marketplace")
		fmt.Println("====================================================")
		fmt.Println(" 1. Browse Market")
		fmt.Println(" 2. View your portfolio")
		fmt.Println(" 0. Exit App")
		fmt.Print(" Choose Option (1-4): ")
		fmt.Scan(option)
	}
}

func RenderMarket(data data.TabNFT, pageNumber, entryPerPage, maxPage int16) {
	fmt.Println("===================================================================================================================")
	fmt.Printf(" |  %s  |           %s           |   %s    |    %s     | %-10s | %-9s | %-10s | %-7s |\n",
		"ID", "Name", "Creator", "Owner", "Blockchain", "Price ETH", "Created At", "Royalty",
	)
	fmt.Println(" ----------------------------------------------------------------------------------------------------------------- ")

	renderMarketList(data, (pageNumber-1)*entryPerPage, pageNumber*entryPerPage)

	fmt.Println("===================================================================================================================")
	fmt.Printf("  PAGES %d/%d\n", pageNumber, maxPage)
	fmt.Println("===================================================================================================================")
}

func RenderMarketMenu(option *int8) {
	fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"01. Next Page",
		"02. Previous Page",
		"03. First Page",
		"04. Last Page",
		"05. Sort by ID",
		"06. Sort by Price",
		"07. Sort by Date",
		"08. Sort by Royalty",
		"09. Purchase NFT",
		"10. Search by ID",
		"11. Search by NFT Name",
		"12. Filter by Blockchain Type",
		"13. Filter by Creator Name",
		"14. Filter by Release Year",
		"15. Filter by Owner Name",
		"0. Back",
	)

	fmt.Println("Choose Option:")
	fmt.Scanln(option)
}

func renderMarketList(arr data.TabNFT, entryStart int16, entryEnd int16) {
	for i := entryStart; i < entryEnd; i++ {
		fmt.Printf(" | %4d | %-24s | %-12s | %-12s | %-10s | %9.2f | %10s | %5.0f %% |\n",
			arr[i].ID, arr[i].Name, arr[i].Creator, arr[i].Owner,
			arr[i].Blockchain, arr[i].PriceETH, arr[i].CreatedAt,
			arr[i].Royalty*100,
		)
	}
}
