package scales

// minor chord only drops 5th one
// major drops into a minor and a 5th ,
// if played out , middle 3rd does not drop
// nashville numbering
// dominant chords - 5th (delta), V7 , natually resolve to I
// 3m
// IM7 Im7
// Alberte Base,
// Broken chords- split triad into lowest and highest (C E G :=> C , EG)
// arpegios - seventh chords - 4 notes
// not clashing : don't move in parallel fifths
// inversions. IV/5 V/3 F/5 F/C same
// harmonize - find chord for 1,3,5, 7th ,6th, 8th that C plays in.
// sequencing - offset variations. keep the rhythm.
// inversions
// major key to shift over to a minor key.
// dynamics - pppp , to FFF , velocity,
// rhythm variation - slow slow quick quick.
// tempo changes .
// reflection,  - flipping it.
// instrumentation - take an existing idea strings play it Forte, horns pick it up, instrument combo. have it bounce around
//
// one going up, two notes changing. slightly more pleasing sound than all notes going up.
// penta tonic - 5ths stacked on eachother.  4th and 7th removed.
// minor pentatonic 2nd 6th removed
// pentatonic for melody only.
// polyrhythms - take two points in time, 1 rhythm plays 4 notes to a measure, another plays 3
// 3/2 , 4/5 per two measures, per 4 measures, etc.
// orchestra ensemble.
// musescore - 4 minutes , vocals,
import (
	"strings"

	"fornof.me/m/v2/src/types"
	"github.com/rs/zerolog/log"
)

var offsetMajorIndex = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII"}
var offsetMinorIndex = []string{"i", "ii", "iii", "iv", "v", "vi", "vii", "viii"}

func HandleRomanChord(chord types.Chord, romanOut string) []types.NBEFNoteRequest {
	offset := -999
	romanChange := types.Constants.Major
	majorScale := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO"}

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

	if romanChange == "up" {
		println(romanChange)
		if strings.Contains(majorScale[offset], "O") {
			// get out of that augmented jail
			if len(chordRequest.ChordNotes) > 2 {
				chordRequest.ChordNotes[2].Halfsteps += 1 // for modifier to act on this
			}
		}
		//it's a major like I, II, III, turn the  roman chord into a major
		if len(chord.Pattern) >= 2 {
			chordRequest.ChordNotes[1].Halfsteps += 1 //todo fix
		}
	}
	if romanChange == "down" {
		println(romanChange)
		if strings.Contains(majorScale[offset], "O") {
			// get out of that augmented jail
			chordRequest.ChordNotes[2].Halfsteps += 1 // for modifier to act on this
		}
		//it's a minor like i, ii, iii, turn the upper roman chord into a minor
		chordRequest.ChordNotes[1].Halfsteps -= 1

	}
	if romanChange == "none" {
		println(romanChange)

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
