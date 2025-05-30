package model

type Nft struct {
	ID         uint16
	Name       string
	Creator    string
	Owner      string
	Blockchain string
	PriceETH   float32
	CreatedAt  string
	Royalty    float32
}

const NMAX = 1024

type TabNFT [NMAX]Nft

/*
This variable tells Market passed data sort status
- 0: Unsorted.
- 1: Sorted by ID asc.
- 2: Sorted by ID dsc.
- 3: Sorted by Price asc.
- 4: Sorted by Price dsc.
- 5: Sorted by Date asc.
- 6. Sorted by Date dsc.
- 7. Sorted by Royalty asc.
- 8. Sorted by Royalty dsc.
Default is one
*/
var MarketSortCode int8 = 1

/*
This variable tells Portfolio passed data sort status
- 0: Unsorted.
- 1: Sorted by ID asc.
- 2: Sorted by ID dsc.
- 3: Sorted by Price asc.
- 4: Sorted by Price dsc.
- 5: Sorted by Date asc.
- 6. Sorted by Date dsc.
- 7. Sorted by Royalty asc.
- 8. Sorted by Royalty dsc.
- 9. Sorted by Bought Time asc.
Default is one
*/
var PortSortCode int8 = 0
