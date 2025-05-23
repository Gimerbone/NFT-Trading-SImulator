package utils

import (
	"app/data"
	"time"
)

func SortIDAsc(arr *data.TabNFT, n int16) {
	var (
		i, j int16
		key  data.Nft
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

func SortIDDsc(arr *data.TabNFT, n int16) {
	var (
		i, j int16
		key  data.Nft
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

func SortPriceAsc(arr *data.TabNFT, n int16) {
	var (
		i, j int16
		key  data.Nft
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

func SortPriceDsc(arr *data.TabNFT, n int16) {
	var (
		i, j int16
		key  data.Nft
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

func SortDateAsc(arr *data.TabNFT, n int16) {
	var (
		i, j               int16
		key                data.Nft
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

func SortDateDsc(arr *data.TabNFT, n int16) {
	var (
		i, j               int16
		key                data.Nft
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

func SortRoyaltyAsc(arr *data.TabNFT, n int16) {
	var (
		i, j int16
		key  data.Nft
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

func SortRoyaltyDsc(arr *data.TabNFT, n int16) {
	var (
		i, j int16
		key  data.Nft
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
