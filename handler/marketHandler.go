package handler

import (
	"app/data"
	"app/renderer"
	"app/utils"
)

var originalList data.TabNFT
var nOriginalData int16

func InitiateMarket() {
	data.InitiateMarketList(&originalList, &nOriginalData)
}

func handleMarket(nftList data.TabNFT, nData int16, pageNumber int16, muxNumber int8) {
	renderer.ClearScreen()

	var (
		maxPage, entryPerPage int16
	)

	entryPerPage = 15
	maxPage = utils.IntDivCeil(nData, entryPerPage)

	renderer.RenderMarket(nftList, pageNumber, entryPerPage, maxPage)

	marketMenuMux(muxNumber, nftList, nData, pageNumber, maxPage)
}

func marketMenuMux(menuNumber int8, nftList data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

	switch menuNumber {
	case 1:
		renderer.RenderMarketMenu1(&option)
		marketOptionMux1(option, &nftList, nData, currentPage, maxPage)
	case 2:
		renderer.RenderMarketMenu2(&option)
	case 3:
		renderer.RenderMarketMenu3(&option)
	}
}

func marketOptionMux1(option int8, nftList *data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	switch option {
	case 2:
		if currentPage == maxPage {
			handleMarket(*nftList, nData, maxPage, 1)
		} else {
			handleMarket(*nftList, nData, currentPage+1, 1)
		}
	case 3:
		if currentPage == 1 {
			handleMarket(*nftList, nData, 1, 1)
		} else {
			handleMarket(*nftList, nData, currentPage-1, 1)
		}
	case 4:
		handleMarket(*nftList, nData, 1, 1)
	case 5:
		handleMarket(*nftList, nData, maxPage, 1)
	case 6:
		utils.SortIDAsc(nftList, nData)
		data.StatusCode = 1
		handleMarket(*nftList, nData, currentPage, 1)
	case 7:
		utils.SortIDDsc(nftList, nData)
		data.StatusCode = 2
		handleMarket(*nftList, nData, currentPage, 1)
	case 8:
		utils.SortPriceAsc(nftList, nData)
		data.StatusCode = 3
		handleMarket(*nftList, nData, currentPage, 1)
	case 9:
		handleMarket(*nftList, nData, currentPage, 2)
	}
}

func marketOptionMux2(option int8, nftList *data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	switch option {
	case 1:
		utils.SortPriceDsc(nftList, nData)
		data.StatusCode = 4
		handleMarket(*nftList, nData, currentPage, 2)
	case 2:
		utils.SortDateAsc(nftList, nData)
		data.StatusCode = 5
		handleMarket(*nftList, nData, currentPage, 2)
	case 3:
		utils.SortDateDsc(nftList, nData)
		data.StatusCode = 6
		handleMarket(*nftList, nData, currentPage, 2)
	case 4:
		utils.SortRoyaltyAsc(nftList, nData)
		data.StatusCode = 7
		handleMarket(*nftList, nData, currentPage, 2)
	case 5:
		utils.SortRoyaltyDsc(nftList, nData)
		data.StatusCode = 8
		handleMarket(*nftList, nData, currentPage, 2)
	case 6:
		code := handleIDSearch(2)
		if code == 1 {
			handleMarket(*nftList, nData, currentPage, 2)
		}
	case 7:
		code := handleNameSearch(2)
		if code == 1 {
			handleMarket(*nftList, nData, currentPage, 2)
		}
	case 8:
		handleMarket(*nftList, nData, currentPage, 3)
	case 9:
		handleMarket(*nftList, nData, currentPage, 1)
	}
}

func handleIDSearch(muxNumber int8) int8 {
	// Returns status code
	// 0: Search success
	// 1. User exit search
	var (
		result               data.TabNFT
		nResult, idx, target int16
	)

	renderer.RenderIdSearch(&target)

	if target == -1 {
		return 1
	} else {
		switch data.StatusCode {
		case 1:
			idx = utils.BSearchAscID(originalList, nOriginalData, uint16(target))
		case 2:
			idx = utils.BSearchDscID(originalList, nOriginalData, uint16(target))
		default:
			idx = utils.LSearchID(originalList, nOriginalData, uint16(target))
		}

		if idx != -1 {
			result[0] = originalList[idx]
			nResult = 1

			handleMarket(result, nResult, 1, muxNumber)
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
		result       data.TabNFT
		nResult, idx int16
		target       string
	)

	renderer.RenderNameSearch(&target)
	if target == "-1" {
		return 1
	} else {
		idx = utils.LSearchName(originalList, nOriginalData, target)
		if idx != -1 {
			result[0] = originalList[idx]
			nResult = 1

			handleMarket(result, nResult, 1, muxNumber)
		} else {
			renderer.RenderNotFound("Name")
			return 1
		}

		return 0
	}
}
