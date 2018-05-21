package population

type Item struct {
	Id ItemId

	Uses int

	NeedsChange map[NeedsType]float64
	SkillsChange map[SkillType]float64
}



