package main

import (
	"encoding/json"
	"fmt"
	"nirs/packages/jsonprocess"
	"time"
)

func CheckAccidentRate(region string, district string) (string, int) {
	json.Unmarshal(jsonprocess.OpenJSON("./config.json"), &config)
	DataPathName := config.DataPathName + "/"
	var regionName string
	switch region {
	case "Москва":
		regionName = "Moscow"
	case "Московская область":
		regionName = "MoscowRegion"
	case "Санкт-Петербург":
		regionName = "SaintPetersburg"
	case "Казань":
		regionName = "Kazan"
	default:
		regionName = "Moscow"
		data.Region = "Москва"
	}

	currentDate := time.Now().Format("01.02.2006")
	if currentDate != config.LastUpdateDate {
		CSVfilenames := SearchFiles(DataPathName)
		for reg, paths := range CSVfilenames {
			fmt.Println(reg, regionName)
			MergingFiles(paths)
			jsonprocess.ParseJSON(DataPathName+regionName+"/"+regionName+"DTP.json", accidents)
			CalculateAcidentRate(accidents)
			jsonprocess.ParseJSON(DataPathName+regionName+"/"+regionName+"District.json", districts)
		}
		config.LastUpdateDate = currentDate
		jsonprocess.ParseJSON("./config.json", config)
	} else {
		if !CheckExist(DataPathName + regionName + "/" + regionName + "District.json") {
			json.Unmarshal(jsonprocess.OpenJSON(DataPathName+regionName+"/"+regionName+"District.json"), &districts)
		} else {
			districts = nil
		}
	}

	return Calculation(district, accidents)
}
