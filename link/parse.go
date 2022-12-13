package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (a tag) in an HTML documents
type Link struct {
	Href, Text string
}

// Parse takes in a HTML ( as an io.Reader ) and retuurns a slice of links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var link Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = getText(n)
	return link
}

func getText(n *html.Node) string {
	var text string

	if n.Type == html.TextNode {
		text += strings.TrimSpace(n.Data)
		return text
	}
	if n.Type != html.ElementNode {
		return text
	}
	for children := n.FirstChild; children != nil; children = children.NextSibling {
		text += strings.TrimSpace(getText(children)) + " "
	}

	return text
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var returnNode []*html.Node

	// c is the first variable initialized int the loop,
	// c would be used in the first loop, as long as it satisfies the condition of the loop
	// after the first loop the value of c is changed to the NextSibling of the first value of c and so on.
	// and when there is no more NextSibling c would be nil which would satisfy the condition of the loop thereby closing the loop
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returnNode = append(returnNode, linkNodes(c)...)
	}
	return returnNode
}
