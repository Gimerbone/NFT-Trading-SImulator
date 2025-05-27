package handler

import (
	"app/data"
	"app/renderer"
	"app/utils"
)

func HandleIDSearch(muxNumber int8) int8 {
	// Returns status code
	// 0: Search success
	// 1. User exit search
	var (
		result      data.TabNFT
		idx, target int16
	)

	renderer.ClearScreen()
	renderer.RenderIdSearch(&target)

	if target == -1 {
		return 1
	} else {
		idx = utils.SearchMarketID(originalList, nOriginalData, uint16(target))

		if idx != -1 {
			result[0] = originalList[idx]

			renderer.ClearScreen()
			handleMarket(result, 1, 1, muxNumber)
		} else {
			renderer.RenderNotFound("ID")
			return 1
		}

		return 0
	}
}

func handleNameSearch(muxNumber int8) int8 {
	// Returns status code
	// 0: Search success
	// 1. User exit search
	var (
		result data.TabNFT
		idx    int16
		target string
	)

	renderer.ClearScreen()
	renderer.RenderNameInput("nft name", &target, true)
	if target == "-1" {
		return 1
	} else {
		idx = utils.LSearchName(originalList, nOriginalData, target)
		if idx != -1 {
			result[0] = originalList[idx]

			renderer.ClearScreen()
			handleMarket(result, 1, 1, muxNumber)
		} else {
			renderer.RenderNotFound("Name")
			return 1
		}

		return 0
	}
}

func handlePurchase() int8 {
	var id int16
	renderer.ClearScreen()
	renderer.RenderIdSearch(&id)
	if id != -1 {
		return purchaseNFT(id)
	}
	return 0
}

func handleFilterBlockchain() int8 {
	var (
		blockchain    string
		filteredList  data.TabNFT
		nFilteredData int16
	)
	renderer.ClearScreen()
	renderer.RenderNameInput("blockchain name", &blockchain, true)
	if blockchain == "-1" {
		return 1
	}
	utils.FilterByBlockchain(blockchain, originalList, nOriginalData, &filteredList, &nFilteredData)
	handleMarket(filteredList, nFilteredData, 1, 3)
	return 0
}

func handleFilterCreator() int8 {
	var (
		creator       string
		filteredList  data.TabNFT
		nFilteredData int16
	)
	renderer.ClearScreen()
	renderer.RenderNameInput("creator name", &creator, true)
	if creator == "-1" {
		return 1
	}
	utils.FilterByCreator(creator, originalList, nOriginalData, &filteredList, &nFilteredData)
	handleMarket(filteredList, nFilteredData, 1, 3)
	return 0
}
