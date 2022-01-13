package main

import (
    "os"
    "fmt"
    "log"
    "flag"
    "bufio"
    "golang.org/x/net/html"
)

type Link struct {
    Href, Text string
}

func researchNode(n *html.Node) {
    fmt.Printf("node.Type: %v - %T\n", n.Type, n.Type)
    fmt.Printf("node.DataAtom: %v - %T\n", n.DataAtom, n.DataAtom)
    fmt.Printf("node.Data: %v - %T\n", n.Data, n.Data)
    fmt.Printf("node.Namespace: %v - %T\n", n.Namespace, n.Namespace)
    fmt.Printf("node.Attr: %v - %T\n", n.Attr, n.Attr)
}


func main() {
    filename := flag.String("filename", "ex1.html", "a string dilename to parse a html")
    flag.Parse()

    l := Link{
        Href: "/dog",
        Text: "Something in a span Text not in a span Bold text!",
    }

    fmt.Sprintf("%s\n%s\n", l.Href, l.Text)

    fmt.Println(*filename)
    file, err := os.Open(*filename)
    if err != nil {
         log.Fatal(err)
    }
    defer file.Close()
    rd := bufio.NewReader(file)

    doc, err := html.Parse(rd)
    if err != nil {
        log.Fatal(err)
    }

//    fmt.Printf("%T\n", doc)

    var links []Link
    var f func(*html.Node)

    f = func(n *html.Node) {
        // fmt.Println(n.Data, n.Type)
            if n.Type == html.ElementNode && n.Data == "a" {
    	    for _, a := range n.Attr {
            // fmt.Println(a)
                if a.Key == "href" {
		            // fmt.Println(a.Val)
                    if n.FirstChild.Type == html.TextNode {
                        // fmt.Println(n.FirstChild.Data)
                        l = Link{a.Val, n.FirstChild.Data}
                        // fmt.Println(l)
                        links = append(links, l)
                    }
                    break
		        }
            }
	    }
	    for c := n.FirstChild; c != nil; c = c.NextSibling {
    	    f(c)
	    }
    }
    f(doc)

    fmt.Println(links)
    researchNode(doc)
}