package main

import (
	"flag"
	"fmt"
	"github.com/mayur-tolexo/A2U/util"
	"os"
	"time"
)

func main() {
	sTime := time.Now()
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
	go util.TextWriter()
	util.StartConverting(fp, *nWorker, *conv)
	util.Wg.Wait()
	count := len(util.Buffer)
	if ofp, err := os.Create(*output); err == nil {
		fmt.Println("Writing in output file")
		for index := 0; index < count; index++ {
			ofp.WriteString(util.Buffer[index] + "\n")
		}
		ofp.Close()
		fmt.Println(count, "Lines converted in", time.Since(sTime).Seconds(), "secs")
	} else {
		fmt.Println("Write file error", err.Error())
	}
}

//set flag values
func getFlags() (input, output, conv *string, nWorker *int) {
	input = flag.String("i", "", "input file")
	output = flag.String("o", "", "output file")
	conv = flag.String("c", "a", "converter a for ASCII or u for UNICODE")
	nWorker = flag.Int("w", 5, "no of workers")
	return
}
