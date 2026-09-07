package main

import (
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	heading, err := getHeadingFromHTML(html)
	if err != nil {
		heading = ""
	}

	firstParagraph, err := getFirstParagraphFromHTML(html)
	if err != nil {
		firstParagraph = ""
	}

	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			Heading:        heading,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	outgoingLinks, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		outgoingLinks = nil
	}

	imagesURLs, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		imagesURLs = nil
	}

	return PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imagesURLs,
	}
}
