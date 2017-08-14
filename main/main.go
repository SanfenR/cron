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

	_, entry := c.AddFunc(spec, func() {
		log.Println("start ", i)
		i ++

	})

	c.Start()

	time.Sleep(5 * time.Second)

	c.RemoveSchedule(entry)
	log.Println("remove")


	time.Sleep(5 * time.Second)

	c.AddFunc(spec, func() {
		log.Println("next ", i)
		i ++
	})

	time.Sleep(5 * time.Second)


	select {}
}
