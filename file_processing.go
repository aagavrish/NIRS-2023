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

type Accident struct {
	Number       int     `json:"Number"`
	Date         string  `json:"Date"`
	District     string  `json:"District"`
	DTPnumber    string  `json:"DTPnumber"`
	DTPform      string  `json:"DTPform"`
	Road         string  `json:"Road"`
	Address      string  `json:"Address"`
	Deapths      float64 `json:"Deapths"`
	Wounded      float64 `json:"Wounded"`
	CARScount    float64 `json:"CARScount"`
	MEMBERScount float64 `json:"MEMBERScount"`
}

var accidents []Accident

func file_processing() {
	CSVfilenames := []string{
		"./data/2021/January2021.csv", "./data/2021/February2021.csv", "./data/2021/March2021.csv",
		"./data/2021/April2021.csv", "./data/2021/May2021.csv", "./data/2021/June2021.csv",
		"./data/2021/July2021.csv", "./data/2021/August2021.csv", "./data/2021/September2021.csv",
		"./data/2021/October2021.csv", "./data/2021/November2021.csv", "./data/2021/December2021.csv",

		"./data/2022/January2022.csv", "./data/2022/February2022.csv", "./data/2022/March2022.csv",
		"./data/2022/April2022.csv", "./data/2022/May2022.csv", "./data/2022/June2022.csv",
		"./data/2022/July2022.csv", "./data/2022/August2022.csv", "./data/2022/September2022.csv",
		"./data/2022/October2022.csv", "./data/2022/November2022.csv", "./data/2022/December2022.csv"}
	jsonFilename := "./data/data.json"

	merging_files(CSVfilenames)
	parseJSON(jsonFilename)
}
