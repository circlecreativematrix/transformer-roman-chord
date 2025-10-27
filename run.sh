onlyMaml=$1
pushd () {
    command pushd "$@" > /dev/null
}

popd () {
    command popd "$@" > /dev/null
}
if [ -z "$1" ]; 
then
    pushd /mnt/c/projects/music-user-reform/transformer-roman-to-standard
    go test -v fornof.me/m/v2/src/songs -run TestYamlMamlOutputChemistry
    #TestYamlMamlOutput 
    popd
    echo "Ran go"
fi
#/mnt/c/projects/music-user-reform/converter-standard-note/maml_test.yml
pushd /mnt/c/projects/music-user-reform/
    source venv-bash/bin/activate
    cd music-central/tests
    pytest -v test_roman_numerals.py
    echo "Ran python"
popd

