package main

import (
	"regexp"
	"os"
)

const (
	logFileDefault = "log-sample.txt"
	numberOfLines  = 20000000 //20 million
	numberWorkers  = 64
)

var regex = regexp.MustCompile(".*#[15][15]1110+$")

func main() {
	logFile := os.Args[1]

	//createLogFile(logFile)
	//sequential(logFile)
	//parallel(logFile)
	scanner(logFile)
}
