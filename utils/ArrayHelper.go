package utils

import "app/data"

func swapTabNFT(a *data.Nft, b *data.Nft) {
	tmp := *a
	*a = *b
	*b = tmp
}
