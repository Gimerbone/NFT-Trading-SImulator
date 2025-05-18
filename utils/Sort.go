package utils

import (
	"app/data"
)

func SortIDAsc(arr *data.TabNFT, n int16) {
	var i int16
	for i = 0; i < n; i++ {
		idx := findMinID(*arr, n, i)
		swapTabNFT(&(*arr)[i], &(*arr)[idx])
	}
}

func SortIDDsc(arr *data.TabNFT, n int16) {
	var i int16
	for i = 0; i < n; i++ {
		idx := findMaxID(*arr, n, i)
		swapTabNFT(&(*arr)[i], &(*arr)[idx])
	}
}

func SortPriceAsc(arr *data.TabNFT, n int16) {}

func SortPriceDsc(arr *data.TabNFT, n int16) {}

func SortDateAsc(arr *data.TabNFT, n int16) {}

func SortDateDsc(arr *data.TabNFT, n int16) {}

func SortRoyaltyAsc(arr *data.TabNFT, n int16) {}

func SortRoyaltyDsc(arr *data.TabNFT, n int16) {}
