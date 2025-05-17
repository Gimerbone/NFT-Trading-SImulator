package handler

import (
	"app/data"
	"app/renderer"
	"app/utils"
)

func handleMarket(nftList data.TabNFT, nData int16, pageNumber int16) {
	var (
		maxPage, entryPerPage int16
	)

	entryPerPage = 15
	maxPage = utils.IntDivCeil(nData, entryPerPage)

	renderer.RenderMarket(nftList, pageNumber, entryPerPage, maxPage)

	marketMenuMux(1)
}

func marketMenuMux(menuNumber int8) {
	var option int8

	switch menuNumber {
	case 1:
		renderer.RenderMarketMenu1(&option)
		marketOptionMux1()
	case 2:
		renderer.RenderMarketMenu2(&option)
	case 3:
		renderer.RenderMarketMenu3(&option)
	}
}

func marketOptionMux1(option int8, nftList data.TabNFT, nData int16, currentPage int16, maxPage int16) {
	switch option {
	case 1:
		if currentPage == maxPage {
			handleMarket(nftList, nData, maxPage)
		} else {
			handleMarket(nftList, nData, currentPage+1)
		}
	case 2:
		if currentPage == 1 {
			handleMarket(nftList, nData, 1)
		} else {
			handleMarket(nftList, nData, currentPage-1)
		}
	case 3:
		handleMarket(nftList, nData, 1)
	case 4:
		handleMarket(nftList, nData, maxPage)
	case 5:
		utils.SortIDAsc(&nftList, nData)
		handleMarket(nftList, nData, currentPage)
	}
}
