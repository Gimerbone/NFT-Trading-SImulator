package renderer

import (
	"app/data"
	"fmt"
)

func RenderPortfolio(data data.TabNFT, pageNumber, entryPerPage, maxPage, nData int16) {
	fmt.Println("===================================================================================================================")
	fmt.Printf(" |  %s  |           %s           |    %s     | %-10s | %-12s | %-10s | %-7s |\n",
		"ID", "Name", "Creator", "Blockchain", "Market Price", "Created At", "Royalty",
	)
	fmt.Println(" ----------------------------------------------------------------------------------------------------------------- ")

	renderPortfolioList(data, (pageNumber-1)*entryPerPage, pageNumber*entryPerPage, nData)

	fmt.Println("===================================================================================================================")
	fmt.Printf("  PAGES %d/%d\n", pageNumber, maxPage)
	fmt.Println("===================================================================================================================")
}

func renderPortfolioList(arr data.TabNFT, entryStart, entryEnd, nData int16) {
	if entryEnd > nData {
		entryEnd = nData
	}

	for i := entryStart; i < entryEnd; i++ {
		fmt.Printf(" | %4d | %-24s | %-12s | %-10s | %12.2f | %10s | %5.0f %% |\n",
			arr[i].ID, arr[i].Name, arr[i].Creator,
			arr[i].Blockchain, arr[i].PriceETH, arr[i].CreatedAt,
			arr[i].Royalty*100,
		)
	}
}
