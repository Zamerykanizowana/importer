package main

import (
	"flag"
	"github.com/Zamerykanizowana/importer/customerimporter"
)

func main() {
	var path, output string
	flag.StringVar(&path, "path", "", "csv file path to analize")
	flag.StringVar(&output, "output", "output.csv", "csv file path for output")
	flag.Parse()
	csvFile := customerimporter.NewCsvFile(path)
	m := csvFile.ReadFile()
	customerimporter.WriteFile(m, output)
}
