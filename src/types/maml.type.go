package types

type NicknameNote struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
type Nicknames struct {
	Notes []NicknameNote `yaml:"notes"`
}
type Header struct {
	SavePath string `yaml:"save_path"`
}
type Input struct {
	Notes string `yaml:"notes"`
}
type Phrase struct {
	Name       string            `yaml:"-"`
	Type       string            `yaml:"type"`
	InNotes    []string          `yaml:"-"`
	OutNotes   []NBEFNoteRequest `yaml:"-"`
	Input      Input             `yaml:"input"`
	OutputMidi string            `yaml:"output_midi"` //output_midi: "midi/stage_song.mid"
}

type Maml struct {
	Header    Header            `yaml:"header"`
	Nicknames Nicknames         `yaml:"nicknames"`
	Phrases   map[string]Phrase `yaml:"phrases"`
}
