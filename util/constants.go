//nolint:exhaustive // constants
package util

// Direction is a direction in 2D space.
type Direction byte

// Directions in 2D space. Index `i` grows South, index `j` grows East.
const (
	N Direction = 1 << iota
	E
	S
	W
	NE           = N | E
	SE           = S | E
	NW           = N | W
	SW           = S | W
	C  Direction = 0
)

// Delta8 is a map of 8 directions to their index delta.
var Delta8 = map[Direction][2]int{
	N:  {-1, 0},
	NE: {-1, 1},
	E:  {0, 1},
	SE: {1, 1},
	S:  {1, 0},
	SW: {1, -1},
	W:  {0, -1},
	NW: {-1, -1},
}

var Delta9 = [9][2]int{
	Delta8[NW], Delta8[N], Delta8[NE],
	Delta8[W], {0, 0}, Delta8[E],
	Delta8[SW], Delta8[S], Delta8[SE],
}

// Delta4 is a map of 4 directions to their index delta.
var Delta4 = map[Direction][2]int{
	N: {-1, 0},
	E: {0, 1},
	S: {1, 0},
	W: {0, -1},
}

// Diagonal4 is a map of 4 diagonal directions to their index delta.
var Diagonal4 = map[Direction][2]int{
	NE: {-1, 1},
	SE: {1, 1},
	SW: {1, -1},
	NW: {-1, -1},
}

// Opposite maps a direction to its opposite direction.
var Opposite = map[Direction]Direction{
	N:  S,
	E:  W,
	S:  N,
	W:  E,
	NE: SW,
	SE: NW,
	NW: SE,
	SW: NE,
}

// ConvertFrom converts a byte of U/D/L/R to a Direction.
var ConvertFrom = map[byte]Direction{
	'U': N,
	'D': S,
	'L': W,
	'R': E,
	'^': N,
	'v': S,
	'<': W,
	'>': E,
}

// Right90 maps a direction to the direction after a 90-degree right turn.
var Right90 = map[Direction]Direction{
	N:  E,
	E:  S,
	S:  W,
	W:  N,
	NE: SE,
	SE: SW,
	SW: NW,
	NW: NE,
}
