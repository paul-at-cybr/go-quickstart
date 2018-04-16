package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type config struct {
	ICBMLaunchCode string
	DefaultTargets []string
}

func main() {
	conf, err := getConfig("./icbm-test.toml")

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Println("\nGot config!")
	fmt.Println("Launch code: " + conf.ICBMLaunchCode)

	fmt.Println("Targets:")
	for _, target := range conf.DefaultTargets {
		fmt.Println("  " + target)
	}
}

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
