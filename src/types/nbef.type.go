package types

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
	Note          int               `yaml:"note,omitempty"`
	Muted         bool              `yaml:"muted,omitempty"`
	Label         string            `yaml:"label,omitempty"`
	Duration      string            `yaml:"duration,omitempty"`
	IsIOAutomatic bool              `yaml:"is_io_automatic,omitempty"`
	Insert        bool              `yaml:"insert,omitempty"`
	Halfsteps     int               `yaml:"halfsteps,omitempty"`
	Entries       map[string]string `yaml:"entries,omitempty"`
}
