package utils

import (
	"app/model"
)

func SearchID(arr model.TabNFT, nData int16, target uint16) int16 {
	switch model.MarketSortCode {
	case 1:
		return BSearchAscID(arr, nData, uint16(target))
	case 2:
		return BSearchDscID(arr, nData, uint16(target))
	default:
		return LSearchID(arr, nData, uint16(target))
	}
}

func BSearchAscID(arr model.TabNFT, nData int16, targetID uint16) int16 {
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

func BSearchDscID(arr model.TabNFT, nData int16, targetID uint16) int16 {
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

func LSearchID(arr model.TabNFT, nData int16, targetID uint16) int16 {
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

func LSearchName(arr model.TabNFT, nData int16, targetName string) int16 {
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
