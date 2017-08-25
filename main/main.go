package main

import (
	"cron"
	"log"
	"time"
)

func main() {
	c := cron.New()
	spec := "*/1 * * * * *"
	i := 0
	j := 0
	k := 0
	_, job1 := c.AddJob(spec, cron.Job{
		Count: -1,
		Run: func() {
			log.Println("work 1 ", i)
			i ++
		},
	})
	c.AddJob(spec, cron.Job{
		Count: 5,
		Run: func() {
			log.Println("work 2 ", j)
			j ++
		},
	})
	c.Start()

	time.Sleep(time.Second * 10)

	c.AddJob(spec, cron.Job{
		Count: -1,
		Run: func() {
			log.Println("work 3 ", k)
			k ++
		},
	})

	c.RemoveSchedule(job1)
	select {}
}
