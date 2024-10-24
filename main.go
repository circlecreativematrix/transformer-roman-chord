//go:build wasm
// +build wasm

package main

import (
	"fmt"
	"regexp"
	"strings"
	"syscall/js"

	"fornof.me/m/v2/src/scales"
	"fornof.me/m/v2/src/services"
	"gopkg.in/yaml.v3"
)

func chordWasm(this js.Value, args []js.Value) interface{} {
	inputYamlStr := args[0].String()
	//regex remove quotes at the very beginning and end of string
	re := regexp.MustCompile(`^"(.*)"$`)
	inputYamlStr = re.ReplaceAllString(inputYamlStr, "$1")
	inputYamlStr = strings.ReplaceAll(inputYamlStr, `\n`, "\n")
	inputYamlStr = strings.ReplaceAll(inputYamlStr, `\"`, "\"")
	inputYamlStr = strings.ReplaceAll(inputYamlStr, `\'`, "'")
	fmt.Println("noties", inputYamlStr) // todo: make this toml

	standardIn := services.MarshalYamlStringToObj(inputYamlStr)
	chordList := scales.ParseStringToChordList(standardIn.Notes)
	outString, err := yaml.Marshal(chordList)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("outString", string(outString))
	return outString
	// todo: make into TOML version of yaml - output each section to different track etc

	// phrases := types.PhraseList{}
	// phrases = append(phrases, types.Phrase{
	// 	Type:       "fornof.standard",
	// 	Name:       "chorus",
	// 	InNotes:    singleChorusChemistry(),
	// 	OutputMidi: "chorus.mid"})
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
	//MakeTheSong(phrases)
	//sleep 5 seconds

}

//export add
func add(x string, y string) string {
	return x + y
}
func main() {
	fmt.Println("welcome to the chords of creative matrix circle.")

	js.Global().Set("chordWasm", js.FuncOf(chordWasm))
	c := make(chan bool)
	// this is needed to keep the program from stopping

	<-c
	//sleep(60)
}
