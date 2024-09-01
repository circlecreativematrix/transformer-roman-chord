package scales

import (
	"testing"

	"fornof.me/m/v2/src/types"
	"github.com/stretchr/testify/assert"
)
// todo: 
/**
 * 1. C Cm C7 Cm7 Cmaj7 CmM7 C6 Cm6 C6/9 C5 C9 Cm9 Cmaj9 C11 Cm11 Cmaj11 C13 Cm13 Cmaj13 Cadd C7-5 C7+5 Csus Cdim Cdim7 Cm7b5 Caug Caug7
   2. with octave C4/5
**/

func TestLetterRegex(t *testing.T) {
	t.Run("testing letter regex", func(t *testing.T) {
		info1 := getAtomicFromLetter("C4#M")
		assert.Equal(t, info1.FullNote, "C4#")
		assert.Equal(t, info1.Halfstep, 1)
		assert.Equal(t, info1.IsMinor, false)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, 4)
	})
	t.Run("testing letter regex without octave", func(t *testing.T) {
		info1 := getAtomicFromLetter("C#M")
		assert.Equal(t, info1.FullNote, "C4#")
		assert.Equal(t, info1.Halfstep, 1)
		assert.Equal(t, info1.IsMinor, false)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, 4)
	})
	t.Run("testing letter regex without sharp, and a minor", func(t *testing.T) {
		info1 := getAtomicFromLetter("Cm")
		assert.Equal(t, info1.FullNote, "C4")
		assert.Equal(t, info1.Halfstep, 0)
		assert.Equal(t, info1.IsMinor, true)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, 4)
	})
	t.Run("testing plain letter regex", func(t *testing.T) {
		info1 := getAtomicFromLetter("C")
		assert.Equal(t, info1.FullNote, "C4")
		assert.Equal(t, info1.Halfstep, 0)
		assert.Equal(t, info1.IsMinor, false)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, 4)
	})
}

func TestLetterChords(t *testing.T) {
	t.Run("testing HandleLetter C/E", func(t *testing.T) {
		chord := types.Chord{Chord: "C5/E5", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "E5", *(result[0].Note))
		assert.Equal(t, "G5", *(result[1].Note))
		assert.Equal(t, "C6", *(result[2].Note))
	})
	t.Run("testing HandleLetter C/3", func(t *testing.T) {
		chord := types.Chord{Chord: "C/3", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "E4", *(result[0].Note))
		assert.Equal(t, "G4", *(result[1].Note))
		assert.Equal(t, "C5", *(result[2].Note))
	})

	t.Run("testing HandleLetter C/5", func(t *testing.T) {
		chord := types.Chord{Chord: "C/5", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "G4", *(result[0].Note))
		assert.Equal(t, "C5", *(result[1].Note))
		assert.Equal(t, "E5", *(result[2].Note))
	})
	t.Run("testing HandleLetter C4/G4", func(t *testing.T) {
		chord := types.Chord{Chord: "C4/G4", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "G4", *(result[0].Note))
		assert.Equal(t, "C5", *(result[1].Note))
		assert.Equal(t, "E5", *(result[2].Note))
	})

	t.Run("testing HandleLetter", func(t *testing.T) {
		chord := types.Chord{Chord: "C", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "C4", *(result[0].Note))
		assert.Equal(t, "E4", *(result[1].Note))
		assert.Equal(t, "G4", *(result[2].Note))
	})
	t.Run("testing HandleLetter C5", func(t *testing.T) {
		chord := types.Chord{Chord: "C5", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "C5", *(result[0].Note))
		assert.Equal(t, "E5", *(result[1].Note))
		assert.Equal(t, "G5", *(result[2].Note))
	})
	t.Run("testing HandleLetter D3M", func(t *testing.T) {
		chord := types.Chord{Chord: "D3M", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "D3", *(result[0].Note))
		assert.Equal(t, "F3#", *(result[1].Note))
		assert.Equal(t, "A3", *(result[2].Note))
	})

	t.Run("testing HandleLetter B3M", func(t *testing.T) {
		chord := types.Chord{Chord: "B3M", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "B3", *(result[0].Note))
		assert.Equal(t, "D4#", *(result[1].Note))
		assert.Equal(t, "F4#", *(result[2].Note))
	})

	t.Run("testing HandleLetter C4m", func(t *testing.T) {
		chord := types.Chord{Chord: "C4m", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "C4", *(result[0].Note))
		assert.Equal(t, "D4#", *(result[1].Note))
		assert.Equal(t, "G4", *(result[2].Note))
	})
	t.Run("testing HandleLetter C4m/D5m", func(t *testing.T) {
		chord := types.Chord{Chord: "C4m/D#", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		result := HandleLetter(chord)
		assert.Equal(t, "D4#", *(result[0].Note))
		assert.Equal(t, "G4", *(result[1].Note))
		assert.Equal(t, "C5", *(result[2].Note))
	})

}

func TestRomanChordsSlash(t *testing.T) {
	t.Run("testing HandleLetter V/5", func(t *testing.T) {
		chord := types.Chord{Chord: "I/5", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		chord.ChordType = types.Constants.Major
		chord.ChordInfo.KeyNote = "C"
		romanOut := getRomanOnly(chord.Chord)
		result := HandleRomanChord(chord, romanOut)
		assert.Equal(t, "4", *(result[0].Note))
		assert.Equal(t, "7", *(result[1].Note))
		assert.Equal(t, "9", *(result[2].Note))
	})
	t.Run("testing HandleLetter V/3", func(t *testing.T) {
		chord := types.Chord{Chord: "I/3", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		chord.ChordType = types.Constants.Major
		chord.ChordInfo.KeyNote = "C"
		romanOut := getRomanOnly(chord.Chord)
		result := HandleRomanChord(chord, romanOut)
		assert.Equal(t, "2", *(result[0].Note))
		assert.Equal(t, "4", *(result[1].Note))
		assert.Equal(t, "7", *(result[2].Note))
	})

	t.Run("testing HandleLetter V/4 throws error", func(t *testing.T) {
		// chord := types.Chord{Chord: "I/4", ChordInfo: types.NBEFNoteRequest{TimeSec: "P"}}
		// chord.ChordType = types.Constants.Major
		// chord.ChordInfo.KeyNote = "C"
		// romanOut := getRomanOnly(chord.Chord)
		// result := HandleRomanChord(chord, romanOut)
		// assert.Equal(t, "4", *(result[0].Note))
		// assert.Equal(t, "4", *(result[1].Note))
		// assert.Equal(t, "4", *(result[2].Note))
	})

}
