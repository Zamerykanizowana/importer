package main

import (
	"flag"
	"github.com/Zamerykanizowana/importer/customerimporter"
)

func main() {
    var path string
    flag.StringVar(&path, "path", "", "csv file path to analize")
    flag.Parse()
    csvFile := customerimporter.NewCsvFile(path)
	csvFile.ReadFile()
}
