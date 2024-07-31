package scales

// minor chord only drops 5th one
// major drops into a minor and a 5th ,
// if played out , middle 3rd does not drop
import (
	"fornof.me/m/v2/src/types"
	"github.com/rs/zerolog/log"
)

func HandleRomanChord(chord types.Chord, romanOut string) []types.NBEFNoteRequest {
	offset := -999
	romanChange := types.Constants.Major
	majorScale := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO"}
	offsetMajorIndex := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII"}
	offsetMinorIndex := []string{"i", "ii", "iii", "iv", "v", "vi", "vii", "viii"}
	for i, v := range majorScale {
		if v == chord.Chord { //includes modifiers such as O
			offset = i
			romanChange = "none"
			chord.Chord = getRomanOnly(chord.Chord) // if modifier is on the scale,
			//remove, it's fine just the way it is
			break
		}
	}
	if offset == -999 {
		for i, v := range offsetMajorIndex {
			if v == romanOut {
				offset = i
				romanChange = "up"
				break
			}
		}
	}
	if offset == -999 {
		for i, v := range offsetMinorIndex {
			if v == romanOut {
				offset = i
				romanChange = "down"
				break
			}
		}

	}
	if offset == -999 {
		log.Error().Msgf("offset not found, using 0 %s" + chord.Chord)
		offset = 0
	}
	chordRequest := types.ChordRequest{}
	chordRequest.Chord = chord
	chordRequest.RomanType = romanChange
	chordRequest.ChordNotes = upperRomanChord(chord, offset)

	if romanChange == "up" && len(chord.Pattern) >= 2 {
		//it's a minor like i, ii, iii, turn the upper roman chord into a minor
		chordRequest.ChordNotes[1].Halfsteps = 1
	}
	if romanChange == "down" && len(chord.Pattern) >= 2 {
		//it's a minor like i, ii, iii, turn the upper roman chord into a minor
		chordRequest.ChordNotes[1].Halfsteps = -1
	}
	if romanChange == "none" && len(chord.Pattern) >= 2 {
		//it's a minor like i, ii, iii, turn the upper roman chord into a minor
		chordRequest.ChordNotes[1].Halfsteps = 0
	}

	handleModifiers(&chordRequest)

	return chordRequest.ChordNotes
	//println(types.StringAllNotes(chord), "after modifiers")
}
func HandleMajor(chord types.Chord) []types.NBEFNoteRequest {

	romanOut := getRomanOnly(chord.Chord)
	if romanOut == "" {
		// assume it's a letter, handle that.
		return HandleLetter(chord)
	}
	return HandleRomanChord(chord, romanOut)

}
