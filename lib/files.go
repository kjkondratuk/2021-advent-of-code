package lib

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadData(file string) []string {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading input file: [%s]", err)
	}

	return strings.Split(string(f), "\n")
}
