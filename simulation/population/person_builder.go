package population

import "github.com/twinj/uuid"

func CreatePerson(x int, y int, prename string, surname string, gender Gender) *Person {
	return &Person{
		UUID: uuid.NewV4().String(),
		X: x,
		Y: y,

		Gender: gender,
		Prename: prename,
		Surname: surname,

		Needs: map[NeedsType]float64{
			NEEDS_ENERGY: 100,
			NEEDS_HUNGER: 100,
			NEEDS_THIRST: 100,
		},

		Skills: map[SkillType]float64{
			SKILLS_FISHING: 0,
			SKILLS_COLLECTING_WATER: 0,
		},
	}
}