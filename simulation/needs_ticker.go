package simulation

import "github.com/FabianK1991/civilization/simulation/population"

func decideNextAction(person *population.Person, sim *Simulation) {
	if(person.CurrentAction == population.ACTIONS_IDLE) {

	}
}

func tickNeeds(person *population.Person, sim *Simulation) {
	tickSleep(person, sim)
	tickHunger(person, sim)
	tickThirst(person, sim)
}

func tickSleep(person *population.Person, sim *Simulation) {
	if(person.CurrentAction == population.ACTIONS_SLEEPING) {
		person.Needs[population.NEEDS_ENERGY] += 12.5 / float64(TICKS_PER_HOUR)

		return
	}

	// Tick stronger between 0am-5am
	if(sim.GetDayTime() < TICKS_PER_HOUR * 5) {
		person.Needs[population.NEEDS_ENERGY] -= 2
	} else {
		person.Needs[population.NEEDS_ENERGY] -= 0.5
	}
}

func tickHunger(person *population.Person, sim *Simulation) {
	if(person.CurrentAction == population.ACTIONS_SLEEPING) {
		person.Needs[population.NEEDS_HUNGER] -= 0.125
	} else {
		person.Needs[population.NEEDS_HUNGER] -= 0.25
	}
}

func tickThirst(person *population.Person, sim *Simulation) {
	if(person.CurrentAction == population.ACTIONS_SLEEPING) {
		person.Needs[population.NEEDS_THIRST] -= 0.25
	} else {
		person.Needs[population.NEEDS_THIRST] -= 0.5
	}
}