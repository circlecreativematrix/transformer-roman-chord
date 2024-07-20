package src

import (
	"regexp"
	"strings"

	"fornof.me/m/v2/src/types"
)

func handleModifiers(roman string, chord *[]types.NBEFNote) {
	regModifiers, _ := regexp.Compile("([-+Oo]?)")
	modifiersOut := regModifiers.FindString(roman)
	if modifiersOut == "" {
		return
	}
	slash := strings.Index(modifiersOut, "/")
	if slash != -1 {
		number := modifiersOut[slash+1:]
		println("number", number)
	}
	plus := strings.Index(modifiersOut, "+")
	if plus != -1 {
		(*chord)[2].Halfsteps += 1
		println("plus", plus)
	}
	if modifiersOut[0] == 'O' {
		println(modifiersOut[0], "diminished")
		(*chord)[1].Halfsteps -= 1
		(*chord)[2].Halfsteps -= 1

	}
	if modifiersOut[0] == 'o' {
		println(modifiersOut[0], "half diminished? what does this look like")
	}

}
func HandleBaseRoman(roman string, pattern []int) []types.NBEFNote {
	chord := []types.NBEFNote{}
	offset := 0
	regRomanOnly, _ := regexp.Compile("([ivIVXCxc]+)")
	romanOut := regRomanOnly.FindString(roman)
	switch romanOut {
	case "I":
		for i := 0; i < len(pattern); i++ {
			chord = append(chord, types.NBEFNote{Note: pattern[i] + offset})
		}

	case "II":
		offset = 1
		for i := 0; i < len(pattern); i++ {
			chord = append(chord, types.NBEFNote{Note: pattern[i] + offset})
		}
	case "III":
		offset = 2
		for i := 0; i < len(pattern); i++ {
			chord = append(chord, types.NBEFNote{Note: pattern[i] + offset})
		}
	case "i":
		offset = 0
		for i := 0; i < len(pattern); i++ {
			chord = append(chord, types.NBEFNote{Note: pattern[i] + offset})
			if i == 1 {
				chord[i].Halfsteps = -1
			}
		}
	case "ii":
		offset = 1
		for i := 0; i < len(pattern); i++ {
			chord = append(chord, types.NBEFNote{Note: pattern[i] + offset})
			if i == 1 {
				chord[i].Halfsteps = -1
			}
		}
	case "iii":
		offset = 2
		for i := 0; i < len(pattern); i++ {
			chord = append(chord, types.NBEFNote{Note: pattern[i] + offset})
			if i == 1 {
				chord[i].Halfsteps = -1
			}
		}
	}
	return chord
}
func handleNumbers(roman string, chord *[]types.NBEFNote) {
	regNumbers := regexp.MustCompile("/?[_0-9]+")
	numbersOut := regNumbers.FindString(roman)
	if numbersOut == "" {
		return
	}
	splitNumeratorDenominator := strings.Split(numbersOut, "_")
	numerator := splitNumeratorDenominator[0]
	denominator := ""
	if len(splitNumeratorDenominator) > 1 {
		denominator = splitNumeratorDenominator[1]
	}
	if numerator != "" {
		println("numerator", numerator)
	}
	if denominator != "" {
		println("denominator", denominator)
	}
	println("numerator", numerator, "denominator", denominator)

}
func handleSharpsFlats(roman string, chord *[]types.NBEFNote) {
	regSharpFlats, _ := regexp.Compile("([#@]?)")
	sharpFlatsOut := regSharpFlats.FindString(roman)
	if sharpFlatsOut == "" {
		return
	}
	if sharpFlatsOut == "#" {
		(*chord)[0].Halfsteps += 1
	}
	if sharpFlatsOut == "@" {
		(*chord)[0].Halfsteps -= 1
	}
}
func FindNotesForChord(roman string, keyType string, pattern []int) []types.NBEFNote {
	if pattern == nil || len(pattern) == 0 {
		pattern = []int{0, 2, 4}
	}
	chord := HandleBaseRoman(roman, pattern)
	handleModifiers(roman, &chord)
	handleNumbers(roman, &chord)
	handleSharpsFlats(roman, &chord)

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
		note := types.NBEFNote{Note: nextNote, Halfsteps: halfsteps[i%len(halfsteps)]}
		if i == 0 {
			note.Label = "root_note"
		}
		chord = append(chord, note)

	}
	return chord
}

//func ConvertNotationToChord(roman string, keyType string, root_note int, pattern []int) []types.NBEFNote {
//https://musictheory.pugetsound.edu/mt21c/DiatonicChordsInMajor.html
//chord := []types.NBEFNote{}

// counting := ["I","II", "III", "IV", "V", "VI" ,"VII", "VIII", "IX", "X"]
// countingAlso := ["i", "ii", "iii", "iv", "v", "vi", "vii", "viii", "ix", "x"]
//M-m-m-M-M-m
//regexSpecials := "[*/+]"

// todo: enter in translations for :
//https://musictheory.pugetsound.edu/mt21c/DiatonicChordsInMinor.html
// find the major or minor offset.
// majorRomanOffset := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO"}
// minorNaturalRomanOffset := []string{"i", "ii0", "III", "iv", "v", "VI", "VII"}
// minorHarmonicRomanOffset := []string{"i", "ii0", "III+", "iv", "V", "VI", "viiO"}
// minorMelodicRomanOffset := []string{"i", "ii", "III+", "IV", "V", "#viO", "viiO"}
// switch keyType {
// case "major":
// 	halfsteps := []int{0, 0, 0}
// 	offset := FindOffset(roman, majorRomanOffset)
// 	chord = MakeChord(root_note, offset, pattern, halfsteps)
// case "minor_natural":
// 	halfsteps := []int{0, 0, 0}
// 	offset := FindOffset(roman, minorNaturalRomanOffset)
// 	chord = MakeChord(root_note, offset, pattern, halfsteps)
// case "minor_harmonic":
// 	halfsteps := []int{0, 0, 0}
// 	offset := FindOffset(roman, minorHarmonicRomanOffset)
// 	chord = MakeChord(root_note, offset, pattern, halfsteps)
// }
// return chord
//}
