package songs

import (
	"fmt"
	"os"
	"strconv"
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

var outputChords = []string{}
var key_note = []string{"F5", "F3"}

func addToOutput(chords []string, output *[]string) {
	if output == nil {
		outputChords = append(outputChords, chords...)
	} else {
		*output = append(*output, chords...)
	}
}

type Beat struct {
	Beat     string
	Count    int
	Duration string
}

func notes(notes []string, beats []Beat, appendEach []string) []string {
	return notesToStringArray(notes, beats, appendEach, false)
}
func chords(notes []string, beats []Beat, appendEach []string) []string {
	return notesToStringArray(notes, beats, appendEach, true)
}
func notesToStringArray(notes []string, beats []Beat, appendEach []string, isChords bool) []string {
	outNotes := []string{}
	if len(appendEach) == 0 {
		appendEach = []string{" "}
	}
	count := beats[0].Count
	selected := 0

	for i, inNotes := range notes {
		splitNotes := strings.Split(inNotes, " ")
		builder := []string{}
		for _, note := range splitNotes {
			if note != "" {
				if isChords {
					builder = append(builder, fmt.Sprintf("chord:%s", note))
				} else {
					builder = append(builder, fmt.Sprintf("note:%s", note))
				}
			}
			if beats[selected].Duration != "" {
				builder = append(builder, fmt.Sprintf("dur:%s", beats[selected].Duration))
			}
			if beats[selected].Beat != "" {
				builder = append(builder, fmt.Sprintf("time:%s", beats[selected].Beat))
			}
			if appendEach[i%len(appendEach)] != "" {
				builder = append(builder, appendEach[i%len(appendEach)])
			}

			outNotes = append(outNotes, strings.Join(builder, ","))
			count--
			if count <= 0 {
				selected = (selected + 1) % len(beats) // go to next one
				count = beats[selected].Count          // reset count to next one
			}
		}
	}
	return outNotes
}
func addToStringArray(array *[]string, amount int) []string {
	output := []string{}
	for _, str := range *array {
		splitArr := strings.Split(str, " ")

		for _, splitStr := range splitArr {
			intStr, err := strconv.ParseInt(splitStr, 10, 64)
			if err != nil {
				println(err)
			}
			output = append(output, strconv.Itoa(int(intStr)+amount))
		}

	}
	return output
}

func alberti(in []string, offset int, output *[]string) {
	newNotes := []string{in[0], in[1], in[2], in[1], in[2], in[1], in[2]}
	if offset != 0 {
		newNotes = addToStringArray(&newNotes, offset)
	}
	addToOutput(notes(newNotes, []Beat{
		{Beat: "P+1/16", Count: 6},
		{Beat: "P+1/8", Count: 1}}, []string{""}), output)

}

// func figure1(in []string, offset int) {
// 	newNotes := []string{in[0], in[1], in[1], "", in[2], in[1], in[1], ""}
// 	if offset != 0 {
// 		newNotes = addToStringArray(&newNotes, offset)
// 	}
// 	addToOutput(notes(newNotes, []Beat{
// 		{Beat: "P+1/4", Count: 1, Duration: "1/4"},
// 		{Beat: "P+1/8", Count: 2, Duration: "1/16"},
// 		{Beat: "P+1/8", Count: 1, Duration: "1/18"}}, []string{""}))

// }
func bounce(in []string, ending string, offset int, output *[]string) {
	if offset != 0 {
		in = addToStringArray(&in, offset)
		endingArr := []string{ending}
		endingArr = addToStringArray(&endingArr, offset)
		ending = strings.Join(endingArr, " ")
	}
	newNotes := []string{fmt.Sprintf("%s %s|%s %s|%s %s %s|%s %s|%s %s",
		in[0], in[1], in[2], in[1], in[2],
		in[0], in[1], in[2], in[1], in[2], ending)}
	addToOutput(notes(newNotes, []Beat{
		{Beat: "P+1/8", Count: 1},
		{Beat: "P+1/16", Count: 2},
		{Beat: "P+1/8", Count: 1},
		{Beat: "P+1/16", Count: 2},
		{Beat: "P+1/16", Count: 4},
		{Beat: "P+1/4", Count: 1},
	}, []string{"vol:127"}), output)
}
func melody(in []string, offset int, output *[]string) {
	if offset != 0 {
		in = addToStringArray(&in, offset)
	}
	// 2 1 3 2 0
	// 0 1 2 3
	newNotes := []string{fmt.Sprintf("%s %s %s %s %s",
		in[2], in[1], in[3], in[2], in[0]), ""}
	addToOutput([]string{"time:0"}, output)
	addToOutput(notes(newNotes, []Beat{
		{Beat: "P+1/4", Count: 1},
		{Beat: "P+1/8", Count: 1},
		{Beat: "P+1/4", Count: 2},
		{Beat: "P+1/4", Count: 1},
		{Beat: "P+1/8", Count: 1},
	}, []string{"vol:127"}), output)

}
func bridge(in []string, offset int, output *[]string) {
	if offset != 0 {
		in = addToStringArray(&in, offset)
	}
	// 2 1 3 2 0
	// 0 1 2 3
	newNotes := []string{fmt.Sprintf("%s %s %s %s %s %s",
		in[2], in[3], in[0], in[3], in[2], in[3])}
	addToOutput(notes(newNotes, []Beat{
		{Beat: "P+1/8", Count: 2},
		{Beat: "P+1/4", Count: 1},
		{Beat: "P+1/8", Count: 2},
		{Beat: "P+1/8", Count: 1},
	}, []string{"vol:127"}), output)

}
func bridgeLower(in []string, offset int, multiplier int, output *[]string) {
	if offset != 0 {
		in = addToStringArray(&in, offset)
	}
	if multiplier != 0 {
		in = multiplyToStringArray(&in, multiplier)
	}
	// 2 1 3 2 0
	// 0 1 2 3
	newNotes := in
	addToOutput(chords(newNotes, []Beat{
		{Beat: "P+1/4", Count: 1},
	}, []string{"vol:100"}), output)

}
func chorus(in []string, offset int, output *[]string) {
	if offset != 0 {
		in = addToStringArray(&in, offset)
	}
	// 2 1 3 2 0
	// 0 1 2 3
	newNotes := []string{fmt.Sprintf("%s %s %s %s %s",
		in[2], in[4], in[5], in[6], in[7])}
	addToOutput(notes(newNotes, []Beat{
		{Beat: "P+1/4", Count: 2, Duration: "3/8"},
		{Beat: "P+1/8", Count: 2},
		{Beat: "P+1/4", Count: 1},
	}, []string{"vol:127"}), output)

}
func chorusLower(in []string, offset int, output *[]string) {
	if offset != 0 {
		in = addToStringArray(&in, offset)
	}
	// 2 1 3 2 0
	// 0 1 2 3
	newNotes := in
	addToOutput(chords(newNotes, []Beat{
		{Beat: "P+1/2", Count: 2},
		{Beat: "P+1/2", Count: 1},
		{Beat: "P+1/2", Count: 1},
	}, []string{"vol:100"}), output)

}
func multiplyToStringArray(in *[]string, multiplier int) []string {
	outMultiplied := []string{}
	for i := 0; i < len(*in); i++ {
		for j := 0; j < multiplier; j++ {
			outMultiplied = append(outMultiplied, (*in)[i])
		}

	}
	return outMultiplied
}
func multiplyToStringArrayRest(in *[]string, multiplier int) []string {
	outMultiplied := []string{}
	for i := 0; i < len(*in); i++ {
		for j := 0; j < multiplier; j++ {
			outMultiplied = append(outMultiplied, (*in)[i])
		}
		outMultiplied = append(outMultiplied, "")

	}
	return outMultiplied
}
func melodyLower(in []string, offset int, multiplier int, output *[]string) {
	if multiplier != 0 {
		in = multiplyToStringArrayRest(&in, 3)
	}
	// 2 1 3 2 0
	// 0 1 2 3
	newNotes := in
	addToOutput(chords(newNotes, []Beat{
		{Beat: "P+1/4", Count: 1},
		{Beat: "P+1/8", Count: 2},
		{Beat: "P+1/4", Count: 1},
	}, []string{"vol:100"}), output)

}
func songLower() {

	addToOutput(header("bridge2-lower", "F3", "major", "100"), nil)
	bridgeLower([]string{"iii", "V", "IV", "iii"}, 0, 0, nil)
	bridgeLower([]string{"iii", "VI", "V", "IV"}, 0, 0, nil)
	bridgeLower([]string{"iii", "VI", "V", "IV"}, 0, 0, nil)
	addToOutput(header("chorus2-lower", "F3", "major", "60"), nil)
	chorusLower([]string{"ii", "IV", "VI", "IV"}, 0, nil)
	chorusLower([]string{"IV", "I", "V", "IV"}, 0, nil)
}

func singleBridge() []string {
	outString := []string{}
	addToOutput([]string{"time:0,dur:0,track:0"}, &outString)
	addToOutput(header("bridge", key_note[0], "major", "60"), &outString)

	bridge([]string{"0", "1", "2", "3"}, 0, &outString)
	bridge([]string{"0", "1", "2", "3"}, 0, &outString)
	addToOutput([]string{"chord:I,time:P+1/4"}, &outString)
	addToOutput([]string{"time:P+1/4"}, &outString)
	addToOutput([]string{"time:0,dur:0,track:1"}, &outString)
	addToOutput(header("bridge-lower", key_note[1], "major", "100"), &outString)
	// addToOutput([]string{"time:P+1/8,dur:1/4,note:22"})//bat signal
	bridgeLower([]string{"iii", "V", "IV", "iii"}, 0, 0, &outString)
	bridgeLower([]string{"iii", "VI", "V", "IV"}, 0, 0, &outString)
	return outString
}
func header(label string, keyNote string, keyType string, tempo string) []string {
	return []string{fmt.Sprintf("label:%s,key_note:%s,key_type:%s,tempo:%s", label, keyNote, keyType, tempo)}
}

type BeatEntries struct {
	Downbeat        string
	Offbeat         string
	BeatsPerMeasure int
	CountBeat       int
}



func singleVerseOne() []string {
	outString := []string{}
	addToOutput([]string{"tempo:60,time:0,dur:0,track:0,label:top-chorus"}, &outString)
	addToOutput(header("chorus", key_note[0], "major", "60"), &outString)
	chorus([]string{"0", "1", "2", "", "4", "5", "6", "7"}, 0, &outString)
	chorus([]string{"0", "1", "2", "", "4", "5", "6", "7"}, 0, &outString)
	chorus([]string{"0", "1", "2", "", "4", "5", "6", "7"}, 0, &outString)
	chorus([]string{"0", "1", "2", "", "4", "5", "6", "7"}, 0, &outString)
	addToOutput([]string{"time:0,dur:0,track:1,label:bottom-chorus"}, &outString)
	addToOutput(header("chorus-lower", key_note[1], "major", "60"), &outString)
	chorusLower([]string{"ii", "IV", "VI", "IV"}, 0, &outString)
	chorusLower([]string{"IV", "I", "V", "IV"}, 0, &outString)
	return outString
}