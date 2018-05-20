package main

import (
	"github.com/FabianK1991/civilization/simulation"
	"fmt"
)

func main() {
	//var world simulation.World = simulation.GenerateWorld()
	//world.SaveToFile("world.json")
	world, _ := simulation.LoadFromFile("world.json")

	fmt.Println(world[0][0])
}