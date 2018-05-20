package simulation

import (
	"github.com/FabianK1991/civilization/simulation/world"
	"sync"
	"time"
	"encoding/json"
	"github.com/FabianK1991/civilization/simulation/population"
)

type Simulation struct {
	sync.Mutex

	started bool

	Time int64
	World world.World
	Population population.Population
}

func (sim *Simulation) Init() {
	sim.World.Init(sim.Population)
}

func (sim *Simulation) Start() {
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

func (sim *Simulation) Stop() {
	sim.Lock()
	defer sim.Unlock()

	sim.started = false
}

func (sim *Simulation) GetJSON() string {
	sim.Lock()
	defer sim.Unlock()

	data, err := json.Marshal(sim); if err != nil {
		return ""
	}

	return string(data)
}

func (sim *Simulation) tick() {
	sim.Lock()
	defer sim.Unlock()

	sim.Time++

	time.Sleep(TICK_TIMEOUT * time.Millisecond)
}
