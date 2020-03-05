package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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

	//fileToWrite, fileToWrteError := openFile("results.csv")

	// if fileToWrteError != nil {
	// 	panic(fileToWrteError)
	// }

	//csvWriter := csv.NewWriter(fileToWrite)

	browserData, browserDataErr := readCSVFiles(filenames)

	if browserDataErr != nil {
		panic(browserDataErr)
	}

	// csvStrings := [][]string{}

	// clientKeys := make(map[ClientModel]int)

	// uniqueUsers := 0
	// totalReturningUsers := 0
	// chromeReturningUsers := 0
	// safariReturningUsers := 0
	// samsungInternetReturningUsers := 0
	// firefoxReturningUsers := 0
	// edgeReturningUsers := 0
	// ieReturningUsers := 0

	//clientModelList := []ClientModel{}

	//var chromeData []BrowserData

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
	chromeUserCount := getReturningUsersForBrowser(groupQuery, "chrome")
	safariUserCount := getReturningUsersForBrowser(groupQuery, "safari")
	ieUserCount := getReturningUsersForBrowser(groupQuery, "internet explorer")
	edgeUserCount := getReturningUsersForBrowser(groupQuery, "edge")
	firefoxUserCount := getReturningUsersForBrowser(groupQuery, "firefox")
	samsungUserCount := getReturningUsersForBrowser(groupQuery, "samsung internet")

	fmt.Println("\nTotal sessions: ", len(browserData))
	fmt.Println("Returning users: ", len(groupQuery))
	fmt.Println("Returning Chrome users: ", chromeUserCount)
	fmt.Println("Returning Safari 12+ users: ", safariUserCount)
	fmt.Println("Returning IE users: ", ieUserCount)
	fmt.Println("Returning Edge users: ", edgeUserCount)
	fmt.Println("Returning Firefox users: ", firefoxUserCount)
	fmt.Println("Returning Samsung Internet users: ", samsungUserCount)

	// for _, qGroup := range query {
	// 	fmt.Println(qGroup.Key)
	// }

	// fmt.Println("Group length: ", len(query))

	// for _, entry := range browserData {
	// 	clientModelList = append(clientModelList, ClientModel{
	// 		Browser:  entry.Browser,
	// 		ClientID: entry.ClientID,
	// 	})
	// }

	// for _, entry := range clientModelList {
	// 	_, exist := clientKeys[entry]

	// 	if exist {
	// 		clientKeys[entry] += 1
	// 	} else {
	// 		clientKeys[entry] = 1
	// 	}
	// }

	// for k, v := range clientKeys {
	// 	if v > 1 {
	// 		totalReturningUsers += 1

	// 		switch strings.ToLower(k.Browser) {
	// 		case "chrome":
	// 			chromeReturningUsers += 1
	// 		case "safari":
	// 			safariReturningUsers += 1
	// 		case "samsung internet":
	// 			samsungInternetReturningUsers += 1
	// 		case "firefox":
	// 			firefoxReturningUsers += 1
	// 		case "edge":
	// 			edgeReturningUsers += 1
	// 		case "internet explorer":
	// 			ieReturningUsers += 1
	// 		}
	// 	}
	// }

	// uniqueUsers = len(clientModelList) - totalReturningUsers

	// fmt.Println("Total unique users: ", uniqueUsers)
	// fmt.Println("Total returning users: ", totalReturningUsers)

	// fmt.Println("Total Chrome returning users: ", chromeReturningUsers)
	// fmt.Println("Total Safari 12+ returning users: ", safariReturningUsers)
	// fmt.Println("Total Samsung Internet returning users: ", samsungInternetReturningUsers)
	// fmt.Println("Total Firefox returning users: ", firefoxReturningUsers)
	// fmt.Println("Total Edge returning users: ", edgeReturningUsers)
	// fmt.Println("Total Internet Explorer returning users: ", ieReturningUsers)

	// totalUsersModel := CSVRow{
	// 	Title: "Total unique users",
	// 	Count: strconv.Itoa(uniqueUsers),
	// }
	// returningUsersModel := CSVRow{
	// 	Title: "Total returning users",
	// 	Count: strconv.Itoa(totalReturningUsers),
	// }
	// chromeUsersModel := CSVRow{
	// 	Title: "Total Chrome returning users",
	// 	Count: strconv.Itoa(chromeReturningUsers),
	// }
	// safariUsersModel := CSVRow{
	// 	Title: "Total Safari 12+ returning users",
	// 	Count: strconv.Itoa(safariReturningUsers),
	// }
	// samsungUsersModel := CSVRow{
	// 	Title: "Total Samsung Internet returning user",
	// 	Count: strconv.Itoa(samsungInternetReturningUsers),
	// }
	// firefoxUsersModel := CSVRow{
	// 	Title: "Total Firefox returning users",
	// 	Count: strconv.Itoa(firefoxReturningUsers),
	// }
	// edgeUsersModel := CSVRow{
	// 	Title: "Total Edge returning users",
	// 	Count: strconv.Itoa(edgeReturningUsers),
	// }
	// ieUsersModel := CSVRow{
	// 	Title: "Total Internet Explorer returning users",
	// 	Count: strconv.Itoa(ieReturningUsers),
	// }

	// csvStrings = append(csvStrings, []string{
	// 	totalUsersModel.Title,
	// 	totalUsersModel.Count,
	// })
	// csvStrings = append(csvStrings, []string{
	// 	returningUsersModel.Title,
	// 	returningUsersModel.Count,
	// })
	// csvStrings = append(csvStrings, []string{
	// 	chromeUsersModel.Title,
	// 	chromeUsersModel.Count,
	// })
	// csvStrings = append(csvStrings, []string{
	// 	safariUsersModel.Title,
	// 	safariUsersModel.Count,
	// })
	// csvStrings = append(csvStrings, []string{
	// 	samsungUsersModel.Title,
	// 	samsungUsersModel.Count,
	// })
	// csvStrings = append(csvStrings, []string{
	// 	firefoxUsersModel.Title,
	// 	firefoxUsersModel.Count,
	// })
	// csvStrings = append(csvStrings, []string{
	// 	edgeUsersModel.Title,
	// 	edgeUsersModel.Count,
	// })
	// csvStrings = append(csvStrings, []string{
	// 	ieUsersModel.Title,
	// 	ieUsersModel.Count,
	// })

	// for _, csvLine := range csvStrings {
	// 	csvWriter.Write(csvLine)
	// }

	// csvWriter.Flush()
	// fileToWrite.Close()

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
