package types

type Atomic struct {
	Letter   string
	Octave   int
	Halfstep int
	IsMinor  bool // if m is present, it is minor, otherwise it is Major
	FullNote string
	BaseNote string
}
