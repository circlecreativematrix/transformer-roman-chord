package main

func main() {
	// assuming key is C
	chord := ConvertNotationToChord("vii*", "major", 0)
	for _, note := range chord {
		println(note.Note)
	}

}
