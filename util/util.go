package util

import (
	"fmt"
	"golang.org/x/net/idna"
	"sync"
	"time"
)

var (
	Counter         int
	Wg              sync.WaitGroup
	Buffer          = make(map[int]string)
	buffWriterQueue = make(chan buffWriter, 100000)
)

const (
	PROFILING_BATCH_SIZE = 100000
)

type Job struct {
	text    string
	line    int
	conv    string
	counter *int
	sTime   time.Time
}
type buffWriter struct {
	text string
	line int
}

type Worker struct {
	WorkerId int
}

func NewWorker(id int) Worker {
	worker := Worker{WorkerId: id}
	return worker
}

func (w *Worker) Start(jobQueue <-chan Job) {
	var (
		wErr error
		text string
	)
	go func() {
		for curJob := range jobQueue {
			switch curJob.conv {
			case "a":
				text, wErr = idna.ToASCII(curJob.text)
				buffWriterQueue <- buffWriter{text: text, line: curJob.line}
			case "u":
				text, wErr = idna.ToUnicode(curJob.text)
				buffWriterQueue <- buffWriter{text: text, line: curJob.line}
			}
			if wErr != nil {
				fmt.Println(w.WorkerId, wErr.Error())
			}
			*curJob.counter++
			if (*curJob.counter % PROFILING_BATCH_SIZE) == 0 {
				fmt.Printf("%d converted in %.6f secs\n", *curJob.counter, time.Since(curJob.sTime).Seconds())
			}
		}
	}()
}

//write text to buffer
func TextWriter() {
	for {
		select {
		case curWriter := <-buffWriterQueue:
			Buffer[curWriter.line] = curWriter.text
			Wg.Done()
		}
	}
}
