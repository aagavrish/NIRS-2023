package forecast

import (
	"encoding/json"
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
	case "Севастополь":
		regionName = "Sevastopol"
	case "Санкт-Петербург":
		regionName = "SaintPetersburg"
	default:
		regionName = "Moscow"
	}
	return regionName
}

func CheckAccidentRate(region string, district string) (string, int, int) {
	json.Unmarshal(jsonprocess.OpenJSON(ConfigPath), &config)
	DataPathName := config.DataPathName + "/"
	var regionName string = RegionValidation(region)

	currentDate := time.Now().Format(TimeFormat)
	if currentDate != config.LastUpdateDate || CheckExist(DataPathName+regionName+"/"+regionName+DISTRICTfile) {
		CSVfilenames := SearchFiles(DataPathName)
		for reg, paths := range CSVfilenames {
			if regionName != reg {
				continue
			}
			districts, accidents = nil, nil

			MergingFiles(paths)
			jsonprocess.ParseJSON(DataPathName+reg+"/"+reg+DTPfile, accidents)
			CalculateAcidentRate(accidents)
			jsonprocess.ParseJSON(DataPathName+reg+"/"+reg+DISTRICTfile, districts)
		}
		config.LastUpdateDate = currentDate
		jsonprocess.ParseJSON(ConfigPath, config)
	} else {
		districts, accidents = nil, nil
		json.Unmarshal(jsonprocess.OpenJSON(DataPathName+regionName+"/"+regionName+DISTRICTfile), &districts)
	}

	return Calculation(district, accidents)
}
