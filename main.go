package main

import (
	"fmt"
	"github.com/FabianK1991/civilization/simulation/world"
	"github.com/FabianK1991/civilization/simulation"
)

func startup() {
	//var world simulation.World = simulation.GenerateWorld()
	//world.SaveToFile("world.json")

	fmt.Println(s[1][5])
}

func main() {
	// go startup()
	// api.StartApi()

	sim := simulation.Simulation{
		World: world.GenerateWorld(),
	}

	sim.World[0][0].TileType = world.TILE_WATER

	sim.SaveToFile("sim.json")

	//sim, _ := simulation.LoadFromFile("sim.json");

	//fmt.Println(sim.World)

	//startup()
}