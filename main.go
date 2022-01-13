package main

import (
    "os"
    "fmt"
    "log"
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
    l := Link{
        Href: "/dog",
        Text: "Something in a span Text not in a span Bold text!",
    }

    fmt.Sprintf("%s\n%s\n", l.Href, l.Text)


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

//    fmt.Printf("%T\n", doc)

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
                        fmt.Println(l)
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

    fmt.Println()
    researchNode(doc)
}