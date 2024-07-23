package scales

import (
	"fornof.me/m/v2/src/types"
)

func HandleMajor(roman string, chord *[]types.NBEFNote, pattern []int, chordInfo *[]types.NBEFNote, keyType string) {
	offset := 0
	romanOut := getRomanOnly(roman)
	// todo - handle if there are less than 3 in pattern
	switch romanOut {
	case "I":
		offset = 0
		upperRomanChord(chord, pattern, offset, chordInfo)
		if keyType == "minor_natural" {
			(*chord)[1].Halfsteps += 1 // 3 is bumped up
		} else {
			//no change for major
		}
	case "II":
		offset = 1
		upperRomanChord(chord, pattern, offset, chordInfo)
		(*chord)[1].Halfsteps -= 1 // 3 is bumped down

	case "III":
		offset = 2
		upperRomanChord(chord, pattern, offset, chordInfo)
		(*chord)[1].Halfsteps -= 1 // 3 is bumped down

	case "IV":
		offset = 3
		upperRomanChord(chord, pattern, offset, chordInfo)

	case "V":
		offset = 4
		upperRomanChord(chord, pattern, offset, chordInfo)

	case "VI":
		offset = 5
		upperRomanChord(chord, pattern, offset, chordInfo)
		(*chord)[1].Halfsteps -= 1 // 3 is bumped down

	case "VII":
		offset = 6
		upperRomanChord(chord, pattern, offset, chordInfo)

	case "i":
		offset = 0
		upperRomanChord(chord, pattern, offset, chordInfo)
		if keyType == "minor_natural" {
			// no change for minor
		} else {
			(*chord)[1].Halfsteps += 1 // 3 is bumped up for major?

		}

	case "ii":
		offset = 1
		upperRomanChord(chord, pattern, offset, chordInfo)

	case "iii":
		offset = 2
		upperRomanChord(chord, pattern, offset, chordInfo)
	case "iv":
		offset = 3
		upperRomanChord(chord, pattern, offset, chordInfo)
		(*chord)[1].Halfsteps -= 1 // 3 is bumped up

	case "v":
		offset = 4
		upperRomanChord(chord, pattern, offset, chordInfo)
		(*chord)[1].Halfsteps -= 1 // 3 is bumped down

	case "vi":
		offset = 5
		upperRomanChord(chord, pattern, offset, chordInfo)
	case "vii":
		offset = 6
		upperRomanChord(chord, pattern, offset, chordInfo)

		// no change if vii, diminished brings down the fifth?

	}
	//println(types.StringAllNotes(chord), "before modifiers")
	handleModifiers(roman, chord)
	//println(types.StringAllNotes(chord), "after modifiers")
}
