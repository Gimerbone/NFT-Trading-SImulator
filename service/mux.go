package service

import (
	"app/model"
	"app/renderer"
	"app/storage"
	"app/utils"
	"fmt"
)

// This is meant to be used for filter by owner name
var isOwnerPlayer bool = true

func MainMenu() {
	var option int8

	fmt.Printf("Welcome, %s!\n", model.User.Name)
	renderer.RenderLogo()
	fmt.Println("1. Browse Market")
	fmt.Println("2. View Your portfolio")
	fmt.Println("0. Exit App")
	fmt.Print("Choose Option: ")
	fmt.Scanf("%d\n", &option)

	switch option {
	case 0:
		renderer.ClearScreen()
		renderer.QuitScreen()
	case 1:
		renderer.ClearScreen()
		Market(storage.MarketList, storage.NMarketData, 1, 1)
	case 2:
		renderer.ClearScreen()
		Portfolio(storage.PortList, storage.NPortData, 1, 1)
	default:
		renderer.ClearScreen()
		fmt.Println("Option does not exist.")
		MainMenu()
	}
}

func MarketMux1(nftList *model.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

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
	fmt.Scanf("%d\n", &option)

	switch option {
	case 0:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.MarketList, storage.NMarketData)
		model.MarketSortCode = 1
		MainMenu()
	case 1:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.MarketList, storage.NMarketData)
		model.MarketSortCode = 1
		Market(storage.MarketList, storage.NMarketData, 1, 1)
	case 2:
		if currentPage == maxPage {
			renderer.ClearScreen()
			Market(*nftList, nData, maxPage, 1)
		} else {
			renderer.ClearScreen()
			Market(*nftList, nData, currentPage+1, 1)
		}
	case 3:
		if currentPage == 1 {
			renderer.ClearScreen()
			Market(*nftList, nData, 1, 1)
		} else {
			renderer.ClearScreen()
			Market(*nftList, nData, currentPage-1, 1)
		}
	case 4:
		renderer.ClearScreen()
		Market(*nftList, nData, 1, 1)
	case 5:
		renderer.ClearScreen()
		Market(*nftList, nData, maxPage, 1)
	case 6:
		utils.SortIDAsc(nftList, nData)
		model.MarketSortCode = 1
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 1)
	case 7:
		utils.SortIDDsc(nftList, nData)
		model.MarketSortCode = 2
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 1)
	case 8:
		utils.SortPriceAsc(nftList, nData)
		model.MarketSortCode = 3
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 1)
	case 9:
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 2)
	default:
		renderer.ClearScreen()
		fmt.Println("Option does not exist.")
		Market(*nftList, nData, currentPage, 1)
	}
}

func MarketMux2(nftList *model.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

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
	fmt.Scanf("%d\n", &option)

	switch option {
	case 0:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.MarketList, storage.NMarketData)
		model.MarketSortCode = 1
		MainMenu()
	case 1:
		utils.SortPriceDsc(nftList, nData)
		model.MarketSortCode = 4
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 2)
	case 2:
		utils.SortDateAsc(nftList, nData)
		model.MarketSortCode = 5
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 2)
	case 3:
		utils.SortDateDsc(nftList, nData)
		model.MarketSortCode = 6
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 2)
	case 4:
		utils.SortRoyaltyAsc(nftList, nData)
		model.MarketSortCode = 7
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 2)
	case 5:
		utils.SortRoyaltyDsc(nftList, nData)
		model.MarketSortCode = 8
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 2)
	case 6:
		var (
			result      model.TabNFT
			idx, target int16
		)

		renderer.ClearScreen()
		utils.InputID(&target)

		if target == -1 {
			renderer.ClearScreen()
			Market(*nftList, nData, currentPage, 2)
		} else {
			idx = utils.SearchID(storage.MarketList, storage.NMarketData, uint16(target))
			if idx != -1 {
				result[0] = storage.MarketList[idx]
				renderer.ClearScreen()
				Market(result, 1, 1, 2)
			} else {
				renderer.ClearScreen()
				fmt.Println("ID not found.")
				fmt.Println("----------------")
				Market(*nftList, nData, currentPage, 2)
			}
		}
	case 7:
		var (
			result model.TabNFT
			idx    int16
			target string
		)

		renderer.ClearScreen()
		utils.InputName("nft name", true, &target)
		if target == "-1" {
			renderer.ClearScreen()
			Market(*nftList, nData, currentPage, 2)
		} else {
			idx = utils.LSearchName(storage.MarketList, storage.NMarketData, target)
			if idx != -1 {
				result[0] = storage.MarketList[idx]

				renderer.ClearScreen()
				Market(result, 1, 1, 2)
			} else {
				renderer.ClearScreen()
				fmt.Println("Name not found.")
				fmt.Println("------------------")
				Market(*nftList, nData, currentPage, 2)
			}
		}
	case 8:
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 3)
	case 9:
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 1)
	default:
		renderer.ClearScreen()
		fmt.Println("Option does not exist.")
		Market(*nftList, nData, currentPage, 2)
	}
}

