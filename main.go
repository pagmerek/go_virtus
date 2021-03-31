package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Version string `json:"version"`
}

func loadConfig() (Config, error) {
	var config Config
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
func startServer(filename string) error {
	config, _ := loadConfig()
	_, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
	log.Fatal(http.ListenAndServe(config.Port, nil))
	return nil
}
func main() {
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runFile := runCmd.String("file", "", "file")
	if len(os.Args) < 2 {
		fmt.Println("expected help, version or run command")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
		startServer(*runFile)

	case "version":
		config, _ := loadConfig()
		fmt.Println("go_virtus ver. " + config.Version)
	case "help":
		fmt.Println("Simple file server for VirtusLab assessment")
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Println("	./<binary name> <command> [argument]")
		fmt.Println("The commands are: ")
		fmt.Println("	version                       show current version of http server")
		fmt.Println("	help                          show this manual")
		fmt.Println("	run --file <filename>         starts listening on port set in the config.json with file <filename> served under host adress")

	}
}
