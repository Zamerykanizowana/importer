package main

import (
	"flag"
	"github.com/Zamerykanizowana/importer/customerimporter"
)

func main() {
	var path, newPath string
	flag.StringVar(&path, "path", "", "csv file path to analize")
	flag.StringVar(&newPath, "newPath", "output.csv", "csv file path for output")
	flag.Parse()
	csvFile := customerimporter.NewCsvFile(path)
	m := csvFile.ReadFile()
	if newPath != "" {
		customerimporter.WriteFile(m, newPath)
	}
}
