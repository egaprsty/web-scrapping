package main

import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run scraper.go <url>")
        return
    }

    url := os.Args[1]

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error fetching the URL:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: Received non-OK status code:", resp.StatusCode)
        return
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    }

    title := findTitle(doc)
    if title != "" {
        fmt.Println("Title of the page:", title)
    } else {
        fmt.Println("Title not found")
    }
}

func findTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := findTitle(c); title != "" {
			return title
		}
	}

	return ""
}