package simulation

import (
	"io/ioutil"
	"encoding/json"
)

func GenerateWorld() World {
	var world World = make(World)

	for x:=0;x<10;x++ {
		world[x] = make(map[int]*Tile)

		for y:=0;y<10;y++ {
			world[x][y] = &Tile{
				X: x,
				Y: y,

				isBlocked: false,
				tileType: TILE_GRASS,
				world: world,
			}
		}
	}

	return world
}

func LoadFromFile(filename string) (World, error) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return GenerateWorld(), nil
	} else {
		var world World

		err := json.Unmarshal(data, &world)

		if err != nil {
			return nil, err
		} else {
			for _, row := range world {
				for _, tile := range row {
					tile.world = world
				}
			}

			return world, nil
		}
	}
}

func (world World) SaveToFile(filename string) error {
	data, err := json.Marshal(world); if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644); if err != nil {
		return err
	}

	return nil
}