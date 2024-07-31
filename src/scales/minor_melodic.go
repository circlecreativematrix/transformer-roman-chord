package scales

/*
*
i, minor, (1, b3, 5)
ii, minor, (2, 4, 6)
III, augmented, (b3, 5, 7)
IV, Major, (4, 6, 8)
V, Major, (5, 7, 2)
vi, diminished, (6, 8, b3)
vii, diminished, (7, 2, 4)
*
*/
// func HandleMinorMelodic(roman string, chord *[]types.NBEFNoteRequest, pattern []int, chordInfo *[]types.NBEFNoteRequest) {
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
// 		// no change for minor_natural

// 	case "IV":
// 		offset = 3
// 		upperRomanChord(chord, pattern, offset, chordInfo)

// 	case "V":
// 		offset = 4
// 		upperRomanChord(chord, pattern, offset, chordInfo)

// 	case "VI":
// 		offset = 5
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped up

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
// 		// no change for minor natural
// 	case "iii":
// 		offset = 2
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped up
// 	case "iv":
// 		offset = 3
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped up

// 	case "v":
// 		offset = 4
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		(*chord)[1].Halfsteps += 1 // 3 is bumped down

// 	case "vi":
// 		offset = 5
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change for vi, diminished brings down the fifth?

// 	case "vii":
// 		offset = 6
// 		upperRomanChord(chord, pattern, offset, chordInfo)
// 		// no change if vii, diminished brings down the fifth?

// 	}
// 	//println(types.StringAllNotes(chord), "before modifiers")
// 	handleModifiers(roman, chord)
// 	//println(types.StringAllNotes(chord), "after modifiers")
// }
