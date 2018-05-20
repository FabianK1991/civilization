package world

func GenerateWorld() World {
	var world World = make(World)

	for x:=0;x<10;x++ {
		world[x] = make(map[int]*Tile)

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

	return world
}