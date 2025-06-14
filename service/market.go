package service

import (
	"app/model"
	"app/renderer"
	"app/utils"
)

func Market(nftList model.TabNFT, nData int16, pageNumber int16, menuNumber int8) {
	var maxPage, entryPerPage int16

	entryPerPage = 15
	maxPage = utils.IntDivCeil(nData, entryPerPage)
	renderer.RenderMarketTable(nftList, nData, pageNumber, entryPerPage, maxPage)

	switch menuNumber {
	case 1:
		MarketMux1(&nftList, nData, pageNumber, maxPage)
	case 2:
		MarketMux2(&nftList, nData, pageNumber, maxPage)
	case 3:
		MarketMux3(&nftList, nData, pageNumber, maxPage)
	}
}
