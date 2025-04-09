package songs

// arpeggios
// r
// dominant
// triplets - minus second note ? minus third note? two back to back/ swing
// trills
//
import (
	"testing"
	"time"

	"fornof.me/m/v2/src/types"
)

func singleChorus() []string {
	outString := []string{}
	addToOutput([]string{"time:0,dur:0,track:0"}, &outString)
	addToOutput(header("intro", key_note[0], "major", "60"), &outString)
	addToOutput([]string{"chord_type:major,key_note:F3,tempo:60"}, &outString)
	// addToOutput(chords([]string{"iii", "V/5", "IV/5", "I"}, []Beat{
	// 	{Beat: "P+1/2", Count: 3},
	// 	{Beat: "P+1/2", Count: 2},
	// }, []string{"vol:100"}), &outString)
	// melodyLower([]string{"I", "iii/5", "vi", "IV/3"}, 0, 0, nil)
	// // should be same as top, need to get a read on it.
	// melodyLower([]string{"Am", "C/E", "B@/F", "F"}, 0, 0, nil)
	// melodyLower([]string{"F", "Am", "Dm", "B@", "F"}, 0, 0, nil)
	// melody([]string{"0", "1", "2", "3"}, 0, &outString)
	// melody([]string{"0", "1", "2", "3"}, 0, &outString)
	// melody([]string{"0", "1", "2", "3"}, 0, &outString)
	// melody([]string{"0", "1", "2", "3"}, 0, &outString)
	// addToOutput([]string{"time:0"}, &outString)
	// addToOutput([]string{"time:P+8/8,dur:0,track:1"}, &outString)
	// addToOutput(header("intro-lower", key_note[1], "major", "100"), &outString)
	//"iii", "V/5", "IV/5", "I", "I","iii/5","vi", "IV/3"
	//"V", "IV", "I", "I", "iii", "vi", "IV"
	//"V", "IV", "I", "I", "iii", "vi", "IV"

	chorusStuff(&outString)
	chorusStuff(&outString)
	chorusStuff(&outString)
	chorusStuff(&outString)
	chorusStuff(&outString)
	return outString
}
func chorusStuff(outString *[]string) {
	melodyLower([]string{"iii", "iii", "iii", ""}, 0, 0, outString)
	melodyLower([]string{"V", "V", "V", ""}, 0, 0, outString)
	melodyLower([]string{"IV", "IV", "IV", ""}, 0, 0, outString)
	melodyLower([]string{"I", "I", "I", ""}, 0, 0, outString)
	melodyLower([]string{"I", "I", "I", ""}, 0, 0, outString)
	melodyLower([]string{"iii", "iii", "iii", ""}, 0, 0, outString)
	melodyLower([]string{"vi", "vi", "vi", ""}, 0, 0, outString)
	melodyLower([]string{"IV", "IV", "IV", ""}, 0, 0, outString)
}
func TestYamlMamlOutput(t *testing.T) {
	t.Run("testing output yaml", func(t *testing.T) {

		//roman := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO", "I"}
		//roman := []string{"i", "iiO", "III", "iv", "v", "VI", "VII"}

		phrases := types.PhraseList{}
		phrases = append(phrases, types.Phrase{
			Type:       "fornof.standard",
			Name:       "chorus",
			InNotes:    singleChorus(),
			OutputMidi: "chorus.mid"})
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
		MakeTheSong(phrases)
		//sleep 5 seconds
		time.Sleep(2 * time.Second)
	})
}
