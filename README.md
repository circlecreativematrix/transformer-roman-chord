# scales definitions: 
C major scale: 
```
roman := []string{
 "chord:I,split:0",
 "chord:ii", 
 "chord:iii", 
 "chord:IV", 
 "chord:V",
 "chord:vi", 
 "chord:viiO", 
 "chord:I",}
```
C minor_natural scale
```
roman := []string{
 "chord:i,split:0",
 "chord:iiO", 
 "chord:III", 
 "chord:iv", 
 "chord:v",
 "chord:VI", 
 "chord:VII",
 "chord:i,octave:1,}
 ```
# questions: 
What happens when I play a vii in a C major scale that has the happy path as viiO? 
What happens when it is VII or VIIO ? 
	roman := `
			    chord:I,split:0,chord_type:major,key_type:major,key_note:C4,time:P+1/4
				chord:ii,time:P+1/4
				chord:iii,time:P+1/4
				chord:IV,time:P+1/4
				chord:V,time:P+1/4
				chord:vi,time:P+1/4
				chord:viiO,time:P+1/4
				chord:I,offset:7,time:P+1/4`

# How to : 
1. get Chord F#M : chord:F#M, 
2. get type "major"/ minor - by M/m 
3. isolate F# --> midi number (44)
4. first pattern for major +3 +3 
5. convert midi back into notes.
6. place notes into succession with eachother, first one with timing  

- todo: make http endpoints and draw a picture of how to make it work / dockerfiles of each 
- todo: possibly redo the recording in the browser to nbef 
- figure out why halfsteps are not working in browser
