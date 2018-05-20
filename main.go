package main

import (
	"github.com/FabianK1991/civilization/simulation/world"
	"github.com/FabianK1991/civilization/simulation"
	"github.com/FabianK1991/civilization/simulation/population"
)

func startup() {
	sim := simulation.Simulation{
		World: world.GenerateWorld(),
		Population: []*population.Person{
			&population.Person{
				X: 0,
				Y: 0,
				Gender: population.GENDER_MALE,
				Prename: "Fabian",
				Surname: "Test",
			},
		},
	}

	sim.Init()
	sim.SaveToFile("sim.json")
}

func main() {
	startup()

	// go startup()
	// api.StartApi()

	//sim, _ := simulation.LoadFromFile("sim.json");

	//fmt.Println(sim.World)

	//startup()
}