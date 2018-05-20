package world

func WorldHasField(world World, x int, y int) bool {
	if(x < len(world) && y < len(world[0])) {
		return true;
	}

	return false
}
