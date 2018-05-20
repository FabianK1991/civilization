package main

import (
	"fmt"
	"github.com/FabianK1991/civilization/simulation"
	"github.com/beefsack/go-astar"
	"strconv"
)


func main() {
	var world simulation.World = simulation.GenerateWorld()

	path, _, _ := astar.Path(world[0][0], world[8][9])

	for i:=len(path) - 1;i>=0;i-- {
		x := path[i].(*simulation.Tile).X
		y := path[i].(*simulation.Tile).Y

		fmt.Println("Take: " + strconv.Itoa(x) + "," + strconv.Itoa(y))
	}
}