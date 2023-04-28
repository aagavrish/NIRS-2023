package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func merging_files(filenames []string) {
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
			accident.DTPnumber = each[3]
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

func parseJSON(filename string) {
	jsonData, err := json.MarshalIndent(accidents, "", " ")
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}

func file_processing() {
	JSONfilename := "./data/data.json"
	DataPathName := "./data/"
	var CSVfilenames []string

	files, err := os.ReadDir(DataPathName)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		curfile, err := os.ReadDir(DataPathName + file.Name())
		if err != nil {
			break
		}
		for _, curcurfile := range curfile {
			CSVfilenames = append(CSVfilenames, DataPathName+file.Name()+"/"+curcurfile.Name())
		}
	}

	merging_files(CSVfilenames)
	parseJSON(JSONfilename)
}
