package utils

import "fmt"

func ScanSentence(str *string) {
	// Scanning words repeatedly until '.' and storing it result into string

	var temp string

	fmt.Scan(&temp)
	*str = temp

	for temp[len(temp)-1] != '.' {
		fmt.Scan(&temp)
		*str = *str + " " + temp
	}
}
