package simulation

func WorldHasField(world World, x int, y int) bool {
	if _, ok := world[x]; ok {
		if _, ok := world[x][y]; ok {
			return true
		}
	}

	return false
}
