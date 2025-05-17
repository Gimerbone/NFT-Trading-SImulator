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
