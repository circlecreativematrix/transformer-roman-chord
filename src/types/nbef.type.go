package types

import (
	"fmt"
	"strings"
)

type ChordRequest struct {
	Chord      Chord
	RomanType  string
	ChordNotes []NBEFNoteRequest
}
type Chord struct {
	Chord     string // this is I, IV, i, #ivO, F/G, IV/G, F_1/3 (first inversion) base is 3rd note
	IsSplit   bool
	Offset    int             // offset on the number on the scale
	ChordType string          // major, minor_natural, minor_harmonic, minor_melodic
	Pattern   []int           // this is the pattern for the chord, to be read by NotesOut render
	ChordInfo NBEFNoteRequest // chord info for entire chord, to be read by NotesOut render
	TimeSec   []string        // this is time for entire chord, to be read by NotesOut render
}

type NBEFNoteRequest struct {
	Midi          int               `yaml:"midi,omitempty"`
	Signal        string            `yaml:"signal,omitempty"`
	TimeSec       string            `yaml:"time_s,omitempty"`
	OriginalTime  string            `yaml:"original_time,omitempty"`
	OriginalNote  string            `yaml:"original_note,omitempty"`
	Track         int               `yaml:"track" default:"-1"`
	Velocity      string            `yaml:"velocity"`
	BeatType      string            `yaml:"beat_type,omitempty"`
	NoteType      string            `yaml:"note_type,omitempty"`
	Tempo         int               `yaml:"tempo,omitempty"`
	KeyNote       string            `yaml:"key_note,omitempty"`
	KeyType       string            `yaml:"key_type,omitempty"`
	Note          *string           `yaml:"note,omitempty"`
	Muted         bool              `yaml:"muted,omitempty"`
	Label         string            `yaml:"label,omitempty"`
	Duration      string            `yaml:"duration,omitempty"`
	IsIOAutomatic bool              `yaml:"is_io_automatic,omitempty"`
	Insert        bool              `yaml:"insert,omitempty"`
	Halfsteps     int               `yaml:"halfsteps,omitempty"`
	Entries       map[string]string `yaml:"entries,omitempty"`
}

func (n NBEFNoteRequest) String(filter ...string) string {
	builder := strings.Builder{}
	println(n.Track, "track")
	if len(filter) != 0 {
		if n.Halfsteps != 0 {
			builder.WriteString(fmt.Sprintf("halfsteps:%d,", n.Halfsteps))
		}
		if n.Note != nil {
			builder.WriteString(fmt.Sprintf("note:%s,", *n.Note))
		}

		if n.TimeSec != "" {
			builder.WriteString(fmt.Sprintf("time:%s,", n.TimeSec))
		}
		if n.Duration != "" {
			builder.WriteString(fmt.Sprintf("dur:%s,", n.Duration))
		}
	} else {
		if n.Halfsteps != 0 {
			builder.WriteString(fmt.Sprintf("halfsteps:%d,", n.Halfsteps))
		}
		if n.Note != nil {
			builder.WriteString(fmt.Sprintf("note:%s,", *n.Note))
		}

		if n.TimeSec != "" {
			builder.WriteString(fmt.Sprintf("time:%s,", n.TimeSec))
		}
		if n.KeyNote != "" {
			builder.WriteString(fmt.Sprintf("key_note:%s,", n.KeyNote))

		}
		if n.KeyType != "" {
			builder.WriteString(fmt.Sprintf("key_type:%s,", n.KeyType))

		}
		if n.Tempo != 0 {
			builder.WriteString(fmt.Sprintf("tempo:%d,", n.Tempo))
		}
		if n.Duration != "" {
			builder.WriteString(fmt.Sprintf("dur:%s,", n.Duration))
		}
		if n.Label != "" {
			builder.WriteString(fmt.Sprintf("label:%s,", n.Label))
		}

	}
	if n.Midi != 0 {
		builder.WriteString(fmt.Sprintf("midi:%d,", n.Midi))
	}
	if n.Signal != "" {
		builder.WriteString(fmt.Sprintf("io:%s,", n.Signal))
	}
	if n.Velocity != "" {
		builder.WriteString(fmt.Sprintf("vol:%s,", n.Velocity))
	}
	if n.Track != -1 {
		builder.WriteString(fmt.Sprintf("track:%d,", n.Track))
	}
	if len(builder.String()) > 1 {
		return builder.String()[0 : builder.Len()-1]
	} else {
		return builder.String()
	}
}

func StringAllNotes(n *[]NBEFNoteRequest) string {
	result := []string{}

	for _, note := range *n {
		result = append(result, note.String())

	}
	return strings.Join(result, "\n")
}
