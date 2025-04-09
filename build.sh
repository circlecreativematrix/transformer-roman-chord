output_dir=${1:-"../music-central-web/public"}
echo 'output main.wasm to' $output_dir
tinygo build -o=chords.wasm -target=wasm -no-debug ./wasm.go
 mv chords.wasm "../music-central-web/public"