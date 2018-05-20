package world

import (
	"github.com/beefsack/go-astar"
	"math"
)

type TileType int

const (
	TILE_UNKNOWN TileType = iota
	TILE_GRASS
	TILE_WATER
)

type Tile struct {
	X int
	Y int

	IsBlocked bool
	TileType TileType
	world World
}

// Implement Pather interface for the astar algorithm
func (tile*Tile) PathNeighbors() []astar.Pather {
	var retSlice []astar.Pather

	retSlice = make([]astar.Pather, 0, 8)

	for i:=-1;i<=1;i+=1 {
		for j:=-1;j<=1;j+=1 {
			if i == 0 && j == 0 {
				continue;
			} else if WorldHasField(tile.world, tile.X + i, tile.Y + j) {
				retSlice = append(retSlice, tile.world[tile.X + i][tile.Y + j])
			}
		}
	}

	return retSlice
}

func (tile*Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

func (tile*Tile) PathEstimatedCost(to astar.Pather) float64 {
	dx := math.Abs(float64(tile.X - to.(*Tile).X))
	dy := math.Abs(float64(tile.Y - to.(*Tile).Y))

	return dx + dy - math.Min(dx, dy)
}
