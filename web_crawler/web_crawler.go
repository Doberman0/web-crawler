package web_crawler

import (
	"github.com/golang-collections/collections/queue"
)

type WebCrawler struct {
	VisitedURLs map[string]string
	URLsToVisit *queue.Queue // Change to queue of string
}

func New(
	entries []string,
) *WebCrawler {

	q := queue.New()
	for _, entry := range entries {
		q.Enqueue(entry)
	}

	return &WebCrawler{
		VisitedURLs: map[string]string{},
		URLsToVisit: q,
	}
}

func (wc *WebCrawler) GetVisitedUrls() map[string]string {
	return wc.VisitedURLs
}

func (wc *WebCrawler) GetUrlsToVisit() *queue.Queue {
	return wc.URLsToVisit
}

func (wc *WebCrawler) AddEntriesFromURL(currentUrl string) {

}
