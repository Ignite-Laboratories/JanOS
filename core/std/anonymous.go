package std

import "fmt"

type Anonymous map[string]any

func (a Anonymous) ComponentLen() uint {
	return uint(len(a))
}

func (a Anonymous) Components(names ...string) map[string]any {
	if len(names) == 0 {
		return a
	}

	result := make(map[string]any)
	for _, name := range names {
		result[name] = a[name]
	}
	return result
}

func (a Anonymous) Component(named string) any {
	return a[named]
}

func (a Anonymous) Set(named string, value any) Vector {
	a[named] = value
	return a
}

func (a Anonymous) From(other Anonymous) Vector {
	for k, v := range other {
		a[k] = v
	}
	return a
}

func (a Anonymous) String() string {
	// TODO: Write this like the letter vectors output themselves, but without a name
	return fmt.Sprintf("%v", a)
}
