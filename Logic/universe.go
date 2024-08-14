package Logic

type Universe struct {
	Dimensions  []*Dimension
	Environment *World
}

func NewUniverse(environment *World, dimensions ...*Dimension) *Universe {
	return &Universe{
		Dimensions:  dimensions,
		Environment: environment,
	}
}

func (u *Universe) BigBang() {

}
