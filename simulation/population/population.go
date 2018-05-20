package population

type Population []*Person
type Gender int

const (
	GENDER_UNKNOWN Gender = iota
	GENDER_MALE
	GENDER_FEMALE
)

type Needs struct {
	Energy float64
	Hunger float64
	Thirst float64
}

type Person struct {
	X int
	Y int

	Gender Gender
	Prename string
	Surname string

	Needs Needs
}


