package scales

import (
	"regexp"
	"strings"

	"fornof.me/m/v2/src/types"
)

func getRomanOnly(roman string) string {
	regRomanOnly, _ := regexp.Compile("([ivIVXCxc]+)")
	return regRomanOnly.FindString(roman)
}
func handleModifiers(roman string, chord *[]types.NBEFNote) {
	handleSharpsFlats(roman, chord)
	handleNumbers(roman, chord)
	regModifiers, _ := regexp.Compile("([-+Oo/]+)")
	modifiersOut := regModifiers.FindString(roman)
	println("modifiers out", modifiersOut, "roman", roman)
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
	if strings.Contains(modifiersOut, "O") {
		println("diminished")
		// does major/minor make a difference here?
		// should just be the top note
		//(*chord)[1].Halfsteps -= 1
		(*chord)[2].Halfsteps -= 1

	}
	if strings.Contains(modifiersOut, "o") {
		println(modifiersOut[0], "half diminished? what does this look like")
	}

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
func upperRomanChord(chord *[]types.NBEFNote, pattern []int, offset int, chordInfo *[]types.NBEFNote) {
	if len(*chordInfo) == 0 {
		(*chordInfo) = append((*chordInfo), types.NBEFNote{Note: nil})
	}
	for i := 0; i < len(pattern); i++ {
		println("note", pattern[i]+offset)
		info := (*chordInfo)[len(*chordInfo)-1]
		offPattern := pattern[i] + offset
		println("offPattern", offPattern)
		info.Note = &offPattern
		(*chord) = append((*chord), info)
	}

}
