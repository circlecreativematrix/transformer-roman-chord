package src

import (
	"testing"

	//"fornof.me/m/v2/src"
	"fornof.me/m/v2/src/types"
)

func TestFindNotesForChord(t *testing.T) {
	tests := []struct {
		name     string
		roman    string
		expected []types.NBEFNote
		keyType  string
	}{
		{
			name:    "Test for chord I",
			roman:   "I",
			keyType: "major",
			expected: []types.NBEFNote{
				types.NBEFNote{Note: 0},
				types.NBEFNote{Note: 2},
				types.NBEFNote{Note: 4},
			}, // Assuming key of C major for simplicity; adjust as necessary
		},
		{
			name:    "Test for chord #IV+",
			roman:   "#IV+",
			keyType: "major",
			expected: []types.NBEFNote{

				types.NBEFNote{Note: 0 + 3, Halfsteps: 1},
				types.NBEFNote{Note: 2 + 3},
				types.NBEFNote{Note: 4 + 3, Halfsteps: 1},
			}, // Assuming key of C major and #IV+ means augmented fourth; adjust as necessary
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindNotesForChord(tt.roman, tt.keyType, []int{0, 2, 4})
			if (result[0].Note != tt.expected[0].Note ||
				result[1].Note != tt.expected[1].Note ||
				result[2].Note != tt.expected[2].Note) &&
				(result[0].Halfsteps != tt.expected[0].Halfsteps ||
					result[2].Halfsteps != tt.expected[2].Halfsteps ||
					result[1].Halfsteps != tt.expected[1].Halfsteps) {
				t.Errorf("Expected %v, got %v", tt.expected[0].Note, result[0].Note)
			}
		})
	}
}
