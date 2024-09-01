package songs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"fornof.me/m/v2/src/scales"
	"fornof.me/m/v2/src/services"
	"fornof.me/m/v2/src/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func MakeTheSong(phrases types.PhraseList) {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// phrases := types.PhraseList{}
	// 	phrases = append(phrases, types.Phrase{
	// 		Type:       "fornof.standard",
	// 		Name:       "chorus",
	// 		InNotes:    singleChorus(),
	// 		OutputMidi: "chorus.mid"})
	// phrases = append(phrases, types.Phrase{
	// 	Type:       "fornof.standard",
	// 	Name:       "verseOneFrenchPoodle",
	// 	InNotes:    singleVerseOne(),
	// 	OutputMidi: "verse1.mid"})

	// phrases = append(phrases, types.Phrase{
	// 	Type:       "fornof.standard",
	// 	Name:       "bridgeFrenchPoodle",
	// 	InNotes:    singleBridge(),
	// 	OutputMidi: "bridge1FrenchPoodle.mid"})
	maml := services.GenerateMaml("/mnt/c/projects/music-user-reform/savemidi/")
	for _, phrase := range phrases {
		chordList := scales.ParseStringToChordList(strings.Join(phrase.InNotes, "\n"))
		phrase.OutNotes = scales.ParseChordList(&chordList)
		services.AddToPhrases(&maml, &phrase)
		// out to file

	}
	resultMamlYamlString, err := yaml.Marshal(maml)
	if err != nil {
		log.Error().Msgf("error_MakeTheSong: %v", err)
	}
	println(string(resultMamlYamlString))
	// out to file
	path := "/mnt/c/projects/music-user-reform/converter-standard-note"
	name := "maml_test.yml"
	err = os.WriteFile(path+"/"+name, []byte(resultMamlYamlString), 0644)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Println("wrote to file", path+"/"+name)
	//sleep 5 seconds
	time.Sleep(2 * time.Second)
}
