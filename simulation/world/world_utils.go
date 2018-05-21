package world

import (
	"github.com/FabianK1991/civilization/simulation/population"
	"math"
)

func WorldHasField(world World, x int, y int) bool {
	if(x < len(world) && y < len(world[0])) {
		return true;
	}

	return false
}

func neighborsAreUnblocked(x int, y int, world World) bool {
	for i:=-1;i<=1;i+=1 {
		for j:=-1;j<=1;j+=1 {
			if i == 0 && j == 0 {
				if world[x + i][y + j].IsBlocked == true {
					return false
				}

				continue;
			} else if WorldHasField(world, x + i, y + j) && world[x + i][y + j].IsBlocked == true {
				return false
			}
		}
	}

	return true
}

/**
	Buildable tile means that the neighbors of the tile are all not blocked
 */
func FindClosestBuildableTile(person *population.Person, world World) (x int, y int, found bool) {
	for distance:=1;distance<len(world);distance++ {
		startX := person.X - distance
		startY := person.Y - distance
		endX := startX + distance * 2
		endY := startY + distance * 2

		shouldBreak := true

		if WorldHasField(world, startX, startY) || WorldHasField(world, startX, endY) {
			for x, y := startX, startY;y <= endY;y++ {
				if WorldHasField(world, x, y) && neighborsAreUnblocked(x, y, world) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if WorldHasField(world, startX, startY) || WorldHasField(world, endX, startY) {
			for x, y := startX, startY;x <= endX;x++ {
				if WorldHasField(world, x, y) && neighborsAreUnblocked(x, y, world) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if WorldHasField(world, endX, startY + 1) || WorldHasField(world, endX, endY) {
			for x, y := endX, startY + 1;y <= endY;y++ {
				if WorldHasField(world, x, y) && neighborsAreUnblocked(x, y, world) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if WorldHasField(world, startX + 1, endY) || WorldHasField(world, endX, endY) {
			for x, y := startX + 1, endY;x <= endX;x++ {
				if WorldHasField(world, x, y) && neighborsAreUnblocked(x, y, world) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if shouldBreak {
			break;
		}
	}

	return 0, 0, false
}

func containsResource(s []Resources, e Resources) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func QuickDistanceCoordinates(fromX int, fromY int, toX int, toY int) float64 {
	dx := math.Abs(float64(fromX - toX))
	dy := math.Abs(float64(fromY - toY))

	return dx + dy - math.Min(dx, dy)
}

func QuickDistance(from *Tile, to *Tile, world World) float64 {
	return from.PathEstimatedCost(to)
}

func FindClosestResource(person *population.Person, resource Resources, world World) (x int, y int, found bool) {
	for distance:=1;distance<len(world);distance++ {
		startX := person.X - distance
		startY := person.Y - distance
		endX := startX + distance * 2
		endY := startY + distance * 2

		shouldBreak := true

		if WorldHasField(world, startX, startY) || WorldHasField(world, startX, endY) {
			for x, y := startX, startY;y <= endY;y++ {
				if WorldHasField(world, x, y) && containsResource(world[x][y].Resources, resource) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if WorldHasField(world, startX, startY) || WorldHasField(world, endX, startY) {
			for x, y := startX, startY;x <= endX;x++ {
				if WorldHasField(world, x, y) && containsResource(world[x][y].Resources, resource) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if WorldHasField(world, endX, startY + 1) || WorldHasField(world, endX, endY) {
			for x, y := endX, startY + 1;y <= endY;y++ {
				if WorldHasField(world, x, y) && containsResource(world[x][y].Resources, resource) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if WorldHasField(world, startX + 1, endY) || WorldHasField(world, endX, endY) {
			for x, y := startX + 1, endY;x <= endX;x++ {
				if WorldHasField(world, x, y) && containsResource(world[x][y].Resources, resource) {
					return x, y, true
				}
			}

			shouldBreak = false
		}

		if shouldBreak {
			break;
		}
	}

	return 0, 0, false
}