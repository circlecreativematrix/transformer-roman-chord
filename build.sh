# GOOS=js GOARCH=wasm go build -o chords.wasm main.go
tinygo build -o=chords.wasm -target=wasm -no-debug ./wasm.go
 mv chords.wasm "../music-central-web/public"