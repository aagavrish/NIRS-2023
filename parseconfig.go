package main

import (
	"encoding/json"
	"io"
	"os"
)

func uncofig() {
	configJSONfile, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer configJSONfile.Close()

	byteResultConf, _ := io.ReadAll(configJSONfile)
	json.Unmarshal(byteResultConf, &config)
}
