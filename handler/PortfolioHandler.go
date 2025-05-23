package handler

import (
	"app/data"
	"app/handler"
	"app/utils"
)

var boughtNFTlist [data.NMAX]data.Nft
var nData int

func buyNFT(boughtID int) {
	var (
		idx int
	)

	switch data.StatusCode {
	case 1:
		idx = utils.BSearchAscID(handler.OriginalList, handler.NOriginalData, uint16(target))
	case 2:
		idx = utils.BSearchDscID(handler.OriginalList,  handler.NOriginalData, uint16(target))
	default:
		idx = utils.LSearchID(handler.OriginalList,  handler.NOriginalData, uint16(target))
	}

	boughtNFTlist[nData] = handler.OriginalList[idx]
	nData++

	delNFT(idx)
}

func delNFT(delIndex int){
	var (
		i int
	)

	for i = delIndex;i < handler.NOriginalData-1; i++ {
		handler.OriginalList[i] = handler.OriginalList[i+1]
	}

	handler.NOriginalData--
}

func portfolioHandler() {

}
