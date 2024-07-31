package scales

/*
*
i,iiO, III+,iv,V,VI,viiO
*
*/
// func HandleMinorHarmonic(roman string, chord *[]types.NBEFNoteRequest, pattern []int, chordInfo *[]types.NBEFNoteRequest) {
// 	offset := 0
// 	romanOut := getRomanOnly(roman)
// 	// todo - handle if there are less than 3 in pattern
// 	switch romanOut {
// 	case "I":
// 		offset = 0
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped up
// 	case "II":
// 		offset = 1
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped up

// 	case "III":
// 		offset = 2
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change for minor_harmonic

// 	case "IV":
// 		offset = 3
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped up

// 	case "V":
// 		offset = 4
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change for harmonic
// 	case "VI":
// 		offset = 5
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change for harmonic
// 	case "VII":
// 		offset = 6
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1

// 	case "i":
// 		offset = 0
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// should be no change as long as minor gets propagated to header

// 	case "ii":
// 		offset = 1
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change for minor harmonic

// 	case "iii":
// 		offset = 2
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps -= 1 // 3 is bumped down

// 	case "iv":
// 		offset = 3
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change
// 	case "v":
// 		offset = 4
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped down

// 	case "vi":
// 		offset = 5
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped down)

// 	case "vii":
// 		offset = 6
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change if vii, diminished brings down the fifth?

// 	}
// 	//println(types.StringAllNotes(chord), "before modifiers")
// 	handleModifiers(roman, chord)

// 	//println(types.StringAllNotes(chord), "after modifiers")
// }
