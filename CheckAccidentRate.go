package main

import (
	"encoding/json"
	"fmt"
	"nirs/packages/jsonprocess"
	"time"
)

const (
	DTPfile      = "DTP.json"
	DISTRICTfile = "District.json"
	TimeFormat   = "01.02.2006" // "MM/DD/YYYY"
)

func RegionValidation(region string) string {
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
	return regionName
}

func CheckAccidentRate(region string, district string) (string, int) {
	json.Unmarshal(jsonprocess.OpenJSON(ConfigPath), &config)
	DataPathName := config.DataPathName + "/"
	var regionName string = RegionValidation(region)

	currentDate := time.Now().Format(TimeFormat)
	if currentDate != config.LastUpdateDate {
		CSVfilenames := SearchFiles(DataPathName)
		for reg, paths := range CSVfilenames {
			fmt.Println(reg, regionName)
			MergingFiles(paths)
			jsonprocess.ParseJSON(DataPathName+regionName+"/"+regionName+DTPfile, accidents)
			CalculateAcidentRate(accidents)
			jsonprocess.ParseJSON(DataPathName+regionName+"/"+regionName+DISTRICTfile, districts)
		}
		config.LastUpdateDate = currentDate
		jsonprocess.ParseJSON(ConfigPath, config)
	} else {
		if !CheckExist(DataPathName + regionName + "/" + regionName + DISTRICTfile) {
			json.Unmarshal(jsonprocess.OpenJSON(DataPathName+regionName+"/"+regionName+DISTRICTfile), &districts)
		} else {
			districts = nil
		}
	}

	return Calculation(district, accidents)
}
