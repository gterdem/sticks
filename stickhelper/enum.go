package stickhelper

// Alias hide the real type of the enum
// and users can use it to define the var for accepting enum
type Alias = int

type list struct {
	Lit        Alias
	Touched    Alias
	Diminished Alias
}

// Enum for public use
var Enum = &list{
	Lit:        0,
	Touched:    1,
	Diminished: 2,
}
