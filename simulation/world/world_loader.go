package world

func GenerateWorld() World {
	var world World = make(World, 10)

	for x:=0;x<10;x++ {
		world[x] = make([]*Tile, 10)

		for y:=0;y<10;y++ {
			world[x][y] = &Tile{
				X: x,
				Y: y,

				IsBlocked: false,
				TileType: TILE_GRASS,
				world: world,
			}
		}
	}

	world[4][5].TileType = TILE_WATER
	world[4][5].Resources = []Resources{RESOURCE_WATER}

	world[5][5].TileType = TILE_WATER
	world[5][5].Resources = []Resources{RESOURCE_FISH, RESOURCE_WATER}

	return world
}