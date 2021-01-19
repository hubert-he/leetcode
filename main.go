package main

import (
	"./runner"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const timeout = 10 * time.Second

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./ex <url>")
		os.Exit(-1)
	}
}

func main() {
	r,err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

	log.Println("Starting work")
	j := runner.New(timeout)
	j.Add(createTask(), createTask(), createTask())
	if err := j.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		case runner.ErrTimeout:
			log.Println("Terminating due to Timeout.")
			os.Exit(1)
		}
	}
	log.Println("job ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}

