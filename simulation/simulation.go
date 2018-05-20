package simulation

import (
	"github.com/FabianK1991/civilization/simulation/world"
	"sync"
	"time"
)

type Simulation struct {
	sync.Mutex

	started bool

	Time int64
	World world.World
	// population population.Population
}

func (sim *Simulation) init() {
	sim.started = false
}

func (sim *Simulation) start() {
	sim.Lock()

	if(sim.started) {
		sim.Unlock()
		return
	}

	sim.started = true
	sim.Unlock()

	for sim.started {
		sim.tick()
	}
}

func (sim *Simulation) stop() {
	sim.Lock()
	defer sim.Unlock()

	sim.started = false
}

func (sim *Simulation) tick() {
	sim.Lock()
	defer sim.Unlock()

	time.Sleep(TICK_TIMEOUT * time.Millisecond)
}
