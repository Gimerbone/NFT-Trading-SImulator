package utils

import "fmt"

func InputID(resultId *int16) {
	fmt.Printf("Enter nft id or -1 to exit:")
	fmt.Scanln(resultId)
}

func InputName(ctx string, canQuit bool, result *string) {
	fmt.Printf("Enter %s (you must end it with '.'", ctx)
	if canQuit {
		fmt.Printf(" or -1 to exit): ")
	} else {
		fmt.Printf("): ")
	}
	scanSentence(result)
}

func scanSentence(str *string) {
	// Scanning words repeatedly until '.' and storing it result into string
	var (
		temp   byte
		result []byte
		argc   int
	)

	fmt.Scanf("%c", &temp)
	for temp != '.' {
		if temp != '\n' && temp != '\r' {
			result = append(result, temp)
		}
		fmt.Scanf("%c", &temp)
	}

	for {
		argc, _ = fmt.Scanf("%c", &temp)
		if argc == 0 || temp == '\n' {
			break
		}
	}

	*str = string(result)
}
