package main

import (
	"encoding/xml"
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/promisefemi/gophercises/link"
)

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls []loc `xml:"url"`
}

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "The URL that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 10, "The maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	var toXml urlset
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	enc := xml.NewEncoder(os.Stdout)

	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}

func bfs(urlStr string, maxDept int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDept; i++ {
		q, nq = nq, make(map[string]struct{})
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range getDestination(url) {
				nq[link] = struct{}{}
			}
		}
	}

	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func getDestination(destinationString string) []string {

	resp, err := http.Get(destinationString)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	baseUrl := &url.URL{
		Scheme: resp.Request.URL.Scheme,
		Host:   resp.Request.URL.Host,
	}
	base := baseUrl.String()
	pages := getHrefs(resp.Body, base)
	return filter(pages, withPrefix(base))

}

func getHrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	var hrefs []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	return hrefs
}

func filter(links []string, keepFn func(string) bool) []string {
	var filteredLinks []string
	for _, l := range links {
		if keepFn(l) {
			filteredLinks = append(filteredLinks, l)
		}
	}
	return filteredLinks
}
func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
