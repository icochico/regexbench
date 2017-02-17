package main

import (
	"log"
	"os"
	"time"
	"fmt"
)

func createLogFile(logFile string) {
	log.Println("Creating log file: " + logFile)
	fo, err := os.Create(logFile)
	if err != nil {
		log.Println("Unable to create file: ", logFile)
		os.Exit(-1)
	}
	now := time.Now()
	for i := 0; i < numberOfLines; i++ {
		date := now.Add(time.Duration(i) * time.Second)
		_, err = fo.WriteString(fmt.Sprintf("%v -- %v\n", date, fmt.Sprintf("log entry #%v", i)))
		if i%100000 == 0 {
			log.Printf("... created %v lines", i)
		}
	}
	log.Printf("Done writing %v lines", logFile)
}
