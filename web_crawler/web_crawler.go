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

func (wc *WebCrawler) getRelativeLinksFromUrl(url string) []string {
	var relativeUrls []string
	return relativeUrls
}

func (wc *WebCrawler) addEntriesFromURL() {
	currentUrl := wc.URLsToVisit.Dequeue().(string)
	urlsToAdd := wc.getRelativeLinksFromUrl(currentUrl)
	for _, url := range urlsToAdd {
		wc.URLsToVisit.Enqueue(url)
	}
}

func (wc *WebCrawler) Crawl() {

}
