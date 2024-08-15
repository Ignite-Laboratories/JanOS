package JanOS

type Symbol string

type Dimension struct {
	Value  float64
	Name   string
	Symbol Symbol
}

type dimensionManager struct {
	dimensions       map[string]*Dimension
	bufferDimensions map[string]*BufferDimension
}

func NewDimensionManager() *dimensionManager {
	return &dimensionManager{
		dimensions:       make(map[string]*Dimension),
		bufferDimensions: make(map[string]*BufferDimension),
	}
}

func (mgr *dimensionManager) GetName() string {
	return "Dimensions"
}

func (mgr *dimensionManager) GetDimension(name string) *Dimension {
	return mgr.dimensions[name]
}

func (mgr *dimensionManager) GetBufferDimension(name string) *BufferDimension {
	return mgr.bufferDimensions[name]
}

func (mgr *dimensionManager) NewDimension(name string, symbol Symbol, defaultValue float64) *Dimension {
	d := &Dimension{
		Name:   name,
		Symbol: symbol,
		Value:  defaultValue,
	}
	mgr.dimensions[name] = d
	Universe.Printf(mgr, "Let '%s' [%s] = %f", name, string(symbol), defaultValue)
	return d
}
