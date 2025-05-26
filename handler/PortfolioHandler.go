package handler

import (
	"app/data"
	"app/utils"
)

var boughtNFTlist [data.NMAX]data.Nft
var nData int

func purchaseNFT(boughtID int16) int8 {
	// Returns 0 if success, -1 if boughtID not found

	var (
		idx int16
	)

	idx = utils.SearchMarketID(originalList, nOriginalData, uint16(boughtID))

	if idx == -1 {
		return -1
	} else {
		boughtNFTlist[nData] = originalList[idx]
		nData++

		delNFT(idx)

		return 0
	}
}

func delNFT(delIndex int16) {
	var (
		i int16
	)

	for i = delIndex; i < nOriginalData-1; i++ {
		originalList[i] = originalList[i+1]
	}

	nOriginalData--
}
