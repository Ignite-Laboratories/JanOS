package JanOS

type Symbol string

type Dimension struct {
	Value  float64
	Name   string
	Symbol Symbol
}

type dimensionManager struct {
	dimensions map[string]*Dimension
}

func NewDimensionManager() *dimensionManager {
	return &dimensionManager{dimensions: make(map[string]*Dimension)}
}

func (mgr *dimensionManager) GetName() string {
	return "Dimensions"
}

func (mgr *dimensionManager) GetDimension(name string) *Dimension {
	return mgr.dimensions[name]
}

func (mgr *dimensionManager) NewDimension(name string, symbol Symbol, defaultValue float64) *Dimension {
	d := &Dimension{
		Name:   name,
		Symbol: symbol,
		Value:  defaultValue,
	}
	mgr.dimensions[name] = d
	Universe.Printf(mgr, "'%s' [%s] set to: %f", name, string(symbol), defaultValue)
	return d
}
