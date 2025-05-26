package renderer

import (
	"app/data"
	"app/utils"
	"fmt"
)

func RenderBalance() {
	fmt.Printf("%s%.2f%s\n",
		"Balance: ",
		data.User.BalanceETH, " ETH",
	)
}

func RenderMarket(data data.TabNFT, pageNumber, entryPerPage, maxPage, nData int16) {
	fmt.Println("===================================================================================================================")
	fmt.Printf(" |  %s  |           %s           |   %s    |    %s     | %-10s | %-9s | %-10s | %-7s |\n",
		"ID", "Name", "Creator", "Owner", "Blockchain", "Price ETH", "Created At", "Royalty",
	)
	fmt.Println(" ----------------------------------------------------------------------------------------------------------------- ")

	renderMarketList(data, (pageNumber-1)*entryPerPage, pageNumber*entryPerPage, nData)

	fmt.Println("===================================================================================================================")
	fmt.Printf("  PAGES %d/%d\n", pageNumber, maxPage)
	fmt.Println("===================================================================================================================")
}

func renderMarketList(arr data.TabNFT, entryStart, entryEnd, nData int16) {
	if entryEnd > nData {
		entryEnd = nData
	}

	for i := entryStart; i < entryEnd; i++ {
		fmt.Printf(" | %4d | %-24s | %-12s | %-12s | %-10s | %9.2f | %10s | %5.0f %% |\n",
			arr[i].ID, arr[i].Name, arr[i].Creator, arr[i].Owner,
			arr[i].Blockchain, arr[i].PriceETH, arr[i].CreatedAt,
			arr[i].Royalty*100,
		)
	}
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

func RenderIdSearch(resultId *int16) {
	for {
		fmt.Printf("Enter nft id or -1 to exit:")
		fmt.Scanln(resultId)

		if *resultId > -2 && *resultId < 1000 && *resultId != 0 {
			break
		}

		ClearScreen()
		fmt.Printf("%s\n%s\n",
			"ID doesn't exist. Choose between 1-1000 or -1 to exit.",
			"-------------------------------------------------------",
		)
	}
}

func RenderNameSearch(resultName *string) {
	fmt.Printf("Enter nft name (you must end it with '.') or -1 to exit:")
	utils.ScanSentence(resultName)
}

func RenderBlockchainInput(blockchain *string) {
	fmt.Println("Enter blockchain name or -1 to exit:")
	fmt.Scan(&blockchain)
}
