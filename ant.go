package main

type Orientation int

const (
	L Orientation = iota
	U
	R
	D
)

// What does an ant know? What can it do?
// It knows where it is? And can modify it's surroundings
type Ant struct {
	grid *Grid
	X    int
	Y    int
	// Where is the ant facing
	orientation Orientation
	rules       map[Color]Move
}

type RotationType int

const (
	LeftRotation RotationType = iota - 1
	NullRotation
	RightRotation
)

// Now each ant has certain rules it lives by.
// A rule is nothing but a mapping from a color to a move
type Move struct {
	// Rotations are L: 0 R: 1
	Rotation RotationType
	Recolor  Color
}

// Initialize an ant with the default rule for a langton ant
func NewAnt(g *Grid, xinit int, yinit int) *Ant {
	a := Ant{
		grid:        g,
		X:           xinit,
		Y:           yinit,
		orientation: U,
		rules: map[Color]Move{
			COLOR_WHITE: {
				Rotation: LeftRotation,
				Recolor:  COLOR_RED,
			},
			COLOR_RED: {
				Rotation: RightRotation,
				Recolor:  COLOR_WHITE,
			},
		},
	}
	return &a
}

func NewAntWithRules(g *Grid, xinit int, yinit int, rules map[Color]Move) *Ant {
	a := NewAnt(g, xinit, yinit)
	a.rules = rules
	return a
}

func (a *Ant) Move() {

	cx, cy, o, clr := a.X, a.Y, a.orientation, a.grid.GetColor(a.X, a.Y)
	// log.Printf("X: %d, Y: %d, O: %d, C: %v", a.X, a.Y, o, clr)
	mv := a.rules[clr]
	// Change color of current cell
	a.grid.SetColor(cx, cy, mv.Recolor)
	// Move to next cell
	a.orientation = Orientation((int(o) + int(mv.Rotation) + 4) % 4)
	switch a.orientation {
	case L:
		a.X -= 1
	case U:
		a.Y -= 1
	case R:
		a.X += 1
	case D:
		a.Y += 1
	}
	a.X = (a.X + a.grid.C) % a.grid.C
	a.Y = (a.Y + a.grid.R) % a.grid.R
}
