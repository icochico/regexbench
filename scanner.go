package main

import (
	"time"
	"log"
	"os"
	"bufio"
	"fmt"
)

func scanner(logFile string) {
	start := time.Now()
	log.Printf("Started: %v", start)
	inputFile, err := os.Open(logFile)
	if err != nil {
		log.Println("Unable to open file: ", logFile)
		os.Exit(-1)
	}
	scanner := bufio.NewScanner(inputFile)
	i := 0
	for scanner.Scan() {
		_ = scanner.Bytes()
		i++
	}
	fmt.Printf("Count: %v\nElapsed time: %v\n", i, time.Since(start))
}
