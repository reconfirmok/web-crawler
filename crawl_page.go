package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	parsedURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Hostname() != parsedURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing URL '%s': %v", rawCurrentURL, err)
		return
	}

	if !cfg.addPageVisit(normalizedURL) {
		return
	}

	fmt.Printf("Crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting HTML '%s': %v", rawCurrentURL, err)
		return
	}

	PageData := extractPageData(htmlBody, rawCurrentURL)
	cfg.setPageData(normalizedURL, PageData)

	for _, nextURL := range PageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
