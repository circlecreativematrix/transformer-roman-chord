package src

import (
	"os"
	"strings"
	"testing"

	"fornof.me/m/v2/src/services"
	"fornof.me/m/v2/src/types"
)

func TestYamlMamlOutput(t *testing.T) {
	t.Run("testing output yaml", func(t *testing.T) {
		key := "major"
		roman := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO"}

		outNotes := []types.NBEFNote{}
		for _, r := range roman {

			chordInfo := types.NBEFNote{KeyType: "major", KeyNote: "C4", TimeSec: "P"}
			notes := FindNotesForChord(r, key, &[]types.NBEFNote{chordInfo})
			outNotes = append(outNotes, notes...)
			outNotes = append(outNotes, types.NBEFNote{Note: nil, KeyType: "major", KeyNote: "C4", TimeSec: "P+1/4"})

			//t.Log(results)
		}
		yamlOutMaml := services.StringNotesYaml(&outNotes)
		t.Log(yamlOutMaml)

		// out to file
		path := "/mnt/c/projects/music-user-reform/converter-standard-note"
		name := "maml_test.yml"
		os.WriteFile(path+"/"+name, []byte(yamlOutMaml), 0644)
	})
}
func TestFindNotesForChord(t *testing.T) {
	t.Run("testing I", func(t *testing.T) {
		result := FindNotesForChord("I", "major", &[]types.NBEFNote{{Note: nil}})
		println(len(result), "len of result")
		if *result[0].Note != 0 ||
			*result[1].Note != 2 ||
			*result[2].Note != 4 {
			t.Errorf("Expected %v, got %v", []int{0, 2, 4}, []int{*result[0].Note, *result[1].Note, *result[2].Note})
		}
	})

	t.Run("output a C major scale", func(t *testing.T) {
		key := "major"
		roman := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO"}
		results := []string{}
		results = append(results, "key_type:major,key_note:C4,tempo:60")
		for _, r := range roman {
			chordInfo := types.NBEFNote{Note: nil, KeyType: "major", KeyNote: "C", TimeSec: "P"}
			notes := FindNotesForChord(r, key, &[]types.NBEFNote{chordInfo})
			for _, note := range notes {
				results = append(results, note.String())
			}
			results = append(results, "time:P+1/4\n")
			//t.Log(results)
		}

		out := "notes: \"" + strings.Join(results, "\n") + "\""
		t.Log(out)
		// out to file
		path := "/mnt/c/projects/music-user-reform/converter-standard-note"
		name := "c_major_scale.yml"
		os.WriteFile(path+"/"+name, []byte(out), 0644)
		//execute a system call and change directory to converter_standard_note
		//command := "." + path + "/standard-converter.exe --input " + name + " --output " + name
		// command := "ls /mnt/c/projects"
		// cmdSplit := strings.Split(command, " ")
		// cmd := exec.Command("/bin/sh", cmdSplit...)
		// t.Log(cmd)
		// outer := bytes.Buffer{}
		// cmd.Stdout = &outer
		// err := cmd.Run()
		// fmt.Printf("%s\n", outer.String())
		// if err != nil {
		// 	t.Error(err)
		// }

	})
}
