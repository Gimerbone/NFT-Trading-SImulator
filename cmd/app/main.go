package main

import (
	"app/data"
	"app/renderer"
)

func main() {

}

func handleMain() {
	var (
		option     int8
		marketList data.TabNFT
	)

	data.InitiateMarket(&marketList)
	renderer.RenderMainMenu(&option)
	mainMux(option, marketList)
}

func mainMux(option int8, marketList data.TabNFT) {
	switch option {
	case 0:
		renderer.ClearScreen()
		renderer.RenderQuitMessage()
	case 1:
		handleMarket(marketList, data.NMAX, 1)
	case 2:
		portfolioHandler()
	}
}

func portfolioHandler() {}
