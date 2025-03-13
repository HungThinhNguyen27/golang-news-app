package main

import (
	"crawl-service/crawler"
	"crawl-service/storage/elasticsearch"
	"fmt"
	"log"
)

// func main() {

// 	// connect database
// 	// db := postgres.ConnectToDB()
// 	// postgres.CreateArticlesTable(db) // IF NOT EXIST

// 	// load time in location VIET NAM
// 	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
// 	if err != nil {
// 		log.Fatal("Failed to load location:", err)
// 	}
// 	c := cron.New()
// 	jobID, err := c.AddFunc("@every 120m", func() {
// 		log.Println("Starting scheduled crawling task at", time.Now().In(loc))
// 		crawler.CrawlArticles()
// 		// for _, article := range articles {
// 		// 	postgres.SaveArticle(db, article) // save articles to postgres
// 		// }

// 	})

// 	if err != nil {
// 		log.Fatal("failed to schedule crawling task:", err)
// 	}
// 	// start cron job
// 	c.Start()

// 	schedule := c.Entry(jobID).Next.In(loc)
// 	log.Println("Scheduler started. Crawling will run every  2 hour")
// 	log.Println("Next run scheduled at:", schedule)
// 	select {}
// }

func main() {
	// connect ElasticSearch

	// crawl
	articles := crawler.CrawlArticles()
	for _, article := range articles {
		fmt.Println("--------")
		err := elasticsearch.SaveDocument(article)
		if err != nil {
			log.Fatalf("Failed to save article: %v", err)
		}
	}
}
