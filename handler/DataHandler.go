package handler

import (
	"app/data"
	"app/hmap"
	"app/utils"
	"fmt"
	"math/rand"
	"time"
)

var originalList data.TabNFT
var nOriginalData int16

var boughtNFTlist [data.NMAX]data.Nft
var nPortfolio int16

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func initiateMarketList() {
	// Preparing nftList array with 999 entries
	// Should only be called from market handler once
	var (
		nameList hmap.HashMap
		i        uint16
	)

	for i = 0; i < data.NMAX; i++ {
		randomDate := time.Now().AddDate(0, 0, -1*rng.Intn(1000))
		originalList[i].ID = i + 1
		originalList[i].Name = uniqueNameGenerator(&nameList)
		originalList[i].Creator = data.Creators[rng.Intn(len(data.Creators))]
		originalList[i].Owner = data.Owners[rng.Intn(len(data.Owners))]
		originalList[i].Blockchain = data.Blockchains[rng.Intn(len(data.Blockchains))]
		originalList[i].PriceETH = utils.TruncateFloat(rng.Float32()*10 + 0.1)
		originalList[i].CreatedAt = randomDate.Format("02-01-2006")
		originalList[i].Royalty = utils.TruncateFloat(rng.Float32() * 0.15)
	}

	nOriginalData = data.NMAX
}

func uniqueNameGenerator(nameList *hmap.HashMap) string {
	var name string

	name = fmt.Sprintf("%s #%d", data.NftNames[rng.Intn(len(data.NftNames))], rng.Intn(9000)+1000)
	if hmap.IsKeyPresent(*nameList, name) {
		return uniqueNameGenerator(nameList)
	} else {
		hmap.AddKey(nameList, name)
		return name
	}
}

func randomizePrice() {
	var i int16
	for i = 0; i < nOriginalData; i++ {
		originalList[i].PriceETH = utils.TruncateFloat(rng.Float32()*10 + 0.1)
	}
}

func purchaseNFT(boughtID int16) int8 {
	// Returns -1 if money is not enough
	var (
		idx int16
	)

	idx = utils.SearchMarketID(originalList, nOriginalData, uint16(boughtID))

	if idx != -1 {
		if originalList[idx].PriceETH > data.User.BalanceETH {
			return -1
		} else {
			boughtNFTlist[nPortfolio] = originalList[idx]
			boughtNFTlist[nPortfolio].Owner = data.User.Name
			nPortfolio++

			delNFT(idx)

			data.User.BalanceETH = data.User.BalanceETH - originalList[idx].PriceETH

			return 0
		}
	}

	return 0
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
