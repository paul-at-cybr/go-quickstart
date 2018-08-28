package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// config - defines the data that we expect to get from the toml config
type config struct {
	ICBMLaunchCode string
	DefaultTargets []string
}

func main() {
	// getConfig() does the config getting
	conf, err := getConfig("./icbm-test.toml")

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	// success!
	fmt.Println("\nGot config!")
	fmt.Println("Launch code: " + conf.ICBMLaunchCode)

	fmt.Println("Targets:")
	for _, target := range conf.DefaultTargets {
		fmt.Println("  " + target)
	}
}

// getConfig - reads config file and ships it to toml.Decode
// then returns the result
func getConfig(path string) (conf *config, err error) {
	conf = &config{}

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, fmt.Errorf("Config file not found at the specified path (%s)", path)
	}

	_, err = toml.Decode(string(bs), conf)
	if err != nil {
		err = fmt.Errorf("Could not decode config file: %s", err.Error())
	}
	return conf, err
}
