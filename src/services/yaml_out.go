package services

import (
	"fornof.me/m/v2/src/types"
)

func NotesToString(notes *[]types.NBEFNoteRequest) string {
	result := ""
	for _, note := range *notes {
		if note.KeyNote != "" || note.KeyType != "" || note.Tempo != 0 || note.Label != "" {
			result += note.String()
		} else {
			result += note.String("SELECT")
		}
		result += "\n"
	}
	return result
}
func AddToPhrases(maml *types.Maml, phrase *types.Phrase) {
	if maml == nil {
		maml = &types.Maml{}
	}
	if maml.Phrases == nil {
		maml.Phrases = map[string]types.Phrase{}
	}
	if phrase.Type == "" {
		phrase.Type = "fornof.standard"
	}
	phrase.Input.Notes = NotesToString(&phrase.OutNotes)
	maml.Phrases[phrase.Name] = *phrase
}
func GenerateMaml(savePath string) types.Maml {
	maml := types.Maml{}
	maml.Header.SavePath = savePath
	maml.Nicknames.Notes = []types.NicknameNote{}
	// 	maml.Nicknames.Notes = append(maml.Nicknames.Notes,
	// 		NicknameNote{Name: "sample",
	// 			Value: `note:4,time:P+1/4\n
	// tempo:120,time:0,key_type:major,key_note:C4`})

	return maml
}
