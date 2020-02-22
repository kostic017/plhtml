package main

import (
	"io/ioutil"
	"path/filepath"
)

func readFromFile(file string) string {
	filebuffer, err := ioutil.ReadFile(filepath.FromSlash(file))

	if err != nil {
		panic(err)
	}

	return string(filebuffer)
}
