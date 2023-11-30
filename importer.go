package main

import (
	"github.com/Zamerykanizowana/importer/customerimporter"
    "os"
)



func main() {
    csvFile := customerimporter.NewCsvFile(os.Args[1])
    csvFile.ReadFile()
}
