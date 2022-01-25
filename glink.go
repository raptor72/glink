package glink

import (
    "io"
    "strings"
    "golang.org/x/net/html"
)

type Link struct {
    Href, Text string
}

func dfsText(n *html.Node) string {
    var s string
    if n.Type == html.TextNode {
        s += n.Data
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        s += dfsText(c) + " "
    }
    return strings.Join(strings.Fields(s), " ")
}

func dfsLink(n *html.Node, l *[]Link) []Link {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                if n.FirstChild.Type == html.TextNode {
                    *l = append(*l, Link{a.Val, dfsText(n)})
                    break
                }
            }
        }   
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        dfsLink(c, l)
    }
    return *l
}

func Parse(r io.Reader) ([]Link, error) {
    doc, err := html.Parse(r)
    if err != nil {
        return nil, err
    }
    var lks []Link
    parsed_lks := dfsLink(doc, &lks)
    return parsed_lks, nil
}
