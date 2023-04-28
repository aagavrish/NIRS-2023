package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const days = 365 * 2

func districts_add_name(accidents []Accident) {
	var district District
	var district_names []string = district_export(accidents, 1)
	for _, curname := range district_names {
		district.Name = curname
		district.AccidentRate = 0.0
		districts = append(districts, district)
	}
}

func districts_calculate_accidentrate(accidents []Accident, flag int) float64 {
	districts_add_name(accidents)
	var carsCount float64 = 0
	var averageRate float64 = 0
	for i := 0; i < len(districts); i++ {
		for j := 0; j < len(accidents); j++ {
			if districts[i].Name == accidents[j].District {
				carsCount += accidents[j].CARScount
			}
		}
		districts[i].AccidentRate = float64(int(carsCount/days*100*100)) / 100
		if flag == 0 {
			fmt.Printf("%d. %s: %.2f\n", i, districts[i].Name, districts[i].AccidentRate)
		}
		averageRate += districts[i].AccidentRate
		carsCount = 0
	}
	averageRate /= float64(len(districts))
	return averageRate
}

func district_export(accidents []Accident, flag int) []string {
	var district_names []string
	district_names = append(district_names, accidents[0].District)

	for index := range accidents {
		if unique_district(accidents[index].District, district_names) {
			district_names = append(district_names, accidents[index].District)
		}
	}

	if flag == 0 {
		for index, districtname := range district_names {
			fmt.Printf("%d. %s\n", index, districtname)
		}
	}

	return district_names
}

func unique_district(curdistrict string, alldistricts []string) bool {
	for _, districtname := range alldistricts {
		if districtname == curdistrict {
			return false
		}
	}
	return true
}

func calculation(accidents []Accident) {
	var averageRate float64 = districts_calculate_accidentrate(accidents, 1)
	var count float64
	var district string

	fmt.Printf("Район: ")
	fmt.Scan(&district)
	for i := 0; i < len(districts); i++ {
		if districts[i].Name == district {
			for _, accident := range accidents {
				if accident.District == district {
					fmt.Printf("Дата: %s. Номер ДТП: %s. Форма ДТП: %s.\tКоличество авто: %.0f\n",
						accident.Date, accident.DTPnumber, accident.DTPform, accident.CARScount)
					count += accident.CARScount
				}
			}
			fmt.Printf("Итого %d автомобилей попали в ДТП\n\n", int(count))
			if districts[i].AccidentRate <= averageRate+10 && districts[i].AccidentRate >= averageRate-10 {
				fmt.Printf("В районе '%s' СРЕДНЕЕ значение аварийности: %.2f%% (среднее значение: %.2f%%)\n", districts[i].Name, districts[i].AccidentRate, averageRate)
			} else if districts[i].AccidentRate > averageRate+10 {
				fmt.Printf("В районе '%s' ВЫСОКОЕ значение аварийности: %.2f%% (среднее значение: %.2f%%)\n", districts[i].Name, districts[i].AccidentRate, averageRate)
			} else if districts[i].AccidentRate < averageRate-10 {
				fmt.Printf("В районе '%s' НИЗКОЕ значение аварийности: %.2f%% (среднее значение: %.2f%%)\n", districts[i].Name, districts[i].AccidentRate, averageRate)
			}
		}
	}
}

func main() {
	file_processing()
	JSONfile, err := os.Open("./data/data.json")
	if err != nil {
		panic(err)
	}
	defer JSONfile.Close()

	byteResult, _ := io.ReadAll(JSONfile)
	var accidents []Accident
	json.Unmarshal(byteResult, &accidents)

	var flag int
	fmt.Printf("1 - вывести список возможных районов\n2 - выбор района для прогнозирования\n3 - обновить данные\nФлаг: ")
	fmt.Scan(&flag)

	switch flag {
	case 1:
		district_export(accidents, 0)
	case 2:
		calculation(accidents)
	case 3:
		file_processing()
	case 4:
		districts_calculate_accidentrate(accidents, 0)
	default:
		fmt.Println("Неккоректное значение флага")
	}
}
