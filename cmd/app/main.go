package main

import (
	"app/data"
	"app/renderer"
	"app/utils"
	"fmt"
)

func main() {

}

func mainHandler() {
	var (
		option     int8
		marketList data.TabNFT
	)

	data.InitiateMarket(&marketList)
	renderer.RenderMainMenu(&option)
	mainMux(option, marketList)
}

func mainMux(option int8, marketList data.TabNFT) {
	switch option {
	case 1:
		marketHandler(marketList, 1)
	case 2:
		portfolioHandler()
	}
}

func marketHandler(nftList data.TabNFT, pageNumber int16) {
	var (
		option                int8
		maxPage, itemsPerPage int16
	)

	maxPage = utils.IntDivCeil(int16(data.NMAX), 15)
	itemsPerPage = 15

	renderer.RenderMarket(nftList, pageNumber, itemsPerPage, maxPage)

	for {
		renderer.RenderMarketMenu(&option)

		if option > -1 && option < 16 {
			break
		}

		renderer.ClearScreen()
		fmt.Printf("Option Does Not Exist\n----------------------\n\n")
	}

	marketMux(option, nftList, pageNumber, maxPage)
}

func marketMux(option int8, nftList data.TabNFT, currentPage int16, maxPage int16) {
	switch option {
	case 1:
		if currentPage == maxPage {
			marketHandler(nftList, maxPage)
		} else {
			marketHandler(nftList, currentPage+1)
		}
	case 2:
		if currentPage == 1 {
			marketHandler(nftList, 1)
		} else {
			marketHandler(nftList, currentPage-1)
		}
	case 3:
		marketHandler(nftList, 1)
	case 4:
		marketHandler(nftList, maxPage)
	case 5:
		sortByID
	}
}

func portfolioHandler() {}
