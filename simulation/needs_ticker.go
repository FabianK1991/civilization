package simulation

import (
	"github.com/FabianK1991/civilization/simulation/population"
	"math"
)

func decideNextAction(person *population.Person, sim *Simulation) {
	if(person.CurrentAction == population.ACTIONS_IDLE) {
		// benefits, cost := getFishingBenefit(person, sim)

		// Check for items
	}
}

func tickNeeds(person *population.Person, sim *Simulation) {
	tickSleep(person, sim)
	tickHunger(person, sim)
	tickThirst(person, sim)
}

func tickSleep(person *population.Person, sim *Simulation) {
	if(person.CurrentAction == population.ACTIONS_SLEEPING) {
		person.Needs[population.NEEDS_ENERGY] = math.Max(100, person.Needs[population.NEEDS_ENERGY] + SLEEPING_ENERGY_RESTORED_PER_TICK)

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