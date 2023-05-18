package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

func CheckExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return false
	}
	if errors.Is(err, os.ErrNotExist) {
		return true
	}
	return true
}

func MergingFiles(filenames []string) {
	var index int = 0
	for _, filename := range filenames {

		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader := csv.NewReader(file)
		reader.Comma = ';'

		csvData, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}

		for i := range csvData {
			for j := range csvData[i] {
				csvData[i][j] = strings.ReplaceAll(csvData[i][j], ",", ".")
			}
		}

		var accident Accident
		for _, each := range csvData {
			accident.Number = index
			accident.Date = each[1]
			accident.District = each[2]
			if each[3] == "" {
				accident.DTPnumber = "Не известен"
			} else {
				accident.DTPnumber = each[3]
			}
			accident.DTPform = each[4]
			accident.Road = each[5]
			accident.Address = each[8]
			accident.Deapths, _ = strconv.ParseFloat(each[9], 32)
			accident.Wounded, _ = strconv.ParseFloat(each[10], 32)
			accident.CARScount, _ = strconv.ParseFloat(each[11], 32)
			accident.MEMBERScount, _ = strconv.ParseFloat(each[12], 32)
			accidents = append(accidents, accident)
			index++
		}
	}
}

func SearchFiles(JSONfilename string, DataPathname string) []string {
	var CSVfilenames []string

	yearPaths, err := os.ReadDir(DataPathname)
	if err != nil {
		panic(err)
	}

	for _, yearPath := range yearPaths {
		monthDataFiles, err := os.ReadDir(DataPathname + yearPath.Name())
		if err != nil {
			break
		}
		for _, monthDataFile := range monthDataFiles {
			CSVfilenames = append(CSVfilenames, DataPathname+yearPath.Name()+"/"+monthDataFile.Name())
		}
	}

	return CSVfilenames
}
