package renderer

import (
	"app/utils"
	"fmt"
)

func RenderIdSearch(resultId *int16) {
	for {
		fmt.Printf("Enter nft id or -1 to exit:")
		fmt.Scanln(resultId)

		if *resultId > -2 && *resultId < 1000 && *resultId != 0 {
			break
		}

		ClearScreen()
		fmt.Printf("%s\n%s\n",
			"ID doesn't exist. Choose between 1-1000 or -1 to exit.",
			"-------------------------------------------------------",
		)
	}
}

func RenderNameInput(objectName string, resultName *string, quitOption bool) {
	fmt.Printf("Enter %s (you must end it with '.'", objectName)
	if quitOption {
		fmt.Printf(" or -1 to exit): ")
	} else {
		fmt.Printf("): ")
	}
	utils.ScanSentence(resultName)
}
