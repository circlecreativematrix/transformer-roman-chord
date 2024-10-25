# GOOS=js GOARCH=wasm go build -o chords.wasm main.go
tinygo build -o=chords.wasm -target=wasm -no-debug ./wasm.go
 mv chords.wasm "/mnt/c/projects/music-user-reform/music-central-web/public"