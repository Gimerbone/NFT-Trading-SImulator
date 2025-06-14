package renderer

import (
	"app/model"
	"app/storage"
	"app/utils"
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

func WelcomeScreen() {
	var (
		username string
	)

	ClearScreen()
	RenderLogo()
	for {
		utils.InputName("username (1-10 characters)", false, &username)
		if len(username) < 11 && username != "" {
			model.User.Name = username
			break
		}
		ClearScreen()
		fmt.Println("Length must be 1-10 characters")
	}
}

func QuitScreen() {
	RenderLogo()
	fmt.Println("Thanks for using our app!")
	fmt.Println()
}

// RenderMarketTable prints a table of NFTs with pagination.
// Table rows are generated from nftList. It renders only the entries
// for the specified pageNumber and shows the user's ETH balance.
// Entry per page is 15 and can be changed in this function.
func RenderMarketTable(nftList model.TabNFT, nData, pageNumber, entryPerPage, maxPage int16) {
	var (
		entryStart, entryEnd int16
	)

	fmt.Printf("%s%.2f%s\n", "Balance: ", model.User.BalanceETH, " ETH")

	fmt.Println("===================================================================================================================")
	fmt.Printf(" |  %s  |           %s           |   %s    |    %s     | %-10s | %-9s | %-10s | %-7s |\n",
		"ID", "Name", "Creator", "Owner", "Blockchain", "Price ETH", "Created At", "Royalty",
	)
	fmt.Println(" ----------------------------------------------------------------------------------------------------------------- ")

	entryStart = (pageNumber - 1) * entryPerPage
	entryEnd = pageNumber * entryPerPage
	if entryEnd > nData {
		entryEnd = nData
	}

	for i := entryStart; i < entryEnd; i++ {
		fmt.Printf(" | %4d | %-24s | %-12s | %-12s | %-10s | %9.2f | %10s | %5.0f %% |\n",
			nftList[i].ID, nftList[i].Name, nftList[i].Creator, nftList[i].Owner,
			nftList[i].Blockchain, nftList[i].PriceETH, nftList[i].CreatedAt,
			nftList[i].Royalty*100,
		)
	}

	fmt.Println("===================================================================================================================")
	fmt.Printf("  PAGES %d/%d\n", pageNumber, maxPage)
	fmt.Println("===================================================================================================================")
}

// RenderPortTable prints a table of NFTs with pagination.
// Table rows are generated from nftList. It renders only the entries
// for the specified pageNumber and shows the user's ETH balance.
// Entry per page is 15 and can be changed in this function.
func RenderPortTable(nftList model.TabNFT, nData, pageNumber, entryPerPage, maxPage int16) {
	var (
		entryStart, entryEnd int16
	)

	fmt.Printf("%s%.2f%s\n", "Balance: ", model.User.BalanceETH, " ETH")
	fmt.Printf("%s%.2f%s\n", "Total Value: ", storage.TotalPortValue, "ETH")

	fmt.Println("===================================================================================================================")
	fmt.Printf(" |  %s  |           %s           |   %s    | %-10s | %-12s | %-10s | %-7s |\n",
		"ID", "Name", "Creator", "Blockchain", "Market Price", "Created At", "Royalty",
	)
	fmt.Println(" ----------------------------------------------------------------------------------------------------------------- ")

	entryStart = (pageNumber - 1) * entryPerPage
	entryEnd = pageNumber * entryPerPage
	if entryEnd > nData {
		entryEnd = nData
	}

	for i := entryStart; i < entryEnd; i++ {
		fmt.Printf(" | %4d | %-24s | %-12s | %-10s | %12.2f | %10s | %5.0f %% |\n",
			nftList[i].ID, nftList[i].Name, nftList[i].Creator,
			nftList[i].Blockchain, nftList[i].PriceETH, nftList[i].CreatedAt,
			nftList[i].Royalty*100,
		)
	}

	fmt.Println("===================================================================================================================")
	fmt.Printf("  PAGES %d/%d\n", pageNumber, maxPage)
	fmt.Println("===================================================================================================================")
}

// RenderSOTable prints a table of NFTs with pagination.
// Table rows are generated from nftList. It renders only the entries
// for the specified pageNumber and shows the user's ETH balance.
// Entry per page is 15 and can be changed in this function.
func RenderSOTable(nftList model.TabNFT, nData, pageNumber, entryPerPage, maxPage int16) {
	var (
		entryStart, entryEnd int16
	)

	fmt.Printf("%s%.2f%s\n", "Balance: ", model.User.BalanceETH, " ETH")

	fmt.Println("===================================================================================================================")
	fmt.Printf(" |  %s  |           %s           |   %s   |   %s   |\n",
		"ID", "Name", "Bid Price", "Market Price",
	)
	fmt.Println(" ----------------------------------------------------------------------------------------------------------------- ")

	entryStart = (pageNumber - 1) * entryPerPage
	entryEnd = pageNumber * entryPerPage
	if entryEnd > nData {
		entryEnd = nData
	}

	for i := entryStart; i < entryEnd; i++ {
		fmt.Printf(" | %4d | %-24s | %-12.2f | %-12.2f |\n",
			nftList[i].ID, nftList[i].Name, 100.0, nftList[i].PriceETH,
		)
	}

	fmt.Println("===================================================================================================================")
	fmt.Printf("  PAGES %d/%d\n", pageNumber, maxPage)
	fmt.Println("===================================================================================================================")
}
