package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func readCSVFiles(filenames []string) ([]BrowserData, error) {
	browserDataList := []BrowserData{}

	for _, filename := range filenames {
		csvLines, csvLinesErr := getCSVLines(filename)

		if csvLinesErr != nil {
			return []BrowserData{}, csvLinesErr
		}

		for i, csvLine := range csvLines {
			if i == 0 {
				continue
			}

			browserData := BrowserData{
				ID:                 i,
				Date:               csvLine[1],
				ClientID:           csvLine[2],
				DeviceCategory:     csvLine[3],
				Sessions:           csvLine[6],
				Transactions:       csvLine[7],
				TransactionRevenue: csvLine[8],
			}

			if strings.ToLower(csvLine[4]) == "safari" {
				if strings.HasPrefix(csvLine[5], "12") || strings.HasPrefix(csvLine[5], "13") {
					browserData.Browser = csvLine[4]
					browserData.BrowserVersion = csvLine[5]
				}
			} else {
				browserData.Browser = csvLine[4]
				browserData.BrowserVersion = csvLine[5]
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
