package main

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

type District struct {
	Name         string
	AccidentRate float64
}

var districts []District
var accidents []Accident
