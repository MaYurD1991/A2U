package util

import (
	"fmt"
	"golang.org/x/net/idna"
	"sync"
	"time"
)

var (
	Wg     sync.WaitGroup
	Buffer = make(map[int]string)
)

const (
	PROFILING_BATCH_SIZE = 100
)

type Job struct {
	text    string
	line    int
	conv    string
	counter *int
	sTime   time.Time
}

type Worker struct {
	WorkerId int
}

func NewWorker(id int) Worker {
	worker := Worker{WorkerId: id}
	return worker
}

func (w *Worker) Start(jobQueue <-chan Job) {
	var wErr error
	go func() {
		for curJob := range jobQueue {
			switch curJob.conv {
			case "a":
				Buffer[curJob.line], wErr = idna.ToASCII(curJob.text)
			case "u":
				Buffer[curJob.line], wErr = idna.ToUnicode(curJob.text)
			}
			if wErr != nil {
				fmt.Println(w.WorkerId, wErr.Error())
			}
			*curJob.counter++
			if (*curJob.counter % PROFILING_BATCH_SIZE) == 0 {
				fmt.Println(*curJob.counter, "converted in ", time.Since(curJob.sTime))
			}
			Wg.Done()
		}
	}()
}
