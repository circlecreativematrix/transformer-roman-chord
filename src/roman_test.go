package src

// func TestMajorChordsGenerateProperly(t *testing.T) {
// 	roman := []string{"I", "ii", "iii", "IV", "V", "vi", "viiO"}
// 	HalfstepNoteExpectations := [][]types.NBEFNoteRequest{}

// 	for offset, _ := range roman {
// 		one := strconv.Itoa(0 + offset)
// 		three := strconv.Itoa(2 + offset)
// 		five := strconv.Itoa(4 + offset)
// 		HalfstepNoteExpectations = append(HalfstepNoteExpectations, []types.NBEFNoteRequest{
// 			{Note: &one, Halfsteps: 0},
// 			{Note: &three, Halfsteps: 0},
// 			{Note: &five, Halfsteps: 0},
// 		})
// 		if offset == 6 {
// 			HalfstepNoteExpectations[offset][2].Halfsteps = -1
// 		}
// 	}

// 	for offset, r := range roman {
// 		chordInfo := types.NBEFNoteRequest{Note: nil, KeyType: "major", KeyNote: "C4", TimeSec: "P"}
// 		notes := FindNotesForChord(r, "major", &[]types.NBEFNoteRequest{chordInfo})
// 		for i, note := range notes {
// 			assert.Equal(t, *note.Note, *HalfstepNoteExpectations[offset][i].Note)
// 			assert.Equal(t, note.Halfsteps, HalfstepNoteExpectations[offset][i].Halfsteps)
// 		}

// 	}
// }
// func TestMinorNaturalChordsGenerateProperly(t *testing.T) {
// 	roman := []string{"i", "iiO", "III", "iv", "v", "VI", "VII"}
// 	HalfstepNoteExpectations := [][]types.NBEFNoteRequest{}

// 	for offset, _ := range roman {
// 		one := strconv.Itoa(0 + offset)
// 		three := strconv.Itoa(2 + offset)
// 		five := strconv.Itoa(4 + offset)
// 		HalfstepNoteExpectations = append(HalfstepNoteExpectations, []types.NBEFNoteRequest{
// 			{Note: &one, Halfsteps: 0},
// 			{Note: &three, Halfsteps: 0},
// 			{Note: &five, Halfsteps: 0},
// 		})
// 		if offset == 6 {
// 			HalfstepNoteExpectations[offset][2].Halfsteps = -1
// 		}
// 	}

// 	for offset, r := range roman {
// 		chordInfo := types.NBEFNoteRequest{Note: nil, KeyType: types.Constants.MinorNatural, KeyNote: "C", TimeSec: "P"}
// 		notes := FindNotesForChord(r, types.Constants.MinorNatural, &[]types.NBEFNoteRequest{chordInfo})
// 		for i, note := range notes {
// 			assert.Equal(t, note.Note, HalfstepNoteExpectations[offset][i].Note)
// 			assert.Equal(t, note.Halfsteps, HalfstepNoteExpectations[offset][i].Halfsteps)
// 		}

// 	}
// }

/*
*
i, minor, (1, b3, 5)
ii, minor, (2, 4, 6)
III, augmented, (b3, 5, 7)
IV, Major, (4, 6, 8)
V, Major, (5, 7, 2)
vi, diminished, (6, 8, b3)
vii, diminished, (7, 2, 4)
*
*/

// func TestMinorMelodicChordsGenerateProperly(t *testing.T) {
// 	roman := []string{"i", "ii", "III", "IV", "V", "viO", "viiO"}
// 	HalfstepNoteExpectations := [][]types.NBEFNoteRequest{}
// 	for offset, _ := range roman {
// 		one := strconv.Itoa(0 + offset)
// 		three := strconv.Itoa(2 + offset)
// 		five := strconv.Itoa(4 + offset)
// 		HalfstepNoteExpectations = append(HalfstepNoteExpectations, []types.NBEFNoteRequest{
// 			{Note: &one, Halfsteps: 0},
// 			{Note: &three, Halfsteps: 0},
// 			{Note: &five, Halfsteps: 0},
// 		})
// 		if offset == 6 || offset == 5 {
// 			HalfstepNoteExpectations[offset][2].Halfsteps = -1
// 		}
// 	}

// 	for offset, r := range roman {
// 		chordInfo := types.NBEFNoteRequest{Note: nil, KeyType: types.Constants.MinorNatural, KeyNote: "C", TimeSec: "P"}
// 		notes := FindNotesForChord(r, types.Constants.MinorMelodic, &[]types.NBEFNoteRequest{chordInfo})
// 		for i, note := range notes {
// 			assert.Equal(t, note.Note, HalfstepNoteExpectations[offset][i].Note)
// 			assert.Equal(t, note.Halfsteps, HalfstepNoteExpectations[offset][i].Halfsteps)
// 		}

// 	}
// }

// func TestMinorHarmonicChordsGenerateProperly(t *testing.T) {
// 	//i,iiO, III+,iv,V,VI,viiO
// 	roman := []string{"i", "iiO", "III+", "iv", "V", "VI", "viiO"}
// 	HalfstepNoteExpectations := [][]types.NBEFNoteRequest{}
// 	for offset, _ := range roman {
// 		one := strconv.Itoa(0 + offset)
// 		three := strconv.Itoa(2 + offset)
// 		five := strconv.Itoa(4 + offset)
// 		HalfstepNoteExpectations = append(HalfstepNoteExpectations, []types.NBEFNoteRequest{
// 			{Note: &one, Halfsteps: 0},
// 			{Note: &three, Halfsteps: 0},
// 			{Note: &five, Halfsteps: 0},
// 		})
// 		if offset == 2 {
// 			HalfstepNoteExpectations[offset][2].Halfsteps = 1
// 		}
// 		if offset == 1 || offset == 6 {
// 			HalfstepNoteExpectations[offset][2].Halfsteps = -1
// 		}
// 	}

// 	for offset, r := range roman {
// 		chordInfo := types.NBEFNoteRequest{Note: nil, KeyType: types.Constants.MinorNatural, KeyNote: "C", TimeSec: "P"}
// 		notes := FindNotesForChord(r, types.Constants.MinorHarmonic, &[]types.NBEFNoteRequest{chordInfo})
// 		for i, note := range notes {
// 			t.Logf("note: %s", r)
// 			assert.Equal(t, note.Note, HalfstepNoteExpectations[offset][i].Note)
// 			assert.Equal(t, note.Halfsteps, HalfstepNoteExpectations[offset][i].Halfsteps)
// 		}

// 	}
// }
