package service

import (
	"app/model"
	"app/renderer"
	"app/utils"
)

func SellOrder(nftList model.TabNFT, nData int16, pageNumber int16) {
	var maxPage, entryPerPage int16

	entryPerPage = 15
	maxPage = utils.IntDivCeil(nData, entryPerPage)
	renderer.RenderSOTable(nftList, nData, pageNumber, entryPerPage, maxPage)

	sOMux(&nftList, nData, pageNumber, maxPage)
}
