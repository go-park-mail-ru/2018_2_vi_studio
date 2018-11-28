package game

import "math/rand"

type Tile [7]int

const (
	FirstTilesCount  = 6
	SecondTilesCount = 14
	ThirdTilesCount  = 14
	FourthTilesCount = 6
	FifthTilesCount  = 14
)

var (
	FirstTile  = Tile{0, 1, 0, 3, 2, 5, 4}
	SecondTile = Tile{1, 3, 2, 1, 0, 5, 4}
	ThirdTile  = Tile{2, 1, 0, 4, 5, 2, 3}
	FourthTile = Tile{3, 3, 4, 5, 0, 1, 2}
	FifthTile  = Tile{4, 3, 5, 4, 0, 2, 1}
)

func NewTileSlice() []Tile {
	tiles := make([]Tile, 0, FirstTilesCount+SecondTilesCount+ThirdTilesCount+FourthTilesCount+FifthTilesCount)

	for i:= 0; i < FirstTilesCount; i++ {
		tiles = append(tiles, FirstTile)
	}

	for i:= 0; i < SecondTilesCount; i++ {
		tiles = append(tiles, SecondTile)
	}

	for i:= 0; i < ThirdTilesCount; i++ {
		tiles = append(tiles, ThirdTile)
	}

	for i:= 0; i < FourthTilesCount; i++ {
		tiles = append(tiles, FourthTile)
	}

	for i:= 0; i < FifthTilesCount; i++ {
		tiles = append(tiles, FifthTile)
	}

	// shuffle
	for i := range tiles {
		j := rand.Intn(i + 1)
		tiles[i], tiles[j] = tiles[j], tiles[i]
	}

	return tiles
}
