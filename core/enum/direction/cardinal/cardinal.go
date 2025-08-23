// Package cardinal provides access to the cardinal.Direction enumeration system.
package cardinal

// Direction represents Primary cardinal direction. If you wish for intercardinal or secondary intercardinal directions,
// please see the Inter and InterSecondary types.  All intercardinal directions through the secondary set are defined here,
// such as NorthByNorthWest and SouthEast.  You may use the following type constraints -
//
// Direction ← N W E S, NW SW SE NE, NNE NNW WNW WSW SSW SSE ESE ENE
// Intercardinal ← N W E S, NW SW SE NE
// Primary ← N W E S
// Latitudinal ← N S
// Longitudinal ← E W
//
// Underneath, East and West are of the Longitudinal type, while North and South are of the Latitudinal type.  Both of these
// are aliases for the Primary type listed above.
//
// All dimensions can be distilled down to an infinitely repeating number line which can be traversed in binary directions -
// but, as you layer these dimensions on top of each other, they orthographically align relative to one another.  The terminology
// used to describe this is entirely dependent upon context, and as such I've provided a robust set of general abstract dimensions
// from which to describe this mechanic in code.  It truly does NOT matter which you use, as long as the called method knows
// how to talk in THAT language. =)
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, you walk "latitudinally" between rows along the Y axis and "longitudinally" between columns along the X axis.  Against a voxel,
// you'd walk negatively "in" or positively "out" along the Z axis.
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
type Direction interface {
	Primary | Inter | InterSecondary
}

// Intercardinal represents any primary cardinal direction, or any intercardinal direction that bisects two cardinal directions.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
type Intercardinal interface {
	Primary | Inter
}

// Primary represents any cardinal direction.
//
// NOTE: For complex intercardinal dimensions, see the Inter and InterSecondary types.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
type Primary byte

// Latitudinal represents the cardinal directions North and South.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
type Latitudinal = Primary

// Longitudinal represents the cadinal directions East and West.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
type Longitudinal = Primary

// InterSecondary represents a secondary intercardinal direction which bisects two intercardinal directions.
// A and B represent the two directions to bisect in creating this secondary intercardinal direction.
//
// I.E. NorthByNorthWest means you go halfway between North and NorthWest.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
type InterSecondary struct {
	A Primary
	B Inter
}

// Inter represents an intercardinal direction which bisects two cardinal directions.
// A and B represent the two directions to bisect in creating this intercardinal direction.
//
// I.E. NorthWest means you go halfway between North and West.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
type Inter struct {
	A Primary
	B Primary
}

// North represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var North Latitudinal

// NorthByNorthEast represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var NorthByNorthEast = InterSecondary{North, Inter{North, East}}

// NorthEast represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var NorthEast = Inter{North, East}

// EastByNorthEast represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var EastByNorthEast = InterSecondary{East, Inter{North, East}}

// East represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var East Longitudinal

// EastBySouthEast represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var EastBySouthEast = InterSecondary{East, Inter{South, East}}

// SouthEast represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var SouthEast = Inter{South, East}

// SouthBySouthEast represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var SouthBySouthEast = InterSecondary{South, Inter{South, East}}

// South represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var South Latitudinal

// SouthBySouthWest represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var SouthBySouthWest = InterSecondary{South, Inter{South, West}}

// SouthWest represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var SouthWest = Inter{South, West}

// WestBySouthWest represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var WestBySouthWest = InterSecondary{West, Inter{South, West}}

// West represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var West Longitudinal

// WestByNorthWest represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var WestByNorthWest = InterSecondary{West, Inter{North, West}}

// NorthWest represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var NorthWest = Inter{North, West}

// NorthByNorthWest represents a cardinal.Primary.
//
// See direction.Any, Direction, Primary, Inter, InterSecondary, Intercardinal, Latitudinal, Longitudinal, North, West, East, and South
var NorthByNorthWest = InterSecondary{North, Inter{North, West}}
