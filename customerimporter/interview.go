// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
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
	fmt.Println(headerLine)
	s := strings.Split(headerLine, ",")
	for index, element := range s {
		if element == "email" {
			fmt.Println(index)
			return index, nil
		}
	}
	return 0, errors.New("No email column in file")
}

func (f csvFile) ReadFile() (map[string]int, error) {
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
	log.Println("Procesing took", elapsed)
	return m, nil
}
