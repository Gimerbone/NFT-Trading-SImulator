package cmd

import (
	"app/model"
	"app/renderer"
	"app/service"
	"app/storage"
)

func Execute() {
	renderer.WelcomeScreen()
	storage.InitiateMarketList()
	model.User.BalanceETH = 10.0

	renderer.ClearScreen()
	service.MainMenu()
}
