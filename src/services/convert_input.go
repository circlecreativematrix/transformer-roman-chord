package services

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// todo, sort nbef by time before outputtting!
func ReadStandardFile(file string) string {
	data,
		err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

type StandardYamlIn struct {
	Notes string `yaml:"notes"`
}

func ReadStandardYamlFile(file string) StandardYamlIn {
	standardString := ReadStandardFile(file)
	return MarshalYamlStringToObj(standardString)
}
func MarshalYamlStringToObj(standardString string) StandardYamlIn {
	var standard StandardYamlIn
	println(standardString, "standardString")
	err := yaml.Unmarshal([]byte(standardString), &standard)
	if err != nil {
		log.Fatal(err)
	}
	return standard
}

// need nbef note
// func MarshalNBEFToString(nbef types.NBEF) string {
// 	data, err := yaml.Marshal(nbef)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return string(data)
// }
