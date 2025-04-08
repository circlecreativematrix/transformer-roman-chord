package scales

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"fornof.me/m/v2/src/services"
	"fornof.me/m/v2/src/types"
)

func getAtomicFromLetter(letterChord string) types.Atomic {
	// can I write diminished O or + augmented?
	letterRegex, _ := regexp.Compile("([A-Ga-g]+)([0-9]?)([#@]?)([Mm]?)")
	fullChord := strings.Split(letterChord, "/")
	if len(fullChord) > 1 {
		atomic := getAtomicFromLetter(fullChord[0])
		atomic.BaseNote = fullChord[1]
		return atomic
	}
	letter := letterRegex.FindAllStringSubmatch(letterChord, -1)
	if letter[0][2] == "" {
		letter[0][2] = "4"
	}
	octaveInt, err := strconv.Atoi(letter[0][2])
	if err != nil {
		log.Println(err)
	}

	atomic := types.Atomic{
		Letter:   letter[0][1],
		Octave:   octaveInt,
		FullNote: (letter[0][1] + letter[0][2] + letter[0][3]),
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
	if atomic.Octave < 0 {
		atomic.Octave = 4
	}

	call := strings.ToLower(atomic.Letter) + strconv.Itoa(atomic.Octave)
	startNote := services.Letter_to_midi[call]
	startNote += atomic.Halfstep
	return startNote

}
func getNoteFromMidi(midi int) *string {
	services.InitConvert()
	strOut := strings.ToUpper(services.Midi_to_letter[midi])
	return &strOut
}
func OffsetNoteThroughMidi(note *string, halfstep int) *string {
	atomicTemp := getAtomicFromLetter(*note)
	midi := getMidiNoteFromAtomic(atomicTemp)
	return getNoteFromMidi(midi + halfstep)
}
func HandleLetter(chord types.Chord) []types.NBEFNoteRequest {
	atomic := getAtomicFromLetter(chord.Chord)
	// get the start note
	chordNotes := []types.NBEFNoteRequest{}
	startNote := getMidiNoteFromAtomic(atomic)
	offsetHalfstepPattern := []int{0, 4, 7}
	offsetMidi := 0
	handledBaseNote := false
	if atomic.BaseNote != "" {

		if atomic.BaseNote == "3" {
			offsetMidi = 4
			offsetHalfstepPattern = []int{0 + offsetMidi, 3 + offsetMidi, 8 + offsetMidi}
			handledBaseNote = true
		}
		if atomic.BaseNote == "5" {
			offsetMidi = 7
			offsetHalfstepPattern = []int{0 + offsetMidi, 5 + offsetMidi, 9 + offsetMidi}
			handledBaseNote = true
		}
		match, err := regexp.MatchString("^[a-zA-Z]", atomic.BaseNote)
		if err != nil {
			log.Default().Println(err)
		}
		if match {
			//baseNoteMidi := getMidiNoteFromAtomic(getAtomicFromLetter(atomic.BaseNote))
			// offsetMidi = (baseNoteMidi - startNote)
			// offsetHalfstepPattern = []int{0, 4, 7}
			// startNote = baseNoteMidi
		}

	}

	if atomic.IsMinor {
		offsetHalfstepPattern = []int{0, 3, 7}
	}

	for i := 0; i < len(offsetHalfstepPattern); i++ {
		note := chord.ChordInfo

		note.Note = getNoteFromMidi(startNote + offsetHalfstepPattern[i])
		chordNotes = append(chordNotes, note)
	}
	if atomic.BaseNote != "" && !handledBaseNote {
		//first := chordNotes[0]
		note := chord.ChordInfo

		noter := getAtomicFromLetter(atomic.BaseNote)
		note.Note = &noter.FullNote
		first := chordNotes[0]
		firstAtomic := getAtomicFromLetter(*first.Note)
		if strings.Contains(*note.Note, firstAtomic.Letter) {
			// already in the chord and at the base, return the chord
			return chordNotes
		}
		third := chordNotes[2]
		thirdAtomic := getAtomicFromLetter(*third.Note)

		second := chordNotes[1]
		secondAtomic := getAtomicFromLetter(*second.Note)

		if strings.Contains(*note.Note, secondAtomic.Letter) {
			// shift first one up , make this one the base
			firstAtomic.Octave += 1
			octaveStr := firstAtomic.Letter + strconv.Itoa(firstAtomic.Octave)
			first.Note = &octaveStr
			chordNotes = []types.NBEFNoteRequest{note, third, first}
			return chordNotes
		}
		if strings.Contains(*note.Note, thirdAtomic.Letter) {
			// shift first one up and second one up , make this one the base
			firstAtomic.Octave += 1
			octaveStr := firstAtomic.Letter + strconv.Itoa(firstAtomic.Octave)
			first.Note = &octaveStr
			secondAtomic.Octave += 1
			secondStr := secondAtomic.Letter + strconv.Itoa(secondAtomic.Octave)
			second.Note = &secondStr
			chordNotes = []types.NBEFNoteRequest{note, first, second}
			return chordNotes
		}

		// check if base note is underneath already , raise octaves maybe?
		chordNotes = []types.NBEFNoteRequest{second, third}
	}

	//outputNotes is the arrangement of notes in the chord
	outputNotes := chordNotes
	if len(chord.Pattern) > 0 {
		outputNotes = []types.NBEFNoteRequest{}
		for i := 0; i < len(chord.Pattern); i++ {
			note := chordNotes[chord.Pattern[i]%(len(chordNotes))]
			note.TimeSec = chord.TimeSec[i%(len(chord.TimeSec))]
			time := GetFractionFromTime(note.TimeSec)
			if note.TimeSec == "P" {

			} else if time == "" {
				numerator := strings.Trim(note.TimeSec, "P+-")
				note.Duration = numerator + "/1"
			} else {
				note.Duration = time
			}

			outputNotes = append(outputNotes, note)
		}
	}
	return outputNotes
}
