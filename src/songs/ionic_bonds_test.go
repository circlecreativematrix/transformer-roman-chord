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

func singleChorusChemistry() []string {
	outString := []string{}
	addToOutput([]string{"time:0,dur:0,track:0"}, &outString)
	addToOutput(header("intro", key_note[0], "major", "60"), &outString)
	addToOutput([]string{"chord_type:major,key_note:F3,tempo:60"}, &outString)
	progressionIntro := []string{"III", "II", "I", "II"}
	progressionVerse := []string{"IV", "III", "VI", "II"}

	chorusUnderlick(&outString, progressionIntro)
	chorusUnderlick(&outString, progressionIntro)
	chorusUnderlick(&outString, progressionVerse)
	chorusUnderlick(&outString, progressionVerse)
	chorusUnderlick(&outString, progressionIntro)
	chorusUnderlick(&outString, progressionIntro)
	chorusUnderlick(&outString, progressionVerse)
	chorusUnderlick(&outString, progressionVerse)
	return outString
}

func chorusUnderlick(outString *[]string, progression []string) {
	for chord := range progression {
		melodyLower([]string{progression[chord], progression[chord], progression[chord], ""}, 0, 0, outString)
	}
}

func TestYamlMamlOutputChemistry(t *testing.T) {
	t.Run("testing output yaml", func(t *testing.T) {

		//roman := []string{progression[2], "ii", "iii", "IV", "V", "vi", "viiO", progression[2]}
		//roman := []string{"i", "iiO", "III", "iv", "v", "VI", "VII"}

		phrases := types.PhraseList{}
		phrases = append(phrases, types.Phrase{
			Type:       "fornof.standard",
			Name:       "chorus",
			InNotes:    singleChorusChemistry(),
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
