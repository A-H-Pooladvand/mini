package file

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func Filenames(path string) (filenames []string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	return filenames
}

func Filename(filename string) string {
	return strings.TrimSuffix(
		filename, filepath.Ext(filename),
	)
}
