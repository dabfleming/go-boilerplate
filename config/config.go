package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
)

type Configuration struct {
	Env         string
	Environment string
	Home        string

	Debug    bool `toml:"debug_mode"`
	Database map[string]interface{}
}

var config *Configuration = nil

func init() {
	config = new(Configuration)

	// TODO More logical lookup of file
	fileLocation := os.Getenv("GOPATH") + "/src/github.com/dabfleming/go-boilerplate/.config.toml"

	log.Print("Loading configuration from: ", fileLocation)

	fileContents, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	_, err = toml.Decode(string(fileContents), config)
	if err != nil {
		debug.PrintStack()
		log.Fatalf("Error decoding config: %v", err)
	}

	log.Printf("Configuration loaded.\n\tEnv: %v\n\tEnvironment: %v", config.Env, config.Environment)
}

func Get() *Configuration {
	return config
}

func IsDebug() bool {
	return config.Debug
}
