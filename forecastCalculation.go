package main

import (
	"strings"
)

const (
	NF = "Не найден"
)

func DistrictsAdd(accidents []Accident) {
	var district District
	var district_names []string = DistrictExport(accidents)
	for _, curname := range district_names {
		district.Name = curname
		district.AccidentRate = 0.0
		districts = append(districts, district)
	}
}

func CalculateAcidentRate(accidents []Accident) {
	DistrictsAdd(accidents)
	var carsCount float64 = 0
	for i := 0; i < len(districts); i++ {
		for j := 0; j < len(accidents); j++ {
			if districts[i].Name == accidents[j].District {
				carsCount += accidents[j].CARScount
			}
		}
		districts[i].AccidentRate = float64(int(carsCount/config.Days*100*100)) / 100
		carsCount = 0
	}
}

func DistrictExport(accidents []Accident) []string {
	var district_names []string
	district_names = append(district_names, accidents[0].District)

	for index := range accidents {
		if FindUnique(accidents[index].District, district_names) {
			district_names = append(district_names, accidents[index].District)
		}
	}

	return district_names
}

func FindUnique(curdistrict string, alldistricts []string) bool {
	for _, districtname := range alldistricts {
		if districtname == curdistrict {
			return false
		}
	}
	return true
}

func StringComparison(defdistrict string, inputdistrict string) bool {
	symbcount := 0.0
	var maxlen int

	defdistrict = strings.ToLower(defdistrict)
	inputdistrict = strings.ToLower(inputdistrict)

	if len(defdistrict) > len(inputdistrict) {
		maxlen = len(inputdistrict)
	} else {
		maxlen = len(defdistrict)
	}

	for i := 0; i < maxlen; i++ {
		if defdistrict[i] == inputdistrict[i] {
			symbcount++
		}
	}

	if float64(symbcount)/float64(maxlen) >= config.InputAccuracy {
		return true
	} else {
		return false
	}
}

func Calculation(district string, accident []Accident) (string, int) {
	var averageRate float64
	for _, district := range districts {
		averageRate += district.AccidentRate
	}
	averageRate /= float64(len(districts))

	var districtName string = NF
	var districtRate int = 0

	for i := 0; i < len(districts); i++ {
		if StringComparison(districts[i].Name, district) {
			districtName = districts[i].Name
			districtRate = int(districts[i].AccidentRate)
			districts, accidents = nil, nil
		}
	}
	return districtName, districtRate
}
