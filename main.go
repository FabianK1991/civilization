package main

import (
	"github.com/FabianK1991/civilization/simulation/world"
	"github.com/FabianK1991/civilization/simulation"
	"github.com/FabianK1991/civilization/simulation/population"
	"github.com/FabianK1991/civilization/api"
)

func startup() *simulation.Simulation {
	sim := simulation.Simulation{
		World: world.GenerateWorld(),
		Population: []*population.Person{
			{
				X: 0,
				Y: 0,
				Gender: population.GENDER_MALE,
				Prename: "Fabian",
				Surname: "Test",
			},
		},
	}

	sim.Init()
	//sim.SaveToFile("sim.json")

	return &sim
}

func main() {
	var sim = startup()

	api.StartApi(sim)
}