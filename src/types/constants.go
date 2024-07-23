package types

type Constant struct {
	Major         string
	MinorNatural  string
	MinorHarmonic string
	MinorMelodic  string
}

var Constants = Constant{
	Major:         "major",
	MinorNatural:  "minor_natural",
	MinorHarmonic: "minor_harmonic",
	MinorMelodic:  `minor_melodic`,
}
