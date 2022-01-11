package main

import (
    "fmt"
)


type Link struct {
    Href string
    Text string
}

func main() {
    l := Link{
        Href: "/dog",
        Text: "Something in a span Text not in a span Bold text!",
    }

    fmt.Printf("%s\n%s\n", l.Href, l.Text)
}