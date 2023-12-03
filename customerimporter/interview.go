// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type csvFile struct {
	path string
}

func NewCsvFile(path string) *csvFile {
	f := csvFile{path: path}
	return &f
}

func findMailColumn(headerLine string) (int, error) {
	s := strings.Split(headerLine, ",")
	for index, element := range s {
		if element == "email" {
			return index, nil
		}
	}
	return 0, errors.New("No email column in file")
}

func (f csvFile) ReadFile() map[string]int {
	log.Println("Start processing...")
	start := time.Now()
	readFile, err := os.Open(f.path)
	defer readFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	i, err := findMailColumn(fileScanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]int)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), ",")
		email := strings.Split(s[i], "@")
		if len(email) != 2 {
			log.Println(fmt.Sprintf("Problem with parsing email: %s, skipping", s[i]))
		} else {
			if _, ok := m[email[1]]; ok {
				m[email[1]] += 1
			} else {
				m[email[1]] = 1
			}
		}
	}

	elapsed := time.Since(start)
	log.Println("Reading data and making map took", elapsed)
	return m
}

func WriteFile(m map[string]int, path string) {
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	if err = w.Write([]string{"domain", "number of customers"}); err != nil {
		log.Fatal("error writing headers to file", err)
	}

	for domain, acc := range m {
		if err = w.Write([]string{domain, fmt.Sprintf("%v", acc)}); err != nil {
			log.Fatal("error writing records to file", err)
		}
	}

    
}
