package scales

// root postion 0 2 4
// first inversion 2 4 7
// second inversion 4 7 9
// 7th chord inversion - root position inversion 4 notes, 0 2 4 6
// 7th first inversion 2,4,6,7
// 7th , second inversion , 4,6,7,9
// 7th third inversion, 6,7,9,11
// GM/d Gm/D G+ , GO/D major minor , all notes of triad, then base note
//Gm_5/D
// chromatic notes
// c e g to 0 2 4 underneath.
import (
	"regexp"
	"strconv"
	"strings"

	"fornof.me/m/v2/src/types"
	"github.com/rs/zerolog/log"
)

func GetFractionFromTime(time string) string {
	r, _ := regexp.Compile(`(\d+\/\d+)`)
	fraction := r.FindString(time)
	if fraction != "" {
		return fraction
	} else {
		log.Error().Msg("no fraction for duration")
		return ""
	}
}
func getRomanOnly(roman string) string {
	regRomanOnly, _ := regexp.Compile("([ivIV]+)")
	return regRomanOnly.FindString(roman)
}
func HandleDiminished(chordRequest *types.ChordRequest) {
	regModifiers, _ := regexp.Compile("([-+Oo]+)")
	modifiersOut := regModifiers.FindString(chordRequest.Chord.Chord)
	println("modifiers out", modifiersOut, "roman", chordRequest.Chord.Chord)
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
		(chordRequest.ChordNotes)[2].Halfsteps += 1
		println("plus", plus)
	}
	if strings.Contains(modifiersOut, "O") {
		println("diminished")
		// does major/minor make a difference here?
		// should just be the top note
		//(*chord)[1].Halfsteps -= 1
		// ignore figured base ? maybe
		// if chordRequest.RomanType == types.Constants.Major {
		// 	// get to a minor first if a major
		// 	chordRequest.ChordNotes[1].Halfsteps -= 1
		// }
		if chordRequest.RomanType == "up" {
			chordRequest.ChordNotes[2].Halfsteps += 1
		} else {
			chordRequest.ChordNotes[2].Halfsteps -= 1
		}

	}
	if strings.Contains(modifiersOut, "o") {
		log.Error().Msgf("%b,%s", modifiersOut[0], "half diminished? psh , skipping like a 7th")
	}
}
func handleModifiers(chordRequest *types.ChordRequest) {
	handleSharpsFlats(chordRequest)
	handleNumbers(chordRequest)
	HandleDiminished(chordRequest)

}
func handleNumbers(chordRequest *types.ChordRequest) {

	regNumbers := regexp.MustCompile("/?[_0-9]+")
	numbersOut := regNumbers.FindString(chordRequest.Chord.Chord)
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
		println("numerator-ignore", numerator)
	}
	if denominator != "" {
		println("denominator-ignore", denominator)
	}
	println("numerator", numerator, "denominator", denominator)

}
func handleSharpsFlats(chordRequest *types.ChordRequest) {
	regSharpFlats, _ := regexp.Compile("([#@]+)")
	sharpFlatsOut := regSharpFlats.FindString(chordRequest.Chord.Chord)
	if sharpFlatsOut == "" {
		return
	}
	if sharpFlatsOut == "#" {
		println("found sharp")
		(chordRequest.ChordNotes)[0].Halfsteps += 1
	}
	if sharpFlatsOut == "@" {
		println("found flat")
		(chordRequest.ChordNotes)[0].Halfsteps -= 1
	}
}
func upperRomanChord(chord types.Chord, offset int) []types.NBEFNoteRequest {
	outNotes := []types.NBEFNoteRequest{}

	if chord.ChordInfo.TimeSec == "" {
		chord.ChordInfo.TimeSec = "P"
	}
	pattern := []int{0, 2, 4}
	for i := 0; i < len(pattern); i++ {

		info := chord.ChordInfo // copies last of the chord infos
		offPattern := strconv.Itoa(pattern[i] + offset + chord.Offset)
		info.Note = &offPattern
		outNotes = append(outNotes, info)
	}

	return outNotes

}
