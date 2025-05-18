package data

import (
	"fmt"
	"math/rand"
	"time"
)

type Nft struct {
	ID         uint16
	Name       string
	Creator    string
	Owner      string
	Blockchain string
	PriceETH   float64
	CreatedAt  string
	Royalty    float32
}

const NMAX = 999

type TabNFT [NMAX]Nft

/*
This variable tells TabNFT sort status
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
var StatusCode int8 = 1

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

var blockchains = [10]string{
	"Ethereum", "Polygon",
	"Solana", "BSC",
	"Avalanche", "Fantom",
	"Tezos", "Near",
	"Arbitrum", "Optimism",
}

var nftNames = [100]string{
	"Meta Relic", "Ethereal Bloom",
	"Void Runner", "Pixel Propher",
	"Celestial Spark", "Chain Phantom",
	"Arcane Ember", "Quantum Fang",
	"Nova Shard", "Starlight Rune",
	"Blade of Echoes", "Aether Prism",
	"Neon Dagger", "Solar Grimoire",
	"Obsidian Warden", "Spectral Helm",
	"Drift Core", "Cosmic Elixir",
	"Phantom Medallion", "Crimson Module",
	"Radiant Flux", "Stormforged Coin",
	"Pulse Saber", "Warp Totem",
	"Shattered Glyph", "Graviton Claw",
	"Frostfire Bloom", "Mirage Chip",
	"Ancient Sparkstone", "Twilight Reactor",
	"Netherstone Ring", "Temporal Blade",
	"Lumina Relic", "Echo Core",
	"Runebound Pendant", "Starforged Mirror",
	"Cyber Antler", "Dream Circuit",
	"Token Whisper", "Lunar Amulet",
	"Bio Synth", "Prism Pulse",
	"Hallowed Byte", "Phantom Badge",
	"Shimmer Chip", "Darklight Crest",
	"Genesis Mask", "Chainbound Ring",
	"Flare Tesseract", "Dusk Saber",
	"Solaris Gem", "Reactor Bloom",
	"Vortex Crown", "Nova Pendant",
	"Sparkshade", "Shadow Bit",
	"Dimensional Core", "Crypto Fang",
	"Chrono Medallion", "Pixel Ghoul",
	"Neon Ember", "Ether Idol",
	"Twilight Blade", "Vault Crystal",
	"Glitch Ember", "Echo Nova",
	"Fusion Crest", "Eclipse Rune",
	"Fractal Warden", "Warp Spark",
	"Null Beacon", "Arc Drift",
	"Chain Wisp", "Oblivion Chip",
	"Meta Halo", "Glimmer Core",
	"Xeno Prism", "Dust Circuit",
	"Rune Flare", "Solar Pulse",
	"Data Shard", "Spectrum Relic",
	"Pulse Totem", "Zephyr Node",
	"Fragment Token", "Phantom Totem",
	"Pixel Rune", "Spirit Code",
	"Quantum Halo", "Dreamflame Ring",
	"Lucid Mask", "Shroud Chip",
	"Token Bloom", "Stellar Crest",
	"Crypto Ember", "Starlit Module",
	"Chain Halo", "Nova Byte",
	"Neon Whisper", "Echo Tether",
}

var creators = [20]string{
	"LarvaLabs", "YugaLabs",
	"ArtBlocks", "OpenSea",
	"Manifold", "Zora",
	"Rarible", "Foundation",
	"Valve", "Nintendo",
	"FromSoft", "CDProjekt",
	"EpicGames", "Ubisoft",
	"Bungie", "Rockstar",
	"NaughtyDog", "Bethesda",
	"SuperRare", "NiftyGate",
}

var owners = [40]string{
	"William H.", "Emily B.",
	"Charles D.", "Grace L.",
	"Henry M.", "Alice T.",
	"Thomas W.", "Zayd A.",
	"Layla S.", "Omar K.",
	"Amina Z.", "Yusuf R.",
	"Hana J.", "Elowen F.",
	"Kaelthor V.", "Nymeria N.",
	"Thalor Q.", "Seraphina Y.",
	"Draven C.", "Aeloria E.",
	"Khalid N.", "Salma R.",
	"Faris H.", "Noor M.",
	"Tariq B.", "Maya L.",
	"Rami S.", "Lina K.",
	"Jamal D.", "Iman F.",
	"Samir A.", "Dana Y.",
	"Zain T.", "Rania W.",
	"Bilal J.", "Nadia Z.",
	"Hassan O.", "Amira E.",
	"Adil Q.", "Fatima C.",
}

func InitiateMarketList(nftList *TabNFT, nData *int16) {
	// Preparing nftList array with 999 entries
	// Should only be called from market handler once
	var i uint16
	for i = 0; i < 999; i++ {
		randomDate := time.Now().AddDate(0, 0, -1*rng.Intn(1000))

		(*nftList)[i].ID = i + 1
		(*nftList)[i].Name = fmt.Sprintf("%s #%d", nftNames[rng.Intn(len(nftNames))], rng.Intn(9000)+1000)
		(*nftList)[i].Creator = creators[rng.Intn(len(creators))]
		(*nftList)[i].Owner = owners[rng.Intn(len(owners))]
		(*nftList)[i].Blockchain = blockchains[rng.Intn(len(blockchains))]
		(*nftList)[i].PriceETH = rng.Float64()*10 + 0.1
		(*nftList)[i].CreatedAt = randomDate.Format("02-01-2006")
		(*nftList)[i].Royalty = rng.Float32() * 0.2
	}

	*nData = 999
}
