package service

import (
	"app/model"
	"app/renderer"
	"app/storage"
	"app/utils"
)

func Portfolio(nftList model.TabNFT, nData int16, pageNumber int16, menuNumber int8) {
	var maxPage, entryPerPage int16

	entryPerPage = 15
	maxPage = utils.IntDivCeil(nData, entryPerPage)
	renderer.RenderPortTable(nftList, nData, pageNumber, entryPerPage, maxPage)

	switch menuNumber {
	case 1:
		PortMux1(&nftList, nData, pageNumber, maxPage)
	case 2:
		PortMux2(&nftList, nData, pageNumber, maxPage)
	case 3:
		PortMux3(&nftList, nData, pageNumber, maxPage)
	}
}

func calculateTotalValue() {
	var i int16
	storage.TotalPortValue = 0
	for i = 0; i < storage.NPortData; i++ {
		storage.TotalPortValue = storage.TotalPortValue + storage.PortList[i].PriceETH
	}
}
