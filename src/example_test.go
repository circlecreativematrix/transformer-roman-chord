package src

// dominant
import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"fornof.me/m/v2/src/scales"
	"fornof.me/m/v2/src/services"
	"github.com/stretchr/testify/assert"
)

func TestChordListToNotes(t *testing.T) {
	t.Run("testing chord list to notes", func(t *testing.T) {
		//chord_pattern:0|2|4|8
		roman := `
		tempo:120,key_note:C4,key_type:major
		chord:I,split:0,chord_type:major,time:P+1/4
		chord:ii
		chord:iii 
		chord:IV 
		chord:V
		chord:vi 
		chord:viiO`
		//roman := []string
		chordList := scales.ParseStringToChordList(roman)
		outNotes := scales.ParseChordList(&chordList)
		assert.Equal(t, *outNotes[1].Note, "0")
		assert.Equal(t, outNotes[0].KeyType, "major")
	})
}

func TestIsSplit(t *testing.T) {
	t.Run("testing chord isSplit, should put time after each chord note", func(t *testing.T) {
		//chord_pattern:0|2|4|8
		roman := `
		tempo:120,key_note:C4,key_type:major
		chord:I,split:1,time:P+1/4`
		//roman := []string
		chordList := scales.ParseStringToChordList(roman)
		outNotes := scales.ParseChordList(&chordList)
		assert.Equal(t, outNotes[1].TimeSec, "P+1/4")
		assert.Equal(t, outNotes[2].TimeSec, "P+1/4")
		assert.Equal(t, outNotes[2].TimeSec, "P+1/4")
		assert.Equal(t, "0", *outNotes[1].Note)
		assert.Equal(t, "2", *outNotes[2].Note)
		assert.Equal(t, "4", *outNotes[3].Note)

	})
}

func TestLetterChords(t *testing.T) {
	// t.Run("testing letter chords", func(t *testing.T) {
	// 	roman := `
	// 		    chord:C,split:0,chord_type:major,key_type:major,key_note:C4,time:P+1/4
	// 			chord:C4#M,time:P+1/4
	// 			chord:E,time:P+1/4
	// 			chord:FM,time:P+1/4
	// 			chord:Gm,time:P+1/4
	// 			chord:A,time:P+1/4
	// 			chord:B,time:P+1/4
	// 			chord:C,offset:7,time:P+1/4`

	// 	chordList := scales.ParseStringToChordList(roman)
	// 	outNotes := scales.ParseChordList(&chordList)
	// 	yamlOutMaml := services.StringNotesYaml(&outNotes)
	// 	t.Log(string(yamlOutMaml))
	// })
}
func rest(time string) string {
	return fmt.Sprintf("time:%s", time)
}
func funkyTown() string {
	chordsOfNote := []string{"I", "V", "vi", "IV", ""}
	beatPattern := []string{"1/4", "1/4", "1/8", "1/8", "1/4"}
	outputChords := []string{}
	outputChords = append(outputChords, "key_type:major,key_note:D4,tempo:60,time:P")
	for i, chord := range chordsOfNote {
		outputChords = append(outputChords, fmt.Sprintf("chord:%s,time:P+%s,dur:%s", chord, beatPattern[i], beatPattern[i]))
	}

	for i := range chordsOfNote {
		outputChords = append(outputChords, fmt.Sprintf("chord:%s,time:P+%s,dur:%s", chordsOfNote[i*2%len(chordsOfNote)], beatPattern[i], beatPattern[i]))
	}
	for i, chord := range chordsOfNote {
		outputChords = append(outputChords, fmt.Sprintf("chord:%s,time:P+%s,dur:%s", chord, beatPattern[i], beatPattern[i]))
	}
	outputChords = append(outputChords, rest("P+1/8"))
	for i := range chordsOfNote {
		outputChords = append(outputChords, fmt.Sprintf("chord:%s,time:P+%s,dur:%s", chordsOfNote[i*2%len(chordsOfNote)], beatPattern[i], beatPattern[i]))
	}
	outputChords = append(outputChords, rest("P+1/8"))
	return strings.Join(outputChords, "\n")
}
func outMajor() string {
	return `
			    chord:I,dur:1/4,split:0,chord_type:major,key_type:major,key_note:C4,time:P+1/4
				chord:ii,time:P+1/4
				chord:iii,time:P+1/4
				chord:IV,time:P+1/4
				chord:V,time:P+1/4
				chord:vi,time:P+1/4
				chord:VII,time:P+1/4
				chord:I,offset:7,time:P+1/4`
}

