package main

import (
	"io/ioutil"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFromFile(file string) string {
	data, err := ioutil.ReadFile(filepath.FromSlash(file))
	check(err)
	return string(data)
}

func writeToFile(file string, text string) {
	// Only the owner can read and write. Everyone else can only read. No one can execute the file.
	err := ioutil.WriteFile(filepath.FromSlash(file), []byte(text), 0644)
	check(err)
}
