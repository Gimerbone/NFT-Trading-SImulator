package utils

import "fmt"

func ScanSentence(str *string) {
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