func TestYamlMamlOutput(t *testing.T) {
	t.Run("testing output yaml", func(t *testing.T) {

		//roman := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO", "I"}
		//roman := []string{"i", "iiO", "III", "iv", "v", "VI", "VII"}
		outputChords := funkyTown() //outMajor()
		//chordsOfNote := []string{"I", "V", "vi", "IV"}

		chordList := scales.ParseStringToChordList(outputChords)
		outNotes := scales.ParseChordList(&chordList)
		yamlOutMaml := services.NotesToString(&outNotes)
		t.Log(string(yamlOutMaml))

		// out to file
		path := "/mnt/c/projects/music-user-reform/converter-standard-note"
		name := "maml_test.yml"
		err := os.WriteFile(path+"/"+name, []byte(yamlOutMaml), 0644)
		if err != nil {
			t.Error(err)
		}
		t.Log("wrote to file", path+"/"+name)
		//sleep 5 seconds
		time.Sleep(2 * time.Second)

	})
}

// func TestFindNotesForChord(t *testing.T) {
// 	t.Run("testing I", func(t *testing.T) {
// 		result := FindNotesForChord("I", "major", &[]types.NBEFNoteRequest{{Note: nil}})
// 		println(len(result), "len of result")
// 		if *result[0].Note != "0" ||
// 			*result[1].Note != "2" ||
// 			*result[2].Note != "4" {
// 			t.Errorf("Expected %v, got %v", []int{0, 2, 4}, []string{*result[0].Note, *result[1].Note, *result[2].Note})
// 		}
// 	})

// 	t.Run("output a C major scale", func(t *testing.T) {
// 		key := "major"
// 		roman := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO"} // melodic minor?
// 		results := []string{}
// 		results = append(results, "key_type:major,key_note:C4,tempo:60")
// 		for _, r := range roman {
// 			chordInfo := types.NBEFNoteRequest{Note: nil, KeyType: "major", KeyNote: "C", TimeSec: "P"}
// 			notes := FindNotesForChord(r, key, &[]types.NBEFNoteRequest{chordInfo})
// 			for _, note := range notes {
// 				results = append(results, note.String())
// 			}
// 			results = append(results, "time:P+1/4\n")
// 			//t.Log(results)
// 		}

// 		out := "notes: \"" + strings.Join(results, "\n") + "\""
// 		t.Log(out)
// 		// out to file
// 		path := "/mnt/c/projects/music-user-reform/converter-standard-note"
// 		name := "c_major_scale.yml"
// 		os.WriteFile(path+"/"+name, []byte(out), 0644)
// 		//execute a system call and change directory to converter_standard_note
// 		//command := "." + path + "/standard-converter.exe --input " + name + " --output " + name
// 		// command := "ls /mnt/c/projects"
// 		// cmdSplit := strings.Split(command, " ")
// 		// cmd := exec.Command("/bin/sh", cmdSplit...)
// 		// t.Log(cmd)
// 		// outer := bytes.Buffer{}
// 		// cmd.Stdout = &outer
// 		// err := cmd.Run()
// 		// fmt.Printf("%s\n", outer.String())
// 		// if err != nil {
// 		// 	t.Error(err)
// 		// }

// 	})
// }
