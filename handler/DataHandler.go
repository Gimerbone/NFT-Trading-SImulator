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
		originalList[i].PriceETH = utils.TruncateFloat(rng.Float64()*10 + 0.1)
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
