package renderer

import "fmt"

func RenderMainMenu(option *int8) {
	RenderLogo()
	fmt.Println("1. Browse Market")
	fmt.Println("2. View Your portfolio")
	fmt.Println("0. Exit App")

	fmt.Print("Choose Option: ")
	fmt.Scanf("%d\n", option)
}

func RenderMarketMenu1(option *int8) {
	fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"1. Refresh Page",
		"2. Next Page",
		"3. Previous Page",
		"4. First Page",
		"5. Last Page",
		"6. Sort by ID (Ascending)",
		"7. Sort by ID (Descending)",
		"8. Sort by Price (Ascending)",
		"9. Next Option",
		"0. Back",
	)

	fmt.Print("Choose Option: ")
	fmt.Scanf("%d\n", option)
}

func RenderMarketMenu2(option *int8) {
	fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"1. Sort by Price (Descending)",
		"2. Sort by Date (Ascending)",
		"3. Sort by Date (Descending)",
		"4. Sort by Royalty (Ascending)",
		"5. Sort by Royalty (Descending)",
		"6. Search by ID",
		"7. Search by NFT Name",
		"8. Next Option",
		"9. Previous Option",
		"0. Back",
	)

	fmt.Print("Choose Option: ")
	fmt.Scanf("%d\n", option)
}

func RenderMarketMenu3(option *int8) {
	fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"1. Purchase NFT",
		"2. Filter by Blockchain Type",
		"3. Filter by Creator Name",
		"4. Filter by Release Year",
		"5. Filter by Owner Name",
		"6. Show only your NFT",
		"9. Previous Option",
		"0. Back",
	)

	fmt.Print("Choose Option: ")
	fmt.Scanf("%d\n", option)
}
