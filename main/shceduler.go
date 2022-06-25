package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func RunAndExecuteJobsMap(tasks func()) {
	scheduler := gocron.NewScheduler(time.Local)
	job, err := scheduler.Every(1).Minute().Do(tasks)
	if err != nil {
		log.Println(err)
	}
	scheduler.StartAsync()
	fmt.Println(job.ScheduledTime())
}
