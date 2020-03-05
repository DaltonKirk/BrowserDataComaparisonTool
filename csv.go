package main

import (
	"encoding/csv"
	"os"
)

func readCSVFiles(filenames []string) ([]BrowserData, error) {
	browserDataList := []BrowserData{}

	for _, filename := range filenames {
		csvLines, csvLinesErr := getCSVLines(filename)

		if csvLinesErr != nil {
			return []BrowserData{}, csvLinesErr
		}

		for _, csvLine := range csvLines {
			browserData := BrowserData{
				ID:                 csvLine[0],
				Date:               csvLine[1],
				ClientID:           csvLine[2],
				DeviceCategory:     csvLine[3],
				Browser:            csvLine[4],
				BrowserVersion:     csvLine[5],
				Sessions:           csvLine[6],
				Transactions:       csvLine[7],
				TransactionRevenue: csvLine[8],
			}

			browserDataList = append(browserDataList, browserData)
		}
	}

	return browserDataList, nil
}

func getCSVLines(filename string) ([][]string, error) {
	file, fileErr := os.Open(filename)

	if fileErr != nil {
		return [][]string{}, fileErr
	}

	defer file.Close()

	lines, linesErr := csv.NewReader(file).ReadAll()

	if linesErr != nil {
		return [][]string{}, linesErr
	}

	return lines, nil
}
