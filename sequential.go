package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func sequential(logFile string) {
	//count := 0
	start := time.Now()
	log.Printf("Started: %v", start)
	inputFile, err := os.Open(logFile)
	if err != nil {
		log.Println("Unable to open file: ", logFile)
		os.Exit(-1)
	}
	defer inputFile.Close()
	reader := bufio.NewReader(inputFile)
	for line, _, err := reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		if regex.Match(line) {
			fmt.Printf("Match: %v\n", string(line))
		}
	}
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}
