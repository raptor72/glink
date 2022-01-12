package main

import (
    "os"
    "fmt"
    "log"
    "bufio"
//    "strings"
    "golang.org/x/net/html"
)



type Link struct {
    Href, Text string
}

func main() {
    l := Link{
        Href: "/dog",
        Text: "Something in a span Text not in a span Bold text!",
    }

    fmt.Printf("%s\n%s\n", l.Href, l.Text)


    file, err := os.Open("ex1.html")
    if err != nil {
         log.Fatal(err)
    }
    defer file.Close()
    rd := bufio.NewReader(file)

    doc, err := html.Parse(rd)
    if err != nil {
        log.Fatal(err)
    }

//    fmt.Println(doc)

    var f func(*html.Node)
    f = func(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
	    for _, a := range n.Attr {
		if a.Key == "href" {
		    fmt.Println(a.Val)
		    break
		}
	    }
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
	    f(c)
	}
    }
    f(doc)
}