package services

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var Letter_to_midi = make(map[string]int)
var Midi_to_letter = make(map[int]string)
var Types = make(map[string]map[int]int)
var Self = make(map[string]interface{})

type NoteBag struct {
	StartNote int
	KeyType   map[int]int
}

var BagConvert = new(NoteBag)

func InitConvert() {
	Letter_to_midi = map[string]int{
		"c0":  0,
		"c1":  12,
		"c2":  24,
		"c3":  36,
		"c4":  48,
		"c5":  60,
		"c6":  72,
		"c7":  84,
		"c8":  96,
		"c9":  108,
		"c10": 120,

		"c0#":  1,
		"c1#":  13,
		"c2#":  25,
		"c3#":  37,
		"c4#":  49,
		"c5#":  61,
		"c6#":  73,
		"c7#":  85,
		"c8#":  97,
		"c9#":  109,
		"c10#": 121,

		"d0@":  1,
		"d1@":  13,
		"d2@":  25,
		"d3@":  37,
		"d4@":  49,
		"d5@":  61,
		"d6@":  73,
		"d7@":  85,
		"d8@":  97,
		"d9@":  109,
		"d10@": 121,

		"d0":  2,
		"d1":  14,
		"d2":  26,
		"d3":  38,
		"d4":  50,
		"d5":  62,
		"d6":  74,
		"d7":  86,
		"d8":  98,
		"d9":  110,
		"d10": 122,

		"d0#":  3,
		"d1#":  15,
		"d2#":  27,
		"d3#":  39,
		"d4#":  51,
		"d5#":  63,
		"d6#":  75,
		"d7#":  87,
		"d8#":  99,
		"d9#":  111,
		"d10#": 123,

		"e0@":  3,
		"e1@":  15,
		"e2@":  27,
		"e3@":  39,
		"e4@":  51,
		"e5@":  63,
		"e6@":  75,
		"e7@":  87,
		"e8@":  99,
		"e9@":  111,
		"e10@": 123,

		"e0":  4,
		"e1":  16,
		"e2":  28,
		"e3":  40,
		"e4":  52,
		"e5":  64,
		"e6":  76,
		"e7":  88,
		"e8":  100,
		"e9":  112,
		"e10": 124,

		"f0":  5,
		"f1":  17,
		"f2":  29,
		"f3":  41,
		"f4":  53,
		"f5":  65,
		"f6":  77,
		"f7":  89,
		"f8":  101,
		"f9":  113,
		"f10": 125,

		"f0#":  6,
		"f1#":  18,
		"f2#":  30,
		"f3#":  42,
		"f4#":  54,
		"f5#":  66,
		"f6#":  78,
		"f7#":  90,
		"f8#":  102,
		"f9#":  114,
		"f10#": 126,

		"g0@":  6,
		"g1@":  18,
		"g2@":  30,
		"g3@":  42,
		"g4@":  54,
		"g5@":  66,
		"g6@":  78,
		"g7@":  90,
		"g8@":  102,
		"g9@":  114,
		"g10@": 126,

		"g0":  7,
		"g1":  19,
		"g2":  31,
		"g3":  43,
		"g4":  55,
		"g5":  67,
		"g6":  79,
		"g7":  91,
		"g8":  103,
		"g9":  115,
		"g10": 127,

		"g0#": 8,
		"g1#": 20,
		"g2#": 32,
		"g3#": 44,
		"g4#": 56,
		"g5#": 68,
		"g6#": 80,
		"g7#": 92,
		"g8#": 104,
		"g9#": 116,

		"a0@": 8,
		"a1@": 20,
		"a2@": 32,
		"a3@": 44,
		"a4@": 56,
		"a5@": 68,
		"a6@": 80,
		"a7@": 92,
		"a8@": 104,
		"a9@": 116,

		"a0": 9,
		"a1": 21,
		"a2": 33,
		"a3": 45,
		"a4": 57,
		"a5": 69,
		"a6": 81,
		"a7": 93,
		"a8": 105,
		"a9": 117,

		"a0#": 10,
		"a1#": 22,
		"a2#": 34,
		"a3#": 46,
		"a4#": 58,
		"a5#": 70,
		"a6#": 82,
		"a7#": 94,
		"a8#": 106,
		"a9#": 118,

		"b0@": 10,
		"b1@": 22,
		"b2@": 34,
		"b3@": 46,
		"b4@": 58,
		"b5@": 70,
		"b6@": 82,
		"b7@": 94,
		"b8@": 106,
		"b9@": 118,

		"b0": 11,
		"b1": 23,
		"b2": 35,
		"b3": 47,
		"b4": 59,
		"b5": 71,
		"b6": 83,
		"b7": 95,
		"b8": 107,
		"b9": 119,
	}
	Midi_to_letter = map[int]string{
		0:   "c0",
		12:  "c1",
		24:  "c2",
		36:  "c3",
		48:  "c4",
		60:  "c5",
		72:  "c6",
		84:  "c7",
		96:  "c8",
		108: "c9",
		120: "c10",

		1:   "c0#",
		13:  "c1#",
		25:  "c2#",
		37:  "c3#",
		49:  "c4#",
		61:  "c5#",
		73:  "c6#",
		85:  "c7#",
		97:  "c8#",
		109: "c9#",
		121: "c10#",

		2:   "d0@",
		14:  "d1",
		26:  "d2",
		38:  "d3",
		50:  "d4",
		62:  "d5",
		74:  "d6",
		86:  "d7",
		98:  "d8",
		110: "d9",
		122: "d10",

		3:   "d0#",
		15:  "d1#",
		27:  "d2#",
		39:  "d3#",
		51:  "d4#",
		63:  "d5#",
		75:  "d6#",
		87:  "d7#",
		99:  "d8#",
		111: "d9#",
		123: "d10#",

		4:   "e0@",
		16:  "e1",
		28:  "e2",
		40:  "e3",
		52:  "e4",
		64:  "e5",
		76:  "e6",
		88:  "e7",
		100: "e8",
		112: "e9",
		124: "e10",

		5:   "f0",
		17:  "f1",
		29:  "f2",
		41:  "f3",
		53:  "f4",
		65:  "f5",
		77:  "f6",
		89:  "f7",
		101: "f8",
		113: "f9",
		125: "f10",

		6:   "f0#",
		18:  "f1#",
		30:  "f2#",
		42:  "f3#",
		54:  "f4#",
		66:  "f5#",
		78:  "f6#",
		90:  "f7#",
		102: "f8#",
		114: "f9#",
		126: "f10#",

		7:   "g0@",
		19:  "g1",
		31:  "g2",
		43:  "g3",
		55:  "g4",
		67:  "g5",
		79:  "g6",
		91:  "g7",
		103: "g8",
		115: "g9",
		127: "g10",

		8:   "g0#",
		20:  "g1#",
		32:  "g2#",
		44:  "g3#",
		56:  "g4#",
		68:  "g5#",
		80:  "g6#",
		92:  "g7#",
		104: "g8#",
		116: "g9#",

		9:   "a0@",
		21:  "a1",
		33:  "a2",
		45:  "a3",
		57:  "a4",
		69:  "a5",
		81:  "a6",
		93:  "a7",
		105: "a8",
		117: "a9",

		10:  "a0#",
		22:  "a1#",
		34:  "a2#",
		46:  "a3#",
		58:  "a4#",
		70:  "a5#",
		82:  "a6#",
		94:  "a7#",
		106: "a8#",
		118: "a9#",

		11:  "b0@",
		23:  "b1",
		35:  "b2",
		47:  "b3",
		59:  "b4",
		71:  "b5",
		83:  "b6",
		95:  "b7",
		107: "b8",
		119: "b9",
	}
	Types = map[string]map[int]int{
		// Major
		// R W W H W W W H
		// 0 1 2 3 4 5 6
		"major": {
			0: 0,
			1: 2,
			2: 2 * 2,
			3: 2*2 + 1,
			4: 2*3 + 1,
			5: 2*4 + 1,
			6: 2*5 + 1,
			7: 2*5 + 1*2,
		},
		"major_pentatonic": {
			0: 0,
			1: 2,
			2: 2 * 2,
			3: 0,
			4: 2*3 + 1,
			5: 2*4 + 1,
			6: 0,
			7: 2*5 + 1*2,
		},
		//natural minor
		//R W H W W H W W
		//0 1 2 3 4 5 6
		"minor_natural": {
			0: 0,
			1: 2,
			2: 2 + 1,
			3: 2*2 + 1,
			4: 2*3 + 1,
			5: 2*3 + 1*2,
			6: 2*4 + 1*2,
		},

		//harmonic minor
		//R W H W W H W+H H
		//0 1 2 3 4 5 6
		"minor_harmonic": {
			0: 0,
			1: 2,
			2: 2 + 1,
			3: 2*2 + 1,
			4: 2*3 + 1,
			5: 2*3 + 1*2,
			6: 2*4 + 1*3,
		},
		//harmonic minor
		//R W H W W H W+H H
		//0 1 2 3 4 5 6
		"minor": {
			0: 0,
			1: 2,
			2: 2 + 1,
			3: 2*2 + 1,
			4: 2*3 + 1,
			5: 2*3 + 1*2,
			6: 2*4 + 1*3,
		},
		"minor_harmonic_pentatonic": {
			0: 0,
			1: 2,
			2: 2 + 1,
			3: 0,
			4: 2*3 + 1,
			5: 2*3 + 1*2,
			6: 0,
		},
		//melodic ascending
		//R W H W W W W H
		//0 1 2 3 4 5 6
		"minor_melodic_ascending": {
			0: 0,
			1: 2,
			2: 2 + 1,
			3: 2*2 + 1,
			4: 2*3 + 1,
			5: 2*4 + 1,
			6: 2*5 + 1,
		},
		//melodic descending -may have this reversed
		//R W W H W W H W
		//0 1 2 3 4 5 6
		"minor_melodic_descending": {
			0: 0,
			1: 2,
			2: 2 * 2,
			3: 2*2 + 1,
			4: 2*3 + 1,
			5: 2*4 + 1,
			6: 2*4 + 1*2,
		},
	}
	BagConvert.StartNote = Letter_to_midi["c4"] //HARDCODED
	BagConvert.KeyType = Types["minor_harmonic"]
}
func SetStartNote(note string) {
	BagConvert.StartNote = Letter_to_midi[strings.ToLower(note)]

}
func SetKeyType(type_in string) {
	//TODO: come up with things here
	BagConvert.KeyType = Types[type_in]

}

// takes note and returns offset
func CalculateOffset(my_note string) (int, int, int) {
	minus := 1
	if my_note == "" {
		return 0, 0, 0
	}
	myNoteInt, _ := strconv.ParseInt(my_note, 10, 64)
	if myNoteInt < 0 {
		minus = -1
	}

	note := myNoteInt % 7
	if note < 0 {
		note = 7 + note
	}
	//print('note!', note)
	octavesInChromaticNotes := 0
	//offset = math.floor((my_note - note) / 7) * 12
	//print(my_note ,'note', note, 'note', math.floor((my_note - note) / 7) * 12)
	if minus == 1 {
		octavesInChromaticNotes = int(math.Abs(float64(myNoteInt-note))/7) * 12
	} else {
		octavesInChromaticNotes = int(math.Abs(float64(myNoteInt-note))/7) * 12
	}
	//print('offset', offset)

	return octavesInChromaticNotes, int(note), minus

}

func HandleKeyType(noteIn string, sharp_flat int) (int, error) {
	offset, note, minus := CalculateOffset(noteIn)
	//print('sharpflat', sharp_flat)
	added_offsets := BagConvert.StartNote + sharp_flat + offset*minus //+ note_offset_from_yaml_read +octave*7
	//print('added_offsets', added_offsets)
	keyType, ok := BagConvert.KeyType[note]
	if !ok {
		return -999, errors.New("note not found, please use minor_harmonic, minor_melodic, or major for key_type")
	}

	return keyType + added_offsets, nil

}

