package world

import (
	"github.com/FabianK1991/civilization/simulation/population"
)

type World [][]*Tile

func (world World) Init(population population.Population) {
	for x, row := range world {
		for y, col := range row {
			col.world = world

			for _, person := range population {
				if person.X == x && person.Y == y {
					col.persons = append(col.persons, person)
				}
			}
		}
	}
}