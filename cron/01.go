package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

type Job struct {
}

func (this *Job) Run() {
	fmt.Println("job exec...")
}

func main() {
	c := cron.New(
		cron.WithLogger(
			cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))), cron.WithSeconds())
	c.AddJob("0 * * * * *", &Job{})

	c.Start()

	time.Sleep(1000 * time.Second)
}
