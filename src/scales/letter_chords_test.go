package scales

import (
	"testing"

	"fornof.me/m/v2/src/types"
	"github.com/stretchr/testify/assert"
)

func TestLetterRegex(t *testing.T) {
	t.Run("testing letter regex", func(t *testing.T) {
		info1 := getAtomicFromLetter("C4#M")
		assert.Equal(t, info1.FullNote, "C4#M")
		assert.Equal(t, info1.Halfstep, 1)
		assert.Equal(t, info1.IsMinor, false)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, "4")
	})
	t.Run("testing letter regex without octave", func(t *testing.T) {
		info1 := getAtomicFromLetter("C#M")
		assert.Equal(t, info1.FullNote, "C#M")
		assert.Equal(t, info1.Halfstep, 1)
		assert.Equal(t, info1.IsMinor, false)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, "")
	})
	t.Run("testing letter regex without sharp, and a minor", func(t *testing.T) {
		info1 := getAtomicFromLetter("Cm")
		assert.Equal(t, info1.FullNote, "Cm")
		assert.Equal(t, info1.Halfstep, 0)
		assert.Equal(t, info1.IsMinor, true)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, "")
	})
	t.Run("testing plain letter regex", func(t *testing.T) {
		info1 := getAtomicFromLetter("C")
		assert.Equal(t, info1.FullNote, "C")
		assert.Equal(t, info1.Halfstep, 0)
		assert.Equal(t, info1.IsMinor, false)
		assert.Equal(t, info1.Letter, "C")
		assert.Equal(t, info1.Octave, "")
	})
}

func TestLetterChords(t *testing.T) {
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
		assert.Equal(t, "C4", *(result[0].Note))
		assert.Equal(t, "D4#", *(result[1].Note))
		assert.Equal(t, "G4", *(result[2].Note))
		assert.Equal(t, "D4#", *(result[3].Note))
	})
	
}
