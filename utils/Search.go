package utils

import (
	"app/data"
)

func findMinID(arr data.TabNFT, nData int16, idxStart int16) int16 {
	var min, i int16

	min = idxStart
	for i = idxStart; i < nData; i++ {
		if arr[min].ID > arr[i].ID {
			min = i
		}
	}
	return min
}

func findMaxID(arr data.TabNFT, nData int16, idxStart int16) int16 {
	var max, i int16

	max = idxStart
	for i = idxStart; i < nData; i++ {
		if arr[max].ID < arr[i].ID {
			max = i
		}
	}
	return max
}

func BSearchAscID(arr data.TabNFT, nData int16, targetID uint16) int16 {
	var (
		low, high, mid, idx int16
	)

	low = 0
	high = nData - 1
	idx = -1

	for (low <= high) && (idx == -1) {
		mid = (low + high) / 2
		if arr[mid].ID < targetID {
			low = mid + 1
		} else if arr[mid].ID > targetID {
			high = mid - 1
		} else {
			idx = mid
		}
	}

	return idx
}

func BSearchDscID(arr data.TabNFT, nData int16, targetID uint16) int16 {
	var (
		low, high, mid, idx int16
	)

	low = 0
	high = nData - 1
	idx = -1

	for (low <= high) && (idx == -1) {
		mid = (low + high) / 2
		if arr[mid].ID > targetID {
			low = mid + 1
		} else if arr[mid].ID < targetID {
			high = mid - 1
		} else {
			idx = mid
		}
	}

	return idx
}

func LSearchID(arr data.TabNFT, nData int16, targetID uint16) int16 {
	var i int16 = 0
	for i < nData && targetID != arr[i].ID {
		i++
	}

	if i < nData {
		return i
	} else {
		return -1
	}
}

func LSearchName(arr data.TabNFT, nData int16, targetName string) int16 {
	var i int16 = 0
	
	for i < nData && targetName != arr[i].Name {
		i++
	}

	if i < nData {
		return i
	} else {
		return -1
	}
}
