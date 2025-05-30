package service

import (
	"app/model"
	"app/renderer"
	"app/utils"
)

func Portfolio(nftList model.TabNFT, nData int16, pageNumber int16, menuNumber int8) {
	var maxPage, entryPerPage int16

	entryPerPage = 15
	maxPage = utils.IntDivCeil(nData, entryPerPage)
	renderer.RenderTable(nftList, nData, pageNumber, entryPerPage, maxPage)

	switch menuNumber {
	case 1:
		PortMux1(&nftList, nData, pageNumber, maxPage)
	case 2:
		PortMux2(&nftList, nData, pageNumber, maxPage)
	case 3:
		PortMux3(&nftList, nData, pageNumber, maxPage)
	}
}
