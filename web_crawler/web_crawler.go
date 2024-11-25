package web_crawler

type WebCrawler struct {
	VisitedURLs map[string]string
	URLsToVisit []string
}

func New(
	entries []string,
) *WebCrawler {
	return &WebCrawler{
		VisitedURLs: map[string]string{},
		URLsToVisit: entries,
	}
}

func (wc *WebCrawler) GetVisitedUrls() map[string]string {
	return wc.VisitedURLs
}

func (wc *WebCrawler) GetUrlsToVisit() []string {
	return wc.URLsToVisit
}

func (wc *WebCrawler) AddEntriesFromURL(currentUrl string) {

}