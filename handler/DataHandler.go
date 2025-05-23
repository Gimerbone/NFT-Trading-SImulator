package handler

import (
	"app/data"
	"app/hmap"
	"app/utils"
	"fmt"
	"math/rand"
	"time"
)

var OriginalList data.TabNFT
var NOriginalData int16

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
		OriginalList[i].ID = i + 1
		OriginalList[i].Name = uniqueNameGenerator(&nameList)
		OriginalList[i].Creator = data.Creators[rng.Intn(len(data.Creators))]
		OriginalList[i].Owner = data.Owners[rng.Intn(len(data.Owners))]
		OriginalList[i].Blockchain = data.Blockchains[rng.Intn(len(data.Blockchains))]
		OriginalList[i].PriceETH = utils.TruncateFloat(rng.Float64()*10 + 0.1)
		OriginalList[i].CreatedAt = randomDate.Format("02-01-2006")
		OriginalList[i].Royalty = utils.TruncateFloat(rng.Float32() * 0.15)
	}

	NOriginalData = data.NMAX
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
