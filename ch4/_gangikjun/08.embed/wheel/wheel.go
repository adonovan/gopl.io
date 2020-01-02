package wheel

type point struct {
	X, Y int
}

type Circle struct {
	point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
