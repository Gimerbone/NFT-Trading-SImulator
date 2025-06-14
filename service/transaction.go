package service

import (
	"app/model"
	"app/renderer"
	"app/storage"
	"app/utils"
)

func purchaseNFT(boughtID uint16) int8 {
	// Returns -1 if money is not enough
	var idx int16 = utils.SearchID(storage.MarketList, storage.NMarketData, boughtID)

	if idx != -1 {
		if storage.MarketList[idx].PriceETH > model.User.BalanceETH {
			return -1
		}

		if storage.NPortData < model.NMAX {
			storage.PortList[storage.NPortData] = storage.MarketList[idx]
			storage.PortList[storage.NPortData].Owner = model.User.Name
			storage.NPortData++

			model.User.BalanceETH = model.User.BalanceETH - storage.MarketList[idx].PriceETH

			delMarketNFT(idx)

			return 0
		}
	}

	return 0
}

func delMarketNFT(delIndex int16) {
	var (
		i int16
	)

	for i = delIndex; i < storage.NMarketData-1; i++ {
		storage.MarketList[i] = storage.MarketList[i+1]
	}

	storage.NMarketData--
}

func enlistNFT() {
	var index int16
	renderer.ClearScreen()
	utils.InputID(&index)

	if storage.NSellOrderList < model.NMAX {
		storage.SellOrderList[storage.NSellOrderList] = storage.PortList[index]
	}
	delMarketNFT(index)
}

func handleSelling() {
	var i int16
	for i = 0; i < storage.NSellOrderList; i++ {
		if utils.RandomizeSelling(storage.BidPriceList[i], storage.SellOrderList[i]) {
			model.User.BalanceETH = model.User.BalanceETH + storage.SellOrderList[i].PriceETH
			delSONFT(i)
		}
	}
}

func delSONFT(delIndex int16) {
	var i int16

	for i = delIndex; i < storage.NSellOrderList-1; i++ {
		storage.SellOrderList[i] = storage.SellOrderList[i+1]
	}

	storage.NSellOrderList--
}

func CancelSellOrder() {
	var index int16
	renderer.ClearScreen()
	utils.InputID(&index)

	if storage.NPortData < model.NMAX {
		storage.PortList[storage.NPortData] = storage.SellOrderList[index]
	}
	delSONFT(index)
}
