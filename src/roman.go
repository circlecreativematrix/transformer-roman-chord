package src

import (
	"fornof.me/m/v2/src/scales"
	"fornof.me/m/v2/src/types"
)

func HandleBaseRoman(roman string, pattern []int, keyType string, chordInfo *[]types.NBEFNote) []types.NBEFNote {
	chord := []types.NBEFNote{}

	switch keyType {
	case types.Constants.Major:
		scales.HandleMajor(roman, &chord, pattern, chordInfo, types.Constants.Major)
	case types.Constants.MinorNatural:
		scales.HandleMajor(roman, &chord, pattern, chordInfo, types.Constants.MinorNatural)
	case types.Constants.MinorHarmonic:
		scales.HandleMinorHarmonic(roman, &chord, pattern, chordInfo)
	case types.Constants.MinorMelodic:
		scales.HandleMinorMelodic(roman, &chord, pattern, chordInfo)
	default:
		println("keyType not implemented:", keyType)
	}

	return chord
}

func FindNotesForChord(roman string, keyType string, chordInfo *[]types.NBEFNote) []types.NBEFNote {

	pattern := []int{0, 2, 4}
	chord := HandleBaseRoman(roman, pattern, keyType, chordInfo)
	//handleModifiers(roman, &chord)
	//handleNumbers(roman, &chord)
	//handleSharpsFlats(roman, &chord)
	return chord
}
func FindOffset(roman string, offsetNumeralArray []string) int {
	offset := 0
	for i, numeral := range offsetNumeralArray {
		if roman == numeral {
			offset = i
		}
	}
	return offset
}
func MakeChord(root_note int, offset int, pattern []int, halfsteps []int) []types.NBEFNote {
	// todo , have the option to have root note as a letter.
	if len(pattern) == 0 {
		pattern = []int{0, 2, 4}
	}
	if len(halfsteps) == 0 {
		halfsteps = []int{0}
	}
	count := len(pattern)

	chord := []types.NBEFNote{}
	for i := range count - 1 {
		nextNote := root_note + offset + pattern[i]
		note := types.NBEFNote{Note: &nextNote, Halfsteps: halfsteps[i%len(halfsteps)]}
		if i == 0 {
			note.Label = "root_note"
		}
		chord = append(chord, note)

	}
	return chord
}
