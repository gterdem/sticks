package stickhelper

// Alias hide the real type of the enum
// and users can use it to define the var for accepting enum
// Credit to: https://medium.com/@iNatata/use-struct-as-enum-in-go-6a314ae78678
type Alias = int

type list struct {
	Lit        Alias
	Touched    Alias
	Diminished Alias
	Growing    Alias
}

// Enum for public use
var Enum = &list{
	Lit:        0,
	Touched:    1,
	Diminished: 2,
	Growing:    3,
}
