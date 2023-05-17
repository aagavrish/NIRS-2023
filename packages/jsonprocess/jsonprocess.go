package jsonprocess

import (
	"encoding/json"
	"io"
	"os"
)

func ParseJSON(outFilename string, dataset interface{}) {
	jsonData, err := json.MarshalIndent(dataset, "", " ")
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create(outFilename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}

func OpenJSON(inFilename string) []byte {
	JSONfile, err := os.Open(inFilename)
	if err != nil {
		panic(err)
	}
	defer JSONfile.Close()

	byteResult, _ := io.ReadAll(JSONfile)
	return byteResult
}
