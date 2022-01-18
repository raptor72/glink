package main

import (
    "os"
    "fmt"
    "log"
    "flag"
    "bufio"
    "strings"
    "golang.org/x/net/html"
)

type Link struct {
    Href, Text string
}

func readFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    return file, err
}

func parseHTML(n *html.Node, l *[]Link) {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                if n.FirstChild.Type == html.TextNode {
                    *l = append(*l, Link{a.Val, strings.TrimSpace(n.FirstChild.Data)})
                }
            break
            }
        }   
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        parseHTML(c, l)
    }
}

func main() {
    filename := flag.String("filename", "ex1.html", "a string filename to parse a html")
    flag.Parse()

    file, err := readFile(*filename)
    if err != nil {
        log.Fatal(err)
    }

    rd := bufio.NewReader(file)

    doc, err := html.Parse(rd)
    if err != nil {
        log.Fatal(err)
    }

    var lks []Link
    parseHTML(doc, &lks)
    fmt.Println(lks)
}