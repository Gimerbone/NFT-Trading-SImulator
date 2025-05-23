package hmap

// This HashMap is specifically developed for market list item.
// The key and value are like this map[key: string] -> bool.
// It's optimized for [1, 12] char length string, returns true if found.

const TABLE_SIZE uint32 = 1024

type HashMap [TABLE_SIZE]string

func hash(key string) uint32 {
	var (
		i    int
		hash uint32
	)

	hash = 0
	for i = 0; i < len(key); i++ {
		hash = hash*31 + uint32(key[i])
	}

	return hash & 1023
}

func IsKeyPresent(hmap HashMap, key string) bool {
	var i, idx, first_position uint32

	first_position = hash(key)
	for i = 0; i < TABLE_SIZE; i++ {
		idx = (first_position + i) & 1023
		if hmap[idx] == "" {
			return false
		}
		if hmap[idx] == key {
			return true
		}
	}

	return false
}

func AddKey(hmap *HashMap, key string) {
	var i uint32

	i = hash(key)
	if !IsKeyPresent(*hmap, key) {
		(*hmap)[i] = key
	} else {
		for {
			i = (i + 1) & 1023
			if hmap[i] != "" {
				break
			}
		}
		hmap[i] = key
	}
}
