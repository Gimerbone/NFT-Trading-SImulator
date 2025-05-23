package utils

import "fmt"

func ScanSentence(str *string) {
	// Scanning words repeatedly until '.' and storing it result into string

	var temp string

	fmt.Scanf("%s\n", &temp)
	*str = temp

	for temp[len(temp)-1] != '.' {
		fmt.Scan(&temp)
		if temp != " " {
			*str = *str + " " + temp
		}
	}
}
