package handler

import (
	"app/data"
	"app/renderer"
	"app/utils"
)

func handlePortfolio(nftList data.TabNFT, nData int16, pageNumber int16, muxNumber int8) {
	var (
		maxPage, entryPerPage int16
	)

	entryPerPage = 15
	maxPage = utils.IntDivCeil(nData, entryPerPage)

	renderer.RenderBalance()
	renderer.RenderMarket(nftList, pageNumber, entryPerPage, maxPage, nData)

	marketMenuMux(muxNumber, nftList, nData, pageNumber, maxPage)
}

func PortMenuMux(menuNumber int8, nftList data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	var option int8

	switch menuNumber {
	case 1:
		renderer.RenderMarketMenu1(&option)
		portOptionMux1(option, &nftList, nData, currentPage, maxPage)
	case 2:
		renderer.RenderMarketMenu2(&option)
		portOptionMux2(option, &nftList, nData, currentPage, maxPage)
	case 3:
		renderer.RenderPortfolioMenu3(&option)
		portOptionMux3(option, &nftList, nData, currentPage, maxPage)
	}
}

func portOptionMux1(option int8, nftList *data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	switch option {
	case 0:
		renderer.ClearScreen()
		randomizePrice()
		data.StatusCode = 1
		handleMain()
	case 1:
		data.StatusCode = 1
		renderer.ClearScreen()
		randomizePrice()
		handleMarket(originalList, nOriginalData, 1, 1)
	case 2:
		if currentPage == maxPage {
			renderer.ClearScreen()
			handleMarket(*nftList, nData, maxPage, 1)
		} else {
			renderer.ClearScreen()
			handleMarket(*nftList, nData, currentPage+1, 1)
		}
	case 3:
		if currentPage == 1 {
			renderer.ClearScreen()
			handleMarket(*nftList, nData, 1, 1)
		} else {
			renderer.ClearScreen()
			handleMarket(*nftList, nData, currentPage-1, 1)
		}
	case 4:
		renderer.ClearScreen()
		handleMarket(*nftList, nData, 1, 1)
	case 5:
		renderer.ClearScreen()
		handleMarket(*nftList, nData, maxPage, 1)
	case 6:
		utils.SortIDAsc(nftList, nData)
		data.StatusCode = 1
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 1)
	case 7:
		utils.SortIDDsc(nftList, nData)
		data.StatusCode = 2
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 1)
	case 8:
		utils.SortPriceAsc(nftList, nData)
		data.StatusCode = 3
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 1)
	case 9:
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 2)
	default:
		renderer.ClearScreen()
		renderer.RenderOptionNotExist()
		handleMarket(*nftList, nData, currentPage, 1)
	}
}

func portOptionMux2(option int8, nftList *data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	switch option {
	case 0:
		renderer.ClearScreen()
		randomizePrice()
		data.StatusCode = 1
		handleMain()
	case 1:
		utils.SortPriceDsc(nftList, nData)
		data.StatusCode = 4
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 2)
	case 2:
		utils.SortDateAsc(nftList, nData)
		data.StatusCode = 5
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 2)
	case 3:
		utils.SortDateDsc(nftList, nData)
		data.StatusCode = 6
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 2)
	case 4:
		utils.SortRoyaltyAsc(nftList, nData)
		data.StatusCode = 7
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 2)
	case 5:
		utils.SortRoyaltyDsc(nftList, nData)
		data.StatusCode = 8
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 2)
	case 6:
		code := HandleMIDSearch(2)
		if code == 1 {
			renderer.ClearScreen()
			handleMarket(*nftList, nData, currentPage, 2)
		}
	case 7:
		code := handleMNameSearch(2)
		if code == 1 {
			renderer.ClearScreen()
			handleMarket(*nftList, nData, currentPage, 2)
		}
	case 8:
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 3)
	case 9:
		renderer.ClearScreen()
		handleMarket(*nftList, nData, currentPage, 1)
	default:
		renderer.ClearScreen()
		renderer.RenderOptionNotExist()
		handleMarket(*nftList, nData, currentPage, 2)
	}
}

func portOptionMux3(option int8, nftList *data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	switch option {
	case 0:
		renderer.ClearScreen()
		randomizePrice()
		data.StatusCode = 1
		handleMain()
	case 2:
		var code int8
		code = handlePFilterBlockchain()
		if code == 1 {
			renderer.ClearScreen()
			handlePortfolio(*nftList, nData, currentPage, 3)
		}
	case 3:
		var code int8
		code = handlePFilterCreator()
		if code == 1 {
			renderer.ClearScreen()
			handlePortfolio(*nftList, nData, currentPage, 3)
		}
	case 4:
		var code int8
		code = handlePFilterYear()
		if code == 1 {
			renderer.ClearScreen()
			handlePortfolio(*nftList, nData, currentPage, 3)
		}
	case 9:
		renderer.ClearScreen()
		handlePortfolio(*nftList, nData, currentPage, 2)
	default:
		renderer.ClearScreen()
		renderer.RenderOptionNotExist()
		handlePortfolio(*nftList, nData, currentPage, 3)
	}
}
