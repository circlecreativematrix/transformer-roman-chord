package scales

import (
	"regexp"
	"strings"

	"fornof.me/m/v2/src/services"
	"fornof.me/m/v2/src/types"
)

func getAtomicFromLetter(letterChord string) types.Atomic {
	// can I write diminished O or + augmented?
	letterRegex, _ := regexp.Compile("([A-Ga-g]+)([0-9]?)([#@]?)([Mm]?)")
	letter := letterRegex.FindAllStringSubmatch(letterChord, -1)
	atomic := types.Atomic{
		Letter:   letter[0][1],
		Octave:   letter[0][2],
		FullNote: letter[0][0],
	}

	if letter[0][3] == "@" {
		atomic.Halfstep = -1
	}
	if letter[0][3] == "#" {
		atomic.Halfstep = 1
	}
	if letter[0][4] == "m" {
		atomic.IsMinor = true
	} else {
		atomic.IsMinor = false
	}
	return atomic
}

func getMidiNoteFromAtomic(atomic types.Atomic) int {
	services.InitConvert()
	if atomic.Octave == "" {
		atomic.Octave = "4"
	}

	call := strings.ToLower(atomic.Letter) + atomic.Octave
	startNote := services.Letter_to_midi[call]
	startNote += atomic.Halfstep
	return startNote

}
func getNoteFromMidi(midi int) *string {
	services.InitConvert()
	strOut := strings.ToUpper(services.Midi_to_letter[midi])
	return &strOut
}
func getNotesFromAtomicAndChord(chord types.Chord, atomic types.Atomic) []types.NBEFNoteRequest {
	// get the start note
	chordNotes := []types.NBEFNoteRequest{}
	startNote := getMidiNoteFromAtomic(atomic)
	offsetHalfstepPattern := []int{0, 4, 7}
	if atomic.IsMinor {
		offsetHalfstepPattern = []int{0, 3, 7}
	}
	for i := 0; i < 3; i++ {
		note := chord.ChordInfo
		note.Note = getNoteFromMidi(startNote + offsetHalfstepPattern[i])
		chordNotes = append(chordNotes, note)
	}

	return chordNotes
}
func HandleLetter(chord types.Chord) []types.NBEFNoteRequest {
	atomic := getAtomicFromLetter(chord.Chord)
	return getNotesFromAtomicAndChord(chord, atomic)
}
