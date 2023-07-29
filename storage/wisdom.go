package storage

type Wisdom map[string]struct{}

func (w Wisdom) Get() string {
	for one := range w {
		return one
	}
	return "Wisdom 0"
}

func NewWisdom() Wisdom {
	return map[string]struct{}{
		"Wisdom 1": {},
		"Wisdom 2": {},
		"Wisdom 3": {},
		"Wisdom 4": {},
	}
}
