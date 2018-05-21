package simulation

import (
	"github.com/FabianK1991/civilization/simulation/world"
	"sync"
	"encoding/json"
	"time"
	"github.com/FabianK1991/civilization/simulation/population"
)

type Simulation struct {
	sync.Mutex

	started bool

	Time int
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

func (sim *Simulation) UpdateTile(updatedTile *world.Tile) {
	sim.Lock()
	defer sim.Unlock()

	sim.World[updatedTile.X][updatedTile.Y].TileType = updatedTile.TileType
	sim.World[updatedTile.X][updatedTile.Y].IsBlocked = updatedTile.IsBlocked
	sim.World[updatedTile.X][updatedTile.Y].Resources = updatedTile.Resources
}

func (sim *Simulation) tick() {
	time.Sleep(TICK_TIMEOUT * time.Millisecond)

	// Lock the simulation
	sim.Lock()
	defer sim.Unlock()

	sim.Time++

	for _, person := range sim.Population {
		tickNeeds(person, sim)
		decideNextAction(person, sim)
	}
}

func (sim *Simulation) GetDayTime() int {
	return sim.Time % TICKS_PER_DAY
}