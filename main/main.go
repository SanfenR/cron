package main

import (
	"cron"
	"log"
)

func main() {

	c := cron.New()

	spec := "*/1 * * * * *"
	i := 0

	c.AddJob(spec, cron.Job{
		Count : 5,
		Run: func(){
			log.Println("start ", i)
			i ++
		},
	})

	c.Start()


	select {}
}
