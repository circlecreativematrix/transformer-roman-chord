// +build wasm

package main
import (
"syscall/js"
"fmt"
"strings"
"regexp"
)



func chordWasm(this js.Value, args []js.Value) interface{} {
	inputYamlStr := args[0].String()
	//regex remove quotes at the very beginning and end of string
	re := regexp.MustCompile(`^"(.*)"$`)
	inputYamlStr = re.ReplaceAllString(inputYamlStr, "$1")
	inputYamlStr = strings.ReplaceAll(inputYamlStr, `\n`, "\n")
	inputYamlStr = strings.ReplaceAll(inputYamlStr, `\"`, "\"")
	inputYamlStr = strings.ReplaceAll(inputYamlStr, `\'`, "'")
	fmt.Println("noties", inputYamlStr)
	standardIn := services.MarshalYamlStringToObj(inputYamlStr)
	nbefIn := services.ParseStandardObj(standardIn)
	//services.WriteNBEFToFile(args.OutputPath, nbefIn)
	nbefStr := services.MarshalNBEFToString(nbefIn)

	return nbefStr

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
