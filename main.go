package main

import (
	"fmt"
	"os"

	"fornof.me/m/v2/src/scales"
	"fornof.me/m/v2/src/services"
)

func mainCLI(path string) string {
	standardIn := services.ReadStandardYamlFile(path)
	chordList := scales.ParseStringToChordList(standardIn.Notes)
	outNotes := scales.ParseChordList(&chordList)
	outStandard := services.NotesToString(&outNotes)
	//outString, err := yaml.Marshal(outNotes)

	fmt.Println("outString", outStandard)
	return string(outStandard)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] != "" {
		fmt.Println("running cli")
		path := os.Args[1]
		fmt.Println(mainCLI(path))
		return

	}

}
