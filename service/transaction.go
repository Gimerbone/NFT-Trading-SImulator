package service

import (
	"app/model"
	"app/storage"
	"app/utils"
)

func purchaseNFT(boughtID uint16) int8 {
	// Returns -1 if money is not enough
	var idx int16 = utils.SearchID(storage.MarketList, storage.NMarketData, boughtID)

	if idx != -1 {
		if storage.MarketList[idx].PriceETH > model.User.BalanceETH {
			return -1
		} else {
			storage.PortList[storage.NPortData] = storage.MarketList[idx]
			storage.PortList[storage.NPortData].Owner = model.User.Name
			storage.NPortData++

			delNFT(idx)

			model.User.BalanceETH = model.User.BalanceETH - storage.MarketList[idx].PriceETH

			return 0
		}
	}

	return 0
}

func delNFT(delIndex int16) {
	var (
		i int16
	)

	for i = delIndex; i < storage.NMarketData-1; i++ {
		storage.MarketList[i] = storage.MarketList[i+1]
	}

	storage.NMarketData--
}
