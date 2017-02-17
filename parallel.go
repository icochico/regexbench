package main

import (
	"sync"
	"time"
	"log"
	"os"
	"bufio"
	"fmt"
)

var waitGroup sync.WaitGroup

func parallel(logFile string) {
	start := time.Now()
	log.Printf("Started: %v", start)
	inputFile, err := os.Open(logFile)
	if err != nil {
		log.Println("Unable to open file: ", logFile)
		os.Exit(-1)
	}

	defer inputFile.Close()

	waitGroup.Add(numberWorkers)

	queue := make(chan []byte)
	for gr := 1; gr <= numberWorkers; gr++ {
		go worker(queue, gr)
	}

	reader := bufio.NewReader(inputFile)
	for line, _, err := reader.ReadLine(); err == nil;
	line, _, err = reader.ReadLine() {
		queue <- line
	}

	close(queue)
	waitGroup.Wait()
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}

func worker(queue chan []byte, id int) {
	defer waitGroup.Done()
	for {
		line, ok := <-queue
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", id)
			return
		}
		if regex.Match(line) {
			fmt.Printf("[%v] Match: %v\n", id, string(line))
		}
	}
}
