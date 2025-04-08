package scales

import (
	"strconv"
	"strings"

	"fornof.me/m/v2/src/types"
	"github.com/rs/zerolog/log"
)

func HandleBaseRoman(chord types.Chord) []types.NBEFNoteRequest {
	noteReturn := []types.NBEFNoteRequest{}

	switch chord.ChordType {
	case types.Constants.Major:
		noteReturn = append(noteReturn, HandleMajor(chord)...)
	// case types.Constants.MinorNatural:
	// 	scales.HandleMajor(roman, &noteReturn, pattern, chordInfo, types.Constants.MinorNatural)
	// case types.Constants.MinorHarmonic:
	// 	scales.HandleMinorHarmonic(roman, &noteReturn, pattern, chordInfo)
	// case types.Constants.MinorMelodic:
	// 	scales.HandleMinorMelodic(roman, &noteReturn, pattern, chordInfo)
	default:
		println("keyType not implemented:", chord.ChordType)
		noteReturn = append(noteReturn, chord.ChordInfo)
	}

	if chord.TimeSec != nil && !chord.IsSplit {
		moveTimeAhead := types.NBEFNoteRequest{TimeSec: chord.TimeSec[0], Track: -1}
		noteReturn = append(noteReturn, moveTimeAhead)
	}
	return noteReturn
}
func ParseStringToChordList(chordStr string) []types.Chord {
	romanLines := strings.Split(chordStr, "\n")

	chordList := []types.Chord{}

	for _, r := range romanLines {
		if r == "" {
			continue
		}

		currentChord := types.Chord{}
		currentChord.ChordInfo.Track = -1
		values := strings.Split(r, ",")
		for _, v := range values {

			if v == "" {
				continue
			}
			entry := strings.Split(v, ":")
			if len(entry) < 2 {
				continue

			}
			key := strings.Trim(entry[0], " \t\n,")
			value := strings.Trim(entry[1], " \t\n,")
			switch key {
			case "label":
				currentChord.ChordInfo.Label = value
			case "vol":
				vol, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					log.Error().Msgf("error parsing int for vol %s", value)
				}
				currentChord.ChordInfo.Velocity = int(vol)
			case "track":
				track, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					log.Error().Msgf("error parsing int for track %s", value)
				}
				currentChord.ChordInfo.Track = int(track)
			case "chord":
				currentChord.Chord = value
				log.Info().Msgf("inchord %s", value)
			case "note":
				currentChord.ChordInfo.Note = &value
			case "io":
				currentChord.ChordInfo.Signal = value
			case "midi":
				midi, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					log.Error().Msgf("error parsing int for midi %s", value)
				}
				currentChord.ChordInfo.Midi = int(midi)
			case "chord_type":
				currentChord.ChordType = value
			case "chord_pattern":
				pattern := strings.Split(value, "|")
				for _, p := range pattern {
					pInt, err := strconv.Atoi(p)
					if err != nil {
						println("error parsing int for pattern", p)
					}
					currentChord.Pattern = append(currentChord.Pattern, pInt)
				}
			case "split":
				isSplit, err := strconv.ParseBool(value)
				if err != nil {
					println("error parsing bool", value)
				}
				currentChord.IsSplit = isSplit
			case "time":
				pattern := strings.Split(value, "|")
				currentChord.TimeSec = append(currentChord.TimeSec, pattern...)

			case "dur":
				currentChord.ChordInfo.Duration = value
			case "tempo":
				tempo, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					log.Error().Msgf("error parsing int for tempo %s", value)
				}
				currentChord.ChordInfo.Tempo = int(tempo)
			case "key_type":
				currentChord.ChordInfo.KeyType = value
			case "key_note":
				currentChord.ChordInfo.KeyNote = value
			case "offset":
				offset, err := strconv.ParseInt(value, 10, 64)
				currentChord.Offset = int(offset)
				if err != nil {
					log.Error().Msgf("error parsing int for offset %s", value)
				}
			case "halfsteps":
				halfsteps, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					log.Error().Msgf("error parsing int for halfsteps %s", value)
				}
				currentChord.ChordInfo.Halfsteps = int(halfsteps)
			default:
				println("default, could not find key value", key, value)

			}

			// todo handle IsSplit with more options with pattern , for now true/false
			// todo soon - test this please

		}
		if currentChord.TimeSec == nil {
			currentChord.TimeSec = []string{}
		}
		if currentChord.Chord == "" || currentChord.ChordInfo.Note != nil || currentChord.IsSplit {
			if len(currentChord.TimeSec) != 0 {
				currentChord.ChordInfo.TimeSec = currentChord.TimeSec[len(currentChord.TimeSec)-1]
			}

		}

		if currentChord.ChordInfo.Duration == "" && currentChord.TimeSec != nil && len(currentChord.TimeSec) > 0 {
			//borrow from time
			currentChord.ChordInfo.Duration = GetFractionFromTime(currentChord.TimeSec[len(currentChord.TimeSec)-1])
		}
		if (currentChord.ChordInfo.TimeSec == "" || currentChord.ChordInfo.Track == -1) && currentChord.TimeSec == nil && len(currentChord.TimeSec) == 0 {
			continue
		}
		chordList = append(chordList, currentChord)
	}
	return chordList
}

func ParseChordList(chordList *[]types.Chord) []types.NBEFNoteRequest {

	outNotes := []types.NBEFNoteRequest{}

	for _, chord := range *chordList {
		// major, melodic minor, harmonic minor, natural minor. 7th chords
		// slash chord
		if chord.Chord == "" && (chord.ChordInfo.Note == nil || chord.ChordInfo.Midi != 0) {
			outNotes = append(outNotes, chord.ChordInfo)
			continue
		} else {
			notes := FindNotesForChord(chord)
			outNotes = append(outNotes, notes...)
		}

	}
	return outNotes
}

func FindNotesForChord(chord types.Chord) []types.NBEFNoteRequest {

	if chord.ChordType == "" {
		chord.ChordType = types.Constants.Major
	}
	if chord.Chord == "" {
		if chord.ChordInfo.TimeSec != "" {
			return []types.NBEFNoteRequest{chord.ChordInfo}
		}
		return nil
	}
	println("CHORD_START", chord.Chord, chord.ChordType)
	// split into just chord
	romanOut := getRomanOnly(chord.Chord)
	if romanOut == "" {
		// assume it's a letter, handle that.
		// this needs to probably be upstreams some more
		return HandleLetter(chord)
	}
	return HandleBaseRoman(chord)
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
func MakeChord(root_note int, offset int, pattern []int, halfsteps []int) []types.NBEFNoteRequest {
	// todo , have the option to have root note as a letter.
	if len(pattern) == 0 {
		pattern = []int{0, 2, 4}
	}
	if len(halfsteps) == 0 {
		halfsteps = []int{0}
	}
	count := len(pattern)

	chord := []types.NBEFNoteRequest{}
	for i := range count - 1 {
		nextNote := strconv.Itoa(root_note + offset + pattern[i])
		note := types.NBEFNoteRequest{Note: &nextNote, Halfsteps: halfsteps[i%len(halfsteps)]}
		if i == 0 {
			note.Label = "root_note"
		}
		chord = append(chord, note)

	}
	return chord
}
