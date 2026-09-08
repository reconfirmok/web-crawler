# Web Crawler

A concurrent web crawler written in Go. It starts from a base URL, follows internal links, and collects page data (heading, first paragraph, outgoing links, and image URLs) into a JSON report.

## Requirements

- Go 1.26 or later
- [goquery](https://github.com/PuerkitoBio/goquery) — HTML parsing and querying

## Installation

Clone the repo and build the binary:

```bash
go build -o crawler .
```

Or run it directly with `go run`:

```bash
go run .
```

## Usage

```bash
./crawler <baseURL> <maxConcurrency> <maxPages>
```

**Arguments:**

| Argument         | Description                                      |
|------------------|---------------------------------------------------|
| `baseURL`        | The URL to start crawling from                    |
| `maxConcurrency` | Max number of pages to crawl at the same time     |
| `maxPages`       | Max number of pages to visit before stopping      |

**Example:**

```bash
./crawler https://example.com 5 20
```

This crawls up to 20 pages on `example.com`, using up to 5 concurrent requests, and writes the results to `report.json` in the current directory.

## Output

The report is a JSON array of page objects, sorted by URL:

```json
[
  {
    "url": "https://example.com",
    "heading": "Example Domain",
    "first_paragraph": "This domain is for use in illustrative examples...",
    "outgoing_links": ["https://example.com/about"]
    "image_urls": ["https://example.com/images/logo.png"]
  }
]
```

