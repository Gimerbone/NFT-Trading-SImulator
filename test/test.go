package main

import "fmt"

const PRIME uint32 = 31

func main() {
	names := []string{
		"William H.", "Emily B.",
		"Charles D.", "Grace L.",
		"Henry M.", "Alice T.",
		"Jacob",
	}

	for _, name := range names {
		fmt.Printf("%s: %d\n", name, hash(name))
	}
}

func hash(key string) uint32 {
	var (
		i    int
		hash uint32
	)

	hash = 0
	for i = 0; i < len(key); i++ {
		hash = hash*PRIME + uint32(key[i])
	}

	return hash & 1023
}
