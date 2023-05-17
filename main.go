package main

import (
	"encoding/json"
	"fmt"
	"nirs/packages/jsonprocess"
	"time"
)

func main() {
	json.Unmarshal(jsonprocess.OpenJSON("./config.json"), &config)

	JSONfilename := config.DataPathName + "/" + config.JSONfilename
	DataPathName := config.DataPathName + "/"

	CSVfilenames := file_processing(JSONfilename, DataPathName)
	MergingFiles(CSVfilenames)
	jsonprocess.ParseJSON(JSONfilename, accidents)

	currentDate := time.Now().Format("01.02.2006")
	if currentDate != config.LastUpdateDate {
		CalculateAcidentRate(accidents)
		config.LastUpdateDate = currentDate
		jsonprocess.ParseJSON("./config.json", config)
		jsonprocess.ParseJSON("./data/districts.json", districts)
	} else {
		json.Unmarshal(jsonprocess.OpenJSON("./data/districts.json"), &districts)
	}

	var flag int
	fmt.Printf("1 - вывести список возможных районов\n2 - выбор района для прогнозирования\nФлаг: ")
	fmt.Scan(&flag)

	switch flag {
	case 1:
		DistrictExport(accidents)
	case 2:
		Calculation(accidents)
	case 3:
		for idx, distrct := range districts {
			fmt.Println(idx+1, distrct.Name, distrct.AccidentRate)
		}
	default:
		fmt.Println("Неккоректное значение флага")
	}
}
