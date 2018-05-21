package simulation

import "time"

const TICK_TIMEOUT time.Duration = 1000
const TICKS_PER_DAY int = 144
const TICKS_PER_HOUR int = TICKS_PER_DAY / 24

const SLEEPING_ENERGY_RESTORED_PER_TICK float64 = 2.1

