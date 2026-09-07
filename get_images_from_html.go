package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("Error parseing HTML: %w", err)
	}

	var imagesURLs []string
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists {
			return
		}

		u, err := url.Parse(src)
		if err != nil {
			fmt.Printf("Error parsing src %q: %v", src, err)
			return
		}

		absloute := baseURL.ResolveReference(u)
		imagesURLs = append(imagesURLs, absloute.String())
	})

	return imagesURLs, nil
}
