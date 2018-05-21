package population

type NeedsType int
type SkillType int
type Gender int
type PeopleActions int

const (
	SKILLS_FISHING = iota
	SKILLS_COLLECTING_WATER
)

const (
	NEEDS_ENERGY NeedsType = iota
	NEEDS_HUNGER
	NEEDS_THIRST
)

const (
	GENDER_UNKNOWN Gender = iota
	GENDER_MALE
	GENDER_FEMALE
)

const (
	ACTIONS_IDLE PeopleActions = iota
	ACTIONS_SLEEPING
	ACTIONS_WALKING
	ACTIONS_FISHING
)

type Person struct {
	UUID string

	X int
	Y int

	Gender Gender
	Prename string
	Surname string

	Health int

	Needs map[NeedsType]float64
	Skills map[SkillType]float64

	CurrentAction PeopleActions
	ActionDuration int

	Inventory []*Item
}

