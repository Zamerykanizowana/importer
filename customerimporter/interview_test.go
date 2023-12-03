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

func testGeneric(pathInput, pathExpectedOutput, pathReceivedOutput string) error {
	expected, err := mapFromCsv(pathExpectedOutput)
	if err != nil {
		return errors.New("Problem with reading data from expected output")
	}

	received := NewCsvFile(pathInput).ReadFile()
	err = checkCounting(expected, received)
	if err != nil {
		return err
	}

	WriteFile(received, pathReceivedOutput)
	defer os.Remove(pathReceivedOutput)
	received, err = mapFromCsv(pathReceivedOutput)
	if err != nil {
		return err
	}

	err = checkCounting(expected, received)
	if err != nil {
		return err
	}
	return nil
}

func TestSmallAmout(t *testing.T) {
	err := testGeneric("testdata/customers_reordered_columns.csv",
		"testdata/customers_reordered_columns_output.csv",
		"testdata/test_small_amout_received_output.csv")
	if err != nil {
		t.Errorf("%v", err)
	}
}
