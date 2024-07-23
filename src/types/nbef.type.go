package types

import (
	"fmt"
	"strings"
)

type NBEFNote struct {
	Midi          int               `yaml:"midi,omitempty"`
	Signal        string            `yaml:"signal,omitempty"`
	TimeSec       string            `yaml:"time_s,omitempty"`
	OriginalTime  string            `yaml:"original_time,omitempty"`
	OriginalNote  string            `yaml:"original_note,omitempty"`
	Track         int               `yaml:"track"`
	Velocity      int               `yaml:"velocity"`
	BeatType      string            `yaml:"beat_type,omitempty"`
	NoteType      string            `yaml:"note_type,omitempty"`
	Tempo         int               `yaml:"tempo,omitempty"`
	KeyNote       string            `yaml:"key_note,omitempty"`
	KeyType       string            `yaml:"key_type,omitempty"`
	Note          *int              `yaml:"note,omitempty"`
	Muted         bool              `yaml:"muted,omitempty"`
	Label         string            `yaml:"label,omitempty"`
	Duration      string            `yaml:"duration,omitempty"`
	IsIOAutomatic bool              `yaml:"is_io_automatic,omitempty"`
	Insert        bool              `yaml:"insert,omitempty"`
	Halfsteps     int               `yaml:"halfsteps,omitempty"`
	Entries       map[string]string `yaml:"entries,omitempty"`
}

func (n NBEFNote) String() string {

	if n.Note != nil {
		return fmt.Sprintf("halfsteps:%d,time:%s,note:%d", n.Halfsteps, n.TimeSec, *n.Note)
	} else {
		return fmt.Sprintf("halfsteps:%d,time:%s", n.Halfsteps, n.TimeSec)
	}
}

func StringAllNotes(n *[]NBEFNote) string {
	result := []string{}
	for _, note := range *n {
		result = append(result, note.String())
	}
	return strings.Join(result, "\n")
}
