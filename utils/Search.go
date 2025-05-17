package utils

import (
	"app/data"
)

func findMinID(arr data.TabNFT, nData int16, idxStart int16) int16 {
	var min int16

	min = 0
	for i := idxStart; i < nData; i++ {
		if arr[min].ID > arr[i].ID {
			min = i
		}
	}
	return min
}
