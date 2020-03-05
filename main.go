package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Start time: ", start.String())

	filenames, filenamesErr := getFilenames("\\data")

	if filenamesErr != nil {
		panic(filenamesErr)
	}

	fileToWrite, fileToWrteError := openFile("results.csv")

	if fileToWrteError != nil {
		panic(fileToWrteError)
	}

	csvWriter := csv.NewWriter(fileToWrite)

	browserData, browserDataErr := readCSVFiles(filenames)

	if browserDataErr != nil {
		panic(browserDataErr)
	}

	csvStrings := [][]string{}

	for _, bd := range browserData {
		csvStrings = append(csvStrings, []string{
			bd.ID,
			bd.Date,
			bd.ClientID,
			bd.DeviceCategory,
			bd.Browser,
			bd.BrowserVersion,
			bd.Sessions,
			bd.Transactions,
			bd.TransactionRevenue,
		})
	}

	for _, csvLine := range csvStrings {
		csvWriter.Write(csvLine)
	}

	csvWriter.Flush()
	fileToWrite.Close()

	end := time.Since(start)

	fmt.Println("Execution time: ", end.String())
}

func getWorkingDir() (string, error) {
	dir, dirErr := os.Getwd()

	if dirErr != nil {
		return "", dirErr
	}

	return dir, nil
}

func getFilenames(folder string) ([]string, error) {
	dir, dirErr := getWorkingDir()

	if dirErr != nil {
		return []string{}, dirErr
	}

	filenames := []string{}
	rootPath := path.Join(dir, folder)

	pathWalkError := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() || filepath.Ext(path) != ".csv" {
				return nil
			}

			filenames = append(filenames, path)

			return nil
		})

	if pathWalkError != nil {
		return []string{}, pathWalkError
	}

	return filenames, nil
}

func openFile(filename string) (*os.File, error) {
	dir, dirErr := getWorkingDir()

	if dirErr != nil {
		return &os.File{}, dirErr
	}

	fileToWrite, fileToWriteError := os.Create(path.Join(dir, filename))

	if fileToWriteError != nil {
		return &os.File{}, fileToWriteError
	}

	return fileToWrite, nil
}
