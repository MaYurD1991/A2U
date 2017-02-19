package util

import (
	"bufio"
	"log"
	"os"
	"time"
)

func StartConverting(fp *os.File, nWorker int, conv string) {
	sTime := time.Now()
	jobQueue := make(chan Job, 100000)
	for i := 0; i < nWorker; i++ {
		worker := NewWorker(i + 1)
		worker.Start(jobQueue)
	}
	scanner := bufio.NewScanner(fp)
	Counter = 0
	line := 0
	for scanner.Scan() {
		text := scanner.Text()
		select {
		case jobQueue <- Job{text: text, line: line, conv: conv, counter: &Counter, sTime: sTime}:
			Wg.Add(1)
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	close(jobQueue)
}
