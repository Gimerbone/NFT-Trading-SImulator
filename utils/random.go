package utils

import (
	"app/model"
	"math/rand"
	"time"
)

var Rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomizePrice(nftList *model.TabNFT, nData int16) {
	var i int16
	for i = 0; i < nData; i++ {
		nftList[i].PriceETH = TruncateFloat(Rng.Float32()*10 + 0.1)
	}
}
