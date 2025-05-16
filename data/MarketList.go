package data

import (
	"fmt"
	"math/rand"
	"time"
)

type nft struct {
	ID         int
	Name       string
	Creator    string
	Owner      string
	Blockchain string
	PriceETH   float64
	CreatedAt  string
	Royalty    float64
}

const NMAX = 1000

var NftList [NMAX]nft

func InitiateMarket() {
	// Preparing NftList array with 1000 entries

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var blockchains = [4]string{"Ethereum", "Polygon", "Solana", "BSC"}

	var nftNames = [36]string{
		"Meta Relic",
		"Ethereal Bloom",
		"Void Runner",
		"Pixel Propher",
		"Celestial Spark",
		"Chain Phantom",
		"Arcane Ember",
		"Quantum Fang",
		"Nova Shard",
		"Starlight Rune",
		"Blade of Echoes",
		"Aether Prism",
		"Neon Dagger",
		"Solar Grimoire",
		"Obsidian Warden",
		"Spectral Helm",
		"Drift Core",
		"Cosmic Elixir",
		"Phantom Medallion",
		"Crimson Module",
		"Radiant Flux",
		"Stormforged Coin",
		"Pulse Saber",
		"Warp Totem",
		"Shattered Glyph",
		"Graviton Claw",
		"Frostfire Bloom",
		"Mirage Chip",
		"Ancient Sparkstone",
		"Twilight Reactor",
		"Netherstone Ring",
		"Temporal Blade",
		"Lumina Relic",
		"Echo Core",
		"Runebound Pendant",
		"Starforged Mirror",
	}

	var creators = [6]string{
		"Ichinomiya",
		"Mike Smith",
		"Abdul Bahri",
		"Ende' Tu'ar",
		"Rafinha B.",
		"David R.",
	}

	var owners = [20]string{
		"William H.",
		"Emily B.",
		"Charles D.",
		"Grace L.",
		"Henry M.",
		"Alice T.",
		"Thomas W.",
		"Zayd A.",
		"Layla S.",
		"Omar K.",
		"Amina Z.",
		"Yusuf R.",
		"Hana J.",
		"Elowen F.",
		"Kaelthor V.",
		"Nymeria N.",
		"Thalor Q.",
		"Seraphina Y.",
		"Draven C.",
		"Aeloria E.",
	}

	for i := 0; i < 1000; i++ {
		randomDate := time.Now().AddDate(0, 0, -rng.Intn(1000))
		NftList[i] = nft{
			ID:         i + 1,
			Name:       fmt.Sprintf("%s #%d", nftNames[rng.Intn(len(nftNames))], i+1),
			Creator:    creators[rng.Intn(len(creators))],
			Owner:      owners[rng.Intn(len(owners))],
			Blockchain: blockchains[rng.Intn(len(blockchains))],
			PriceETH:   rng.Float64()*10 + 0.1,
			CreatedAt:  randomDate.Format("02-01-2006"),
			Royalty:    rng.Float64() - (rng.Float64() * 0.1),
		}
	}
}
