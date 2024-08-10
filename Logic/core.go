package Logic

// Drivable represents a system that can be driven at an interval defined by the hosting architecture.
type Drivable interface {
	Tick()
}
