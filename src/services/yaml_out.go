package services

import (
	"fornof.me/m/v2/src/types"
	"gopkg.in/yaml.v3"
)

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
	Type       string `yaml:"type"`
	Input      Input  `yaml:"input"`
	OutputMidi string `yaml:"output_midi"` //output_midi: "midi/stage_song.mid"
}

type Maml struct {
	Header    Header            `yaml:"header"`
	Nicknames Nicknames         `yaml:"nicknames"`
	Phrases   map[string]Phrase `yaml:"phrases"`
}

func notesToString(notes *[]types.NBEFNote) string {
	result := ""
	for _, note := range *notes {
		result += note.String() + "\n"
	}
	return result
}
func GenerateYaml(notes *[]types.NBEFNote) Maml {
	maml := Maml{}
	maml.Header.SavePath = "C:/projects/music-user-reform/save"
	maml.Nicknames.Notes = []NicknameNote{}
	maml.Nicknames.Notes = append(maml.Nicknames.Notes,
		NicknameNote{Name: "header",
			Value: "tempo:120,time:0,key_type:major,key_note:C4"})
	maml.Phrases = map[string]Phrase{}
	phrase := Phrase{}
	phrase.Type = "fornof.standard"
	phrase.OutputMidi = "midi/goosbumps.mid"
	phrase.Input.Notes = `$header` + "\n" + notesToString(notes)
	maml.Phrases["phrase.standard.1"] = phrase
	return maml
}

func StringNotesYaml(notes *[]types.NBEFNote) string {
	maml := GenerateYaml(notes)
	result, _ := yaml.Marshal(maml)
	println(result)
	return string(result)
}