func MarketMux3(nftList *model.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

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
	fmt.Scanf("%d\n", &option)

	switch option {
	case 0:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.MarketList, storage.NMarketData)
		model.MarketSortCode = 1
		MainMenu()
	case 1:
		var id int16
		renderer.ClearScreen()
		utils.InputID(&id)

		renderer.ClearScreen()
		if id != -1 && purchaseNFT(uint16(id)) == 1 {
			fmt.Println("Balance not enough.")
		}
		Market(storage.MarketList, storage.NMarketData, currentPage, 3)
	case 2:
		var (
			blockchain    string
			filteredList  model.TabNFT
			nFilteredData int16
		)
		renderer.ClearScreen()
		utils.InputName("blockchain name", true, &blockchain)
		if blockchain == "-1" {
			renderer.ClearScreen()
			Market(*nftList, nData, currentPage, 3)
		}
		utils.FilterByBlockchain(blockchain, storage.MarketList, storage.NMarketData, &filteredList, &nFilteredData)
		Market(filteredList, nFilteredData, 1, 3)
	case 3:
		var (
			creator       string
			filteredList  model.TabNFT
			nFilteredData int16
		)
		renderer.ClearScreen()
		utils.InputName("creator name", true, &creator)
		if creator == "-1" {
			renderer.ClearScreen()
			Market(*nftList, nData, currentPage, 3)
		}
		utils.FilterByCreator(creator, storage.MarketList, storage.NMarketData, &filteredList, &nFilteredData)
		Market(filteredList, nFilteredData, 1, 3)
	case 4:
		var (
			year          string
			filteredList  model.TabNFT
			nFilteredData int16
		)
		renderer.ClearScreen()
		utils.InputName("year name", true, &year)
		if year == "-1" {
			renderer.ClearScreen()
			Market(*nftList, nData, currentPage, 3)
		}
		utils.FilterByYear(year, storage.MarketList, storage.NMarketData, &filteredList, &nFilteredData)
		Market(filteredList, nFilteredData, 1, 3)
	case 5:
		isOwnerPlayer = false
		fallthrough
	case 6:
		var (
			owner         string
			filteredList  model.TabNFT
			nFilteredData int16
		)
		renderer.ClearScreen()
		if !isOwnerPlayer {
			utils.InputName("owner name", true, &owner)
			if owner == "-1" {
				renderer.ClearScreen()
				Market(*nftList, nData, currentPage, 3)
			}
		} else {
			owner = model.User.Name
		}
		utils.FilterByOwner(owner, storage.MarketList, storage.NMarketData, &filteredList, &nFilteredData)
		isOwnerPlayer = true
		Market(filteredList, nFilteredData, 1, 3)
	case 9:
		renderer.ClearScreen()
		Market(*nftList, nData, currentPage, 2)
	default:
		renderer.ClearScreen()
		fmt.Println("Option does not exist.")
		Market(*nftList, nData, currentPage, 3)
	}
}

func PortMux1(nftList *model.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

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
	fmt.Scanf("%d\n", &option)

	switch option {
	case 0:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.PortList, storage.NPortData)
		model.PortSortCode = 9
		MainMenu()
	case 1:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.PortList, storage.NPortData)
		model.PortSortCode = 9
		Portfolio(storage.PortList, storage.NPortData, 1, 1)
	case 2:
		if currentPage == maxPage {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, maxPage, 1)
		} else {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, currentPage+1, 1)
		}
	case 3:
		if currentPage == 1 {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, 1, 1)
		} else {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, currentPage-1, 1)
		}
	case 4:
		renderer.ClearScreen()
		Portfolio(*nftList, nData, 1, 1)
	case 5:
		renderer.ClearScreen()
		Portfolio(*nftList, nData, maxPage, 1)
	case 6:
		utils.SortIDAsc(nftList, nData)
		model.PortSortCode = 1
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 1)
	case 7:
		utils.SortIDDsc(nftList, nData)
		model.PortSortCode = 2
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 1)
	case 8:
		utils.SortPriceAsc(nftList, nData)
		model.PortSortCode = 3
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 1)
	case 9:
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 2)
	default:
		renderer.ClearScreen()
		fmt.Println("Option does not exist.")
		Portfolio(*nftList, nData, currentPage, 1)
	}
}

