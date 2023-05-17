package main

import (
	"fmt"
	"strings"
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

func Calculation(accidents []Accident) {
	var averageRate float64
	for _, district := range districts {
		averageRate += district.AccidentRate
	}
	averageRate /= float64(len(districts))

	var found bool
	var district string
	fmt.Printf("Район: ")
	fmt.Scan(&district)
	for i := 0; i < len(districts); i++ {
		if StringComparison(districts[i].Name, district) {
			district = districts[i].Name
			found = true
			if districts[i].AccidentRate <= averageRate+config.Interval && districts[i].AccidentRate >= averageRate-config.Interval {
				fmt.Printf("В районе '%s' СРЕДНЕЕ значение аварийности: %.2f%% (среднее значение: %.2f%%)\n", districts[i].Name, districts[i].AccidentRate, averageRate)
			} else if districts[i].AccidentRate > averageRate+config.Interval {
				fmt.Printf("В районе '%s' ВЫСОКОЕ значение аварийности: %.2f%% (среднее значение: %.2f%%)\n", districts[i].Name, districts[i].AccidentRate, averageRate)
			} else if districts[i].AccidentRate < averageRate-config.Interval {
				fmt.Printf("В районе '%s' НИЗКОЕ значение аварийности: %.2f%% (среднее значение: %.2f%%)\n", districts[i].Name, districts[i].AccidentRate, averageRate)
			}
		}
	}

	if !found {
		return
	}

	fmt.Printf("Предоставить отчёт по ДТП в районе %s? Да/нет: ", district)
	var answer string
	fmt.Scan(&answer)
	switch answer {
	case "Да", "да", "Yes", "yes", "y":
		var carsCount int = 0
		for _, accident := range accidents {
			if accident.District == district {
				fmt.Printf("Район: %s. Дата: %s. Номер ДТП: %s. Форма ДТП: %s. Количество авто: %.0f.\n",
					accident.District, accident.Date, accident.DTPnumber, accident.DTPform, accident.CARScount)
				carsCount += int(accident.CARScount)
			}
		}
		fmt.Printf("Итого в ДТП попало: %d\n", carsCount)
	}
}
