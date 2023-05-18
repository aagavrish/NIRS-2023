package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nirs/packages/jsonprocess"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "main")
	json.Unmarshal(jsonprocess.OpenJSON("./config.json"), &config)

	JSONfilename := config.DataPathName + "/" + config.JSONfilename
	DataPathName := config.DataPathName + "/"

	currentDate := time.Now().Format("01.02.2006")
	if currentDate != config.LastUpdateDate || CheckExist(JSONfilename) {
		CSVfilenames := SearchFiles(JSONfilename, DataPathName)
		MergingFiles(CSVfilenames)
		jsonprocess.ParseJSON(JSONfilename, accidents)
		CalculateAcidentRate(accidents)

		config.LastUpdateDate = currentDate
		jsonprocess.ParseJSON("./config.json", config)
		jsonprocess.ParseJSON("./data/districts.json", districts)
	} else {
		json.Unmarshal(jsonprocess.OpenJSON(JSONfilename), &accidents)
		json.Unmarshal(jsonprocess.OpenJSON("./data/districts.json"), &districts)
	}

	var flag int
	fmt.Printf("1 - вывести список возможных районов\n2 - выбор района для прогнозирования\nФлаг: ")
	fmt.Scan(&flag)

	switch flag {
	case 1:
		for idx, distrct := range districts {
			fmt.Printf("%d. %s\n", idx+1, distrct.Name)
		}
	case 2:
		Calculation(accidents)
	default:
		fmt.Println("Неккоректное значение флага")
	}
}
