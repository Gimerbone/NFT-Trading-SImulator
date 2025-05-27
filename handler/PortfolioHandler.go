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
