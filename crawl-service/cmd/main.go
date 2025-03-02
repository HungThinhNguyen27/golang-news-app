package main

import (
	"crawl-service/crawler"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// load time in location VIET NAM
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Fatal("Failed to load location:", err)
	}

	c := cron.New()
	jobID, err := c.AddFunc("@every 120m", func() {
		log.Println("Starting scheduled crawling task at", time.Now().In(loc))
		crawler.CrawlArticles()
	})

	if err != nil {
		log.Fatal("failed to schedule crawling task:", err)
	}
	// start cron job
	c.Start()

	schedule := c.Entry(jobID).Next.In(loc)
	log.Println("Scheduler started. Crawling will run every  2 hour")
	log.Println("Next run scheduled at:", schedule)
	select {}
}
