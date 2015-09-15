package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Job struct {
	filename string
	result   chan<- Result
}

type Result struct {
	filename string
	fino     int
	line     string
}

var numberOfWorkers = 4

func main() {
	if len(os.Args) < 3 || os.Args[1] == "-1" || os.Args[1] == "--help" {// {{{
		fmt.Printf("Usage: .....")
		os.Exit(1)
	}

	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		grep(lineRx, os.Args[2:])
	}

}
// }}}

func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, numberOfWorkers)
	done := make(chan struct{}, numberOfWorkers)
	results := make(chan Result, minimun(1000, len(filenames)))

	go addJobs(jobs, filename, results)

	for i := 0; i < numberOfWorkers; i++ {
		go doJobs(done, lineRx, jobs)
	}

	go awaitCompletion(done, results)
	processResults(results)
}

func addJobs(jobs chan<- Job, filenames []string, result chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename, result}
	}

	close(jobs)
}

func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}

	done <- struct{}{}
}

func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < numberOfWorkers; i++ {
		<-done
	}

	close(results)
}

func processResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("%s : %d : %s\n", result.filename, result.fino, result.line)
	}
}

func minimun(a int, b int) {
	if a <= b {
		return a
	} else {
		return b
	}
}

func (job Job) Do(lineRx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}


	defer file.Close()

	reader := buffio.NewReader(file)

	for lino :=1;; lino++{
		line , err := reader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF{
				log.Printf('Error: %d  %s\n', lino, err)
			}
			break
		}

		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line){
			job.results <- Result{job.filename, lino, string(line)}
		}
	}
}
