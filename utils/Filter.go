package utils

import (
	"app/data"
)

func FilterByBlockchain(blockchain string, arr data.TabNFT, nData int16, newArr *data.TabNFT, newN *int16) {
	var i int16
	*newN = 0
	for i = 0; i < nData; i++ {
		if arr[i].Blockchain == blockchain {
			(*newArr)[*newN] = arr[i]
		}
	}
}
