package structures

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

type Config struct {
	Days           float64 `json:"Days"`
	JSONfilename   string  `json:"JSONfilename"`
	Interval       float64 `json:"Interval"`
	LastUpdateDate string  `json:"LastUpdateDate"`
	DataPathName   string  `json:"DataPathName"`
	InputAccuracy  float64 `json:"InputAccuracy"`
}

type District struct {
	Name         string
	AccidentRate float64
}
