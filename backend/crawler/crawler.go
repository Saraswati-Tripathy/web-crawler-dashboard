package crawler

import (
	"crawler-backend/models"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CrawlPage(url string) (*models.CrawlResult, error) {
	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Load HTML into goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Extract information
	htmlVersion := getHTMLVersion(doc)
	title := doc.Find("title").Text()

	h1Count := doc.Find("h1").Length()
	h2Count := doc.Find("h2").Length()
	h3Count := doc.Find("h3").Length()

	internalLinks, externalLinks, inaccessible := countLinks(doc, url)
	hasLogin := hasLoginForm(doc)

	result := &models.CrawlResult{
		URL:               url,
		Title:             title,
		HTMLVersion:       htmlVersion,
		H1Count:           h1Count,
		H2Count:           h2Count,
		H3Count:           h3Count,
		InternalLinks:     internalLinks,
		ExternalLinks:     externalLinks,
		InaccessibleLinks: inaccessible,
		HasLoginForm:      hasLogin,
	}

	return result, nil
}

func getHTMLVersion(doc *goquery.Document) string {
	node := doc.Nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == "xmlns" {
			return "HTML5"
		}
	}
	return "Unknown"
}

func countLinks(doc *goquery.Document, baseURL string) (internal int, external int, inaccessible int) {
	baseDomain := getDomain(baseURL)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			if strings.HasPrefix(href, "http") {
				if strings.Contains(href, baseDomain) {
					internal++
				} else {
					external++
				}
				// Check accessibility
				resp, err := http.Head(href)
				if err != nil || resp.StatusCode >= 400 {
					inaccessible++
				}
			} else if strings.HasPrefix(href, "/") {
				internal++
			}
		}
	})
	return
}

func hasLoginForm(doc *goquery.Document) bool {
	found := false
	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		if s.Find("input[type='password']").Length() > 0 {
			found = true
		}
	})
	return found
}

func getDomain(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 3 {
		return parts[2]
	}
	return url
}