func PortMux2(nftList *model.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

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
	fmt.Scanf("%d\n", &option)

	switch option {
	case 0:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.PortList, storage.NPortData)
		model.PortSortCode = 9
		MainMenu()
	case 1:
		utils.SortPriceDsc(nftList, nData)
		model.PortSortCode = 4
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 2)
	case 2:
		utils.SortDateAsc(nftList, nData)
		model.PortSortCode = 5
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 2)
	case 3:
		utils.SortDateDsc(nftList, nData)
		model.PortSortCode = 6
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 2)
	case 4:
		utils.SortRoyaltyAsc(nftList, nData)
		model.PortSortCode = 7
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 2)
	case 5:
		utils.SortRoyaltyDsc(nftList, nData)
		model.PortSortCode = 8
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 2)
	case 6:
		var (
			result      model.TabNFT
			idx, target int16
		)

		renderer.ClearScreen()
		utils.InputID(&target)

		if target == -1 {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, currentPage, 2)
		} else {
			idx = utils.SearchID(storage.PortList, storage.NPortData, uint16(target))
			if idx != -1 {
				result[0] = storage.PortList[idx]
				renderer.ClearScreen()
				Portfolio(result, 1, 1, 2)
			} else {
				renderer.ClearScreen()
				fmt.Println("ID not found.")
				fmt.Println("----------------")
				Portfolio(*nftList, nData, currentPage, 2)
			}
		}
	case 7:
		var (
			result model.TabNFT
			idx    int16
			target string
		)

		renderer.ClearScreen()
		utils.InputName("nft name", true, &target)
		if target == "-1" {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, currentPage, 2)
		} else {
			idx = utils.LSearchName(storage.PortList, storage.NPortData, target)
			if idx != -1 {
				result[0] = storage.PortList[idx]

				renderer.ClearScreen()
				Portfolio(result, 1, 1, 2)
			} else {
				renderer.ClearScreen()
				fmt.Println("Name not found.")
				fmt.Println("------------------")
				Portfolio(*nftList, nData, currentPage, 2)
			}
		}
	case 8:
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 3)
	case 9:
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 1)
	default:
		renderer.ClearScreen()
		fmt.Println("Option does not exist.")
		Portfolio(*nftList, nData, currentPage, 2)
	}
}

func PortMux3(nftList *model.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

	fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"1. Sort by bought time",
		"1. Filter by Blockchain Type",
		"2. Filter by Creator Name",
		"3. Filter by Release Year",
		"9. Previous Option",
		"0. Back",
	)

	fmt.Print("Choose Option: ")
	fmt.Scanf("%d\n", &option)

	switch option {
	case 0:
		renderer.ClearScreen()
		utils.RandomizePrice(&storage.PortList, storage.NPortData)
		model.PortSortCode = 9
		MainMenu()
	case 1:
		model.PortSortCode = 9
		Portfolio(storage.PortList, storage.NPortData, 1, 3)
	case 2:
		var (
			blockchain    string
			filteredList  model.TabNFT
			nFilteredData int16
		)
		renderer.ClearScreen()
		utils.InputName("blockchain name", true, &blockchain)
		if blockchain == "-1" {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, currentPage, 3)
		}
		utils.FilterByBlockchain(blockchain, storage.PortList, storage.NPortData, &filteredList, &nFilteredData)
		Portfolio(filteredList, nFilteredData, 1, 3)
	case 3:
		var (
			creator       string
			filteredList  model.TabNFT
			nFilteredData int16
		)
		renderer.ClearScreen()
		utils.InputName("creator name", true, &creator)
		if creator == "-1" {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, currentPage, 3)
		}
		utils.FilterByCreator(creator, storage.PortList, storage.NPortData, &filteredList, &nFilteredData)
		Portfolio(filteredList, nFilteredData, 1, 3)
	case 4:
		var (
			year          string
			filteredList  model.TabNFT
			nFilteredData int16
		)
		renderer.ClearScreen()
		utils.InputName("year name", true, &year)
		if year == "-1" {
			renderer.ClearScreen()
			Portfolio(*nftList, nData, currentPage, 3)
		}
		utils.FilterByYear(year, storage.PortList, storage.NPortData, &filteredList, &nFilteredData)
		Portfolio(filteredList, nFilteredData, 1, 3)
	case 9:
		renderer.ClearScreen()
		Portfolio(*nftList, nData, currentPage, 2)
	default:
		renderer.ClearScreen()
		fmt.Println("Option does not exist.")
		Portfolio(*nftList, nData, currentPage, 3)
	}
}
