package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) (string, error) {
	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	return string(filebuffer), nil
}
