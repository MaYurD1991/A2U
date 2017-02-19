package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mayur-tolexo/A2U/util"
	"golang.org/x/net/idna"
	"log"
	"os"
)

func main() {
	input, output, conv, nWorker := getFlags()
	flag.Parse()
	var (
		fp  *os.File
		err error
	)

	if *output == "" {
		fmt.Println("Output can't be blank")
		return
	}
	if fp, err = os.Open(*input); err != nil {
		fmt.Println("Input file error", err.Error())
	}
	jobQueue := make(chan Job, 100)
	for i := 0; i < *nWorker; i++ {
		worker := util.NewWorker(i + 1)
		worker.Start(jobQueue)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		curTxt := scanner.Text()
	}
}
func getFlags() (input, output, conv *string, nWorker *int) {
	input = flag.String("i", "", "input file")
	output = flag.String("o", "", "output file")
	conv = flag.String("c", "a", "converter a for ASCII or u for UNICODE")
	nWorker = flag.Int("w", "5", "no of workers")
	return
}
