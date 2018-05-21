package main

import (
	"github.com/FabianK1991/civilization/simulation/world"
	"github.com/FabianK1991/civilization/simulation"
	"github.com/FabianK1991/civilization/simulation/population"
	"github.com/FabianK1991/civilization/api"
)

func startup() *simulation.Simulation {
	person := population.CreatePerson(0, 0, "Fabian", "Test", population.GENDER_MALE)

	sim := simulation.Simulation{
		World: world.GenerateWorld(),
		Population: map[string]*population.Person{
			person.UUID: person,
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