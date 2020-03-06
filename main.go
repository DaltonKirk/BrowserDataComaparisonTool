package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	. "github.com/ahmetb/go-linq"
)

func main() {
	start := time.Now()

	fmt.Println("Start time: ", start.String())

	filenames, filenamesErr := getFilenames("\\data")

	if filenamesErr != nil {
		panic(filenamesErr)
	}

	browserData, browserDataErr := readCSVFiles(filenames)

	if browserDataErr != nil {
		panic(browserDataErr)
	}

	var groupQuery []Group

	// Group browser data by Client Ids
	From(browserData).GroupBy(func(bd interface{}) interface{} {
		return bd.(BrowserData).ClientID
	}, func(bd interface{}) interface{} {
		return bd.(BrowserData)
	}).Where(func(group interface{}) bool {
		convertedGroup := group.(Group)

		if len(convertedGroup.Group) > 1 {
			return true
		}

		return false
	}).ToSlice(&groupQuery)

	// Get User data
	chromeUserCount := []string{"Chrome", intToString(getReturningUsersForBrowser(groupQuery, "chrome"))}
	safariUserCount := []string{"Safari 12+", intToString(getReturningUsersForBrowser(groupQuery, "safari"))}
	ieUserCount := []string{"Internet Explorer", intToString(getReturningUsersForBrowser(groupQuery, "internet explorer"))}
	edgeUserCount := []string{"Edge", intToString(getReturningUsersForBrowser(groupQuery, "edge"))}
	firefoxUserCount := []string{"Firefox", intToString(getReturningUsersForBrowser(groupQuery, "firefox"))}
	samsungUserCount := []string{"Samsung Internet", intToString(getReturningUsersForBrowser(groupQuery, "samsung internet"))}

	userCountList := [][]string{
		chromeUserCount,
		safariUserCount,
		ieUserCount,
		edgeUserCount,
		firefoxUserCount,
		samsungUserCount,
	}

	fmt.Println("\nTotal sessions: ", len(browserData))
	fmt.Println("Returning users: ", len(groupQuery))

	day := intToString(start.Day())
	month := intToString(int(start.Month()))
	year := intToString(start.Year())

	hour := intToString(start.Hour())
	minutes := intToString(start.Minute())
	seconds := intToString(start.Second())

	dateString := day + "-" + month + "-" + year + "-" + hour + "-" + minutes + "-" + seconds

	fileToWrite, fileToWrteError := openFile("results-" + dateString + ".csv")

	if fileToWrteError != nil {
		panic(fileToWrteError)
	}

	defer fileToWrite.Close()

	csvWriter := csv.NewWriter(fileToWrite)

	defer csvWriter.Flush()

	csvWriter.Write([]string{"Total sessions", intToString(len(browserData))})
	csvWriter.Write([]string{"Returning users", intToString(len(groupQuery))})

	for _, userCount := range userCountList {
		fmt.Println("Returning "+userCount[0]+" users: ", userCount[1])

		csvWriter.Write(userCount)
	}

	end := time.Since(start)

	fmt.Println("Execution time: ", end.String())
}

func getReturningUsersForBrowser(groupQuery []Group, browserName string) int {
	var browserUsers []Group

	From(groupQuery).Where(func(group interface{}) bool {
		convertedGroup := group.(Group)

		browser := convertedGroup.Group[1].(BrowserData).Browser

		if strings.ToLower(browser) == browserName {
			return true
		}

		return false
	}).ToSlice(&browserUsers)

	return len(browserUsers)
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

func intToString(x int) (string) {
	return strconv.Itoa(x)
}