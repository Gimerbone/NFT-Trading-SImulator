package utils

import (
	"app/model"
	"time"
)

func FilterByBlockchain(bcName string, arr model.TabNFT, nData int16, newArr *model.TabNFT, newN *int16) {
	var i int16
	*newN = 0
	for i = 0; i < nData; i++ {
		if arr[i].Blockchain == bcName {
			(*newArr)[*newN] = arr[i]
			*newN++
		}
	}
}

func FilterByCreator(cName string, arr model.TabNFT, nData int16, newArr *model.TabNFT, newN *int16) {
	var i int16
	*newN = 0
	for i = 0; i < nData; i++ {
		if arr[i].Creator == cName {
			(*newArr)[*newN] = arr[i]
			*newN++
		}
	}
}

func FilterByYear(year string, arr model.TabNFT, nData int16, newArr *model.TabNFT, newN *int16) {
	var (
		i    int16
		date time.Time
	)
	*newN = 0
	for i = 0; i < nData; i++ {
		date, _ = time.Parse("02-01-2006", arr[i].CreatedAt)
		if date.Format("2006") == year {
			(*newArr)[*newN] = arr[i]
			*newN++
		}
	}
}

func FilterByOwner(oName string, arr model.TabNFT, nData int16, newArr *model.TabNFT, newN *int16) {
	var i int16
	*newN = 0
	for i = 0; i < nData; i++ {
		if arr[i].Owner == oName {
			(*newArr)[*newN] = arr[i]
			*newN++
		}
	}
}
