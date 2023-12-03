package customerimporter

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
)

func mapFromCsv(path string) (map[string]int, error) {
	readFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer readFile.Close()

	m := make(map[string]int)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// to omit first line with headers
	fileScanner.Scan()

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), ",")
		m[s[0]], err = strconv.Atoi(s[1])
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func checkCounting(expected, received map[string]int) error {
	for domain, occ := range expected {
		if received[domain] != occ {
			return errors.New("Wrong number of occurrence")
		}
	}
	return nil
}

func TestReorderedColumn(t *testing.T) {
	expected, err := mapFromCsv("testdata/customers_reordered_columns_output.csv")
	if err != nil {
		t.Errorf("Fail test")
	}

	received := NewCsvFile("testdata/customers_reordered_columns.csv").ReadFile()
	err = checkCounting(expected, received)
	if err != nil {
		t.Errorf("%v", err)
	}

	receivedFilePath := "testdata/test_reordered_column_output.csv"
	WriteFile(received, receivedFilePath)
	defer os.Remove(receivedFilePath)
	received, err = mapFromCsv(receivedFilePath)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = checkCounting(expected, received)
	if err != nil {
		t.Errorf("%v", err)
	}
}
