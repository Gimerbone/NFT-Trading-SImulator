package utils

import (
	"app/data"
)

func FilterByBlockchain(bcName string, arr data.TabNFT, nData int16, newArr *data.TabNFT, newN *int16) {
	var i int16
	*newN = 0
	for i = 0; i < nData; i++ {
		if arr[i].Blockchain == bcName {
			(*newArr)[*newN] = arr[i]
			*newN++
		}
	}
}

func FilterByCreator(oName string, arr data.TabNFT, nData int16, newArr *data.TabNFT, newN *int16) {
	var i int16
	*newN = 0
	for i = 0; i < nData; i++ {
		if arr[i].Creator == oName {
			(*newArr)[*newN] = arr[i]
			*newN++
		}
	}
}
