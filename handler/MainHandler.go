package handler

import (
	"app/data"
	"app/renderer"
)

func StartApp() {
	handleWelcome()
	initiateMarketList()
	data.User.BalanceETH = 10.0

	renderer.ClearScreen()
	handleMain()
}

func handleMain() {
	var (
		option int8
	)

	renderer.RenderWelcome()
	renderer.RenderMainMenu(&option)
	mainMux(option)
}

func mainMux(option int8) {
	switch option {
	case 0:
		renderer.ClearScreen()
		renderer.RenderQuitMessage()
	case 1:
		renderer.ClearScreen()
		handleMarket(originalList, nOriginalData, 1, 1)
	default:
		renderer.ClearScreen()
		renderer.RenderOptionNotExist()
		handleMain()
	}
}

func handleWelcome() {
	var (
		username string
	)

	renderer.ClearScreen()
	for {
		renderer.RenderUsernameInput(&username)
		if len(username) < 11 && username != "" {
			data.User.Name = username
			break
		}
		renderer.ClearScreen()
		renderer.RenderInvalidName()
	}
}
