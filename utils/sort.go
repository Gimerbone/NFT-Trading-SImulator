package utils

import (
	"app/model"
	"time"
)

func SortIDAsc(arr *model.TabNFT, n int16) {
	var (
		i, j int16
		key  model.Nft
	)

	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		for j > -1 && key.ID < (*arr)[j].ID {
			(*arr)[j+1] = (*arr)[j]
			j--
		}

		(*arr)[j+1] = key
	}
}

func SortIDDsc(arr *model.TabNFT, n int16) {
	var (
		i, j int16
		key  model.Nft
	)

	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		for j > -1 && key.ID > (*arr)[j].ID {
			(*arr)[j+1] = (*arr)[j]
			j--
		}

		(*arr)[j+1] = key
	}
}

func SortPriceAsc(arr *model.TabNFT, n int16) {
	var (
		i, j int16
		key  model.Nft
	)

	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		for j > -1 && key.PriceETH < (*arr)[j].PriceETH {
			(*arr)[j+1] = (*arr)[j]
			j--
		}

		(*arr)[j+1] = key
	}
}

func SortPriceDsc(arr *model.TabNFT, n int16) {
	var (
		i, j int16
		key  model.Nft
	)

	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		for j > -1 && key.PriceETH > (*arr)[j].PriceETH {
			(*arr)[j+1] = (*arr)[j]
			j--
		}

		(*arr)[j+1] = key
	}
}

func SortDateAsc(arr *model.TabNFT, n int16) {
	var (
		i, j               int16
		key                model.Nft
		layout             string
		timeKey, timeIndex time.Time
	)

	layout = "02-01-2006"
	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		timeKey, _ = time.Parse(layout, key.CreatedAt)
		timeIndex, _ = time.Parse(layout, (*arr)[j].CreatedAt)

		for j > -1 && timeKey.Before(timeIndex) {
			(*arr)[j+1] = (*arr)[j]
			timeIndex, _ = time.Parse(layout, (*arr)[j].CreatedAt)
			j--
		}

		(*arr)[j+1] = key
	}
}

func SortDateDsc(arr *model.TabNFT, n int16) {
	var (
		i, j               int16
		key                model.Nft
		layout             string
		timeKey, timeIndex time.Time
	)

	layout = "02-01-2006"
	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		timeKey, _ = time.Parse(layout, key.CreatedAt)
		timeIndex, _ = time.Parse(layout, (*arr)[j].CreatedAt)

		for j > -1 && timeKey.After(timeIndex) {
			(*arr)[j+1] = (*arr)[j]
			timeIndex, _ = time.Parse(layout, (*arr)[j].CreatedAt)
			j--
		}

		(*arr)[j+1] = key
	}
}

func SortRoyaltyAsc(arr *model.TabNFT, n int16) {
	var (
		i, j int16
		key  model.Nft
	)

	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		for j > -1 && key.Royalty < (*arr)[j].Royalty {
			(*arr)[j+1] = (*arr)[j]
			j--
		}

		(*arr)[j+1] = key
	}
}

func SortRoyaltyDsc(arr *model.TabNFT, n int16) {
	var (
		i, j int16
		key  model.Nft
	)

	for i = 1; i < n; i++ {
		key = (*arr)[i]
		j = i - 1

		for j > -1 && key.Royalty > (*arr)[j].Royalty {
			(*arr)[j+1] = (*arr)[j]
			j--
		}

		(*arr)[j+1] = key
	}
}