// func named_note_offset(note string, offset int, use_accidental bool) (int, error) {
// 	name := ""
// 	octave := 3
// 	accidental := ""
// 	number := 0
// 	if len(note) == 3 {
// 		name = string(note[0])
// 		octaveIn, _ := strconv.Atoi(string(note[1]))
// 		octave = int(octaveIn)
// 		accidental = string(note[2])
// 	}
// 	if len(note) == 2 {
// 		name = string(note[0])
// 		octaveIn, _ := strconv.Atoi(string(note[1]))
// 		octave = int(octaveIn)
// 	}
// 	start := Midi_to_letter[BagConvert.StartNote]
// 	resulting_accidental := 0.0
// 	if len(accidental) > 0 && use_accidental {
// 		if accidental == "#" {
// 			resulting_accidental = .6 //sharp
// 		}
// 		if accidental == "@" {
// 			resulting_accidental = .4 //flat - todo change this.
// 		}
// 	}
// 	startOctave, err := strconv.Atoi(string(start[1]))
// 	if err != nil {
// 		return -999, err
// 	}

// 	octave_out := math.Abs(float64(octave-startOctave)) * 7
// 	if int(octave) >= int(start[1]) {
// 		//search positive
// 		number = int(name[0]) - int(start[0])
// 		return number + int(octave_out) + offset + int(resulting_accidental), nil
// 	} else {
// 		//search negative.
// 		number = int(start[0]) - int(name[0]) - 1 //offset of -1
// 		return number - int(octave_out) + offset + int(resulting_accidental), nil
// 	}
// }

