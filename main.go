package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	data, err := os.ReadFile("ex4.html")
	if err != nil {
		panic(err)
	}

	var links []Link

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	var f func(*html.Node, bool) string
	f = func(n *html.Node, parentA bool) string {
		isA := parentA
		var link Link
		var text string
		if n.Type == html.ElementNode && n.Data == "a" {
			link = Link{
				Href: n.Attr[0].Val,
			}
			isA = true
		}
		if n.Type == html.TextNode && isA {
			text = n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			text += f(c, isA)
		}
		if (link != Link{}) {
			link.Text = text
			links = append(links, link)
			text = ""
		}
		return text
	}
	f(doc, false)
	for _, link := range links {
		fmt.Printf("\n%s: %s", link.Href, link.Text)
	}
}