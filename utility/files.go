package utility

import (
	"io/ioutil"
	"path/filepath"
)

func ReadFile(file string) string {
	data, err := ioutil.ReadFile(filepath.FromSlash(file))
	Check(err)
	return string(data)
}

func WriteFile(file string, text string) {
	// Only the owner can read and write. Everyone else can only read. No one can execute the file.
	err := ioutil.WriteFile(filepath.FromSlash(file), []byte(text), 0644)
	Check(err)
}
