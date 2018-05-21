package simulation

import (
	"github.com/FabianK1991/civilization/simulation/population"
	"github.com/FabianK1991/civilization/simulation/world"
	"github.com/FabianK1991/civilization/utils"
)

type NeedChange struct {
	change float64
	useableLater bool
}

type Benefit struct {
	needsChanges map[population.NeedsType]NeedChange
	skillsChanges map[population.SkillType]float64
}

type Cost struct {
	ticks float64
}

func getCollectingWaterBenefit(person *population.Person, sim *Simulation) (*Benefit, *Cost) {
	x, y, found := world.FindClosestResource(person, world.RESOURCE_WATER, sim.World)

	if(!found) {
		return nil, nil
	}

	return &Benefit{
		needsChanges: map[population.NeedsType]NeedChange{
			population.NEEDS_THIRST: {
				change: population.GetWaterBottle().NeedsChange[population.NEEDS_THIRST] * 5,
				useableLater: true,
			},
		},
		skillsChanges: map[population.SkillType]float64{
			population.SKILLS_COLLECTING_WATER: 0.1,
		},
	}, &Cost {
		ticks: world.QuickDistanceCoordinates(person.X, person.Y, x, y) + utils.ControlledGrowth(1, 12, 0.3, person.Skills[population.SKILLS_COLLECTING_WATER] + 1, false),
	}
}

func getFishingBenefit(person *population.Person, sim *Simulation) (*Benefit, *Cost) {
	x, y, found := world.FindClosestResource(person, world.RESOURCE_FISH, sim.World)

	if(!found) {
		return nil, nil
	}

	return &Benefit{
		needsChanges: map[population.NeedsType]NeedChange{
			population.NEEDS_HUNGER: {
				change: population.GetFish().NeedsChange[population.NEEDS_HUNGER],
				useableLater: true,
			},
		},
		skillsChanges: map[population.SkillType]float64{
			population.SKILLS_FISHING: 0.1,
		},
	}, &Cost {
		ticks: world.QuickDistanceCoordinates(person.X, person.Y, x, y) + utils.ControlledGrowth(1, 12, 0.3, person.Skills[population.SKILLS_FISHING] + 1, false),
	}
}