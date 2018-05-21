package population

type ItemId int

const (
	ITEM_FISH = iota
	ITEM_WATER_BOTTLE
)

func GetWaterBottle() *Item {
	return &Item {
		Id: ITEM_WATER_BOTTLE,

		Uses: 5,
		NeedsChange:  map[NeedsType]float64{
			NEEDS_THIRST: 20,
		},
	}
}

func GetFish() *Item {
	return &Item {
		Id: ITEM_FISH,

		Uses: 1,
		NeedsChange:  map[NeedsType]float64{
			NEEDS_HUNGER: 50,
		},
	}
}
