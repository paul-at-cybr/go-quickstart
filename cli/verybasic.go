package main

import (
	"flag"
	"fmt"

	sh "github.com/paul-at-start/simple-help"
)

const (
	programTitle       = "Very Basic CLI utility"
	programDescription = `A CLI utility based on the golang-default flag package,
  with improved human-readable help output (via paul-at-start/simple-help)`
)

var (
	// Flags
	count       int64  // how many sheep to count
	favCarBrand string // your favorite car brand
)

func init() {
	flag.Int64Var(&count, "count", 0, "Number of sheep to count")
	flag.StringVar(&favCarBrand, "string", "Honda", "Your favorite car brand")
}

func main() {
	// override built-in --help with our own version
	help := &sh.SimpleHelp{
		ProgramTitle:       programTitle,
		ProgramDescription: programDescription,
		Indentation:        15,
	}

	flag.CommandLine.Usage = help.Help

	flag.Parse()

	if count <= 0 { // must count at least 1 sheep
		help.Help() // output help
		return
	}

	// Do the counting!
	fmt.Print("Starting sheep count:\n\n")
	for i := int64(0); i < count; i++ {
		fmt.Printf("%d sheep in a %s\n", (i + 1), favCarBrand)
	}

	fmt.Print("\nDone\n\n")
}
