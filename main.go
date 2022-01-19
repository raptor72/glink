package main

import (
    // "io"
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

// fucn Parse(r io.Reader) ([]Link, error) {
//     return nil, nil
// }

func readFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    return file, err
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

func dfsLink(n *html.Node, l *[]Link) {
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
    dfsLink(doc, &lks)
    fmt.Println(lks)
}