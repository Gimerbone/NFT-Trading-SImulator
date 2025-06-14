package utils

import (
	"app/model"
	"math"
	"math/rand"
	"time"
)

var Rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomizePrice(nftList *model.TabNFT, nData int16) {
	var i, cases int16

	cases = int16(Rng.Intn(5))

	switch cases {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		for i = 0; i < nData; i++ {
			nftList[i].PriceETH = nftList[i].PriceETH - TruncateFloat(Rng.Float32()) + 0.5
		}
	case 4:
		for i = 0; i < nData; i++ {
			nftList[i].PriceETH = TruncateFloat(Rng.Float32()*10 + 0.1)
		}
	}
}

func RandomizeSelling(bidPrice float32, product model.Nft) bool {
	var (
		priceDiff                 float32
		successChance, failChance float64
	)

	failChance = Rng.Float64() * 100.0

	priceDiff = product.PriceETH - bidPrice
	successChance = 100.0 / (1 + math.Exp(-1.0*float64(priceDiff)/10))

	return successChance > failChance
}