// func convert_duration_to_beat( duration, bpm){
// //250 ms
// bps = bpm / 60.0
// spb = 1 / bps  //seconds per beat
// mspb = 1000 * spb  //ms per beat
// print(duration, 'duration', mspb , 'mspb')
// fraction = Fraction(int(duration)*100, int(round(mspb, 0))).limit_denominator(128)
// if str(fraction) == 1:
// 	return "1/1"

// if len(str(fraction).split("/")) == 1:
// 	print(str(fraction) + "/1", 'outputty')
// 	return str(fraction)+"/1"
// return str(fraction)

// }
// func convert_beat_to_dur_single( beat, bpm){
// bps = bpm / 60.0
// spb = 1 / bps  //seconds per beat
// mspb = 1000 * spb  //ms per beat
// //get first
// //get second
// str_beat = str(beat).split("/")
// //assuming no non-decimal numbers
// //assuming last is not zero
// ##print(str_beat, 'beat')
// fraction = float(str_beat[0]) / float(str_beat[1])
// if fraction < 0:
// 	//11/7/2023 - this is confusing as $!%^, I don't like it
// 	//but it's there, might as well use it for now , come back
// 	//later and change it
// 	//its a rest if its negative!
// 	##print('added_midi_note -1 ')
// 	//"dur": fraction * mspb *4*-1,

// 	return {"dur": fraction * mspb * 4  * -1, "midi_note": -1}
// else:
// 	return {"dur": fraction * mspb * 4 }

// }
// func convert_beat_to_dur( obj, bpm){
// beat = obj.get('beat', None)
// if not beat:
// 	return 1
// bps = bpm / 60.0
// spb = 1 / bps  //seconds per beat
// mspb = 1000 * spb  //ms per beat
// //get first
// //get second
// str_beat = str(beat).split(".")
// //assuming no non-decimal numbers
// //assuming last is not zero
// fraction = float(str_beat[0]) / float(str_beat[1])

// if fraction < 0:
// 	//its a rest if its negative!
// 	obj["midi_note"] = -1
// 	##print('added_midi_note -1 ')
// 	del obj['note']
// 	obj['dur'] = fraction * mspb * -1
// else:
// 	obj['dur'] = fraction * mspb

// }
// func load_yaml( path){
// with open(path, 'r') as file:
// 	return yaml.safe_load(file.read())

// }	Mi
// func get_type( obj){
// matcher = re.search(r"'(.*)'", str(type(obj)))
// return matcher.group(1)

// }
// func format_to_nbef_notes( note, beat, signal, velocity, tempo, track = 0){
// return {'beat_type': 'signal_ms',
// 		'note_type': 'midi',
// 		'midi': note,
// 		'track': track,
// 		'tempo': tempo,
// 		'velocity': velocity,
// 		'time_s': beat,
// 		'signal': signal}

// }
// func format_to_nbef( notes, notebeatArrayYaml){
// return {
// 	'note_type': 'midi',
// 	'beat_type': 'signal_ms',
// 	'tempo': notebeatArrayYaml['tempo'],
// 	'notes': notes
// }
