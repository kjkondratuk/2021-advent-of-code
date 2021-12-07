package lib

import (
	"io/ioutil"
	"log"
	"strings"
)

var (
	defaultDataLoader = func(data []string) interface{} {
		return data
	}
)

type DataLoader func([]string) interface{}

type dataReader struct {
	file   string
	loader DataLoader
}

type DataReader interface {
	Read() interface{}
}

func NewDataReader(file string) DataReader {
	return &dataReader{
		file:   file,
		loader: defaultDataLoader,
	}
}

func NewDataReaderWithLoader(file string, loader DataLoader) DataReader {
	return &dataReader{
		file:   file,
		loader: loader,
	}
}

func (r *dataReader) Read() interface{} {
	f, err := ioutil.ReadFile(r.file)
	if err != nil {
		log.Fatalf("Error reading input file: [%s]", err)
	}

	return r.loader(strings.Split(string(f), "\n"))
}
