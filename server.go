package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Version string `json:"version"`
}

func loadConfig(filename string) Config {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
func main() {
	config := loadConfig("config.json")
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runFile := runCmd.String("file", "", "file")

	if len(os.Args) < 2 {
		fmt.Println("expected help, version or run command")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
	case "version":
		fmt.Println("go_virtus ver. " + config.Version)
	case "help":
		fmt.Println("")
	}
}
