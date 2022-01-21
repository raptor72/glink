package glink

import (
    "fmt"
	"strings"
	"testing"
	"golang.org/x/net/html"
)


func createNode(content string) (n *html.Node) {
	n, err := html.Parse(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
    fmt.Println(n)
	return n
}

// func createAnchorNode(href string, text string) *html.Node {
// 	textNode := createNode(text)
// 	attrs := make([]html.Attribute, 1)
// 	attrs[0] = html.Attribute{
// 		Namespace: "",
// 		Key:       "href",
// 		Val:       href,
// 	}
// 	return &html.Node{
// 		Parent:      nil,
// 		FirstChild:  textNode,
// 		LastChild:   textNode,
// 		PrevSibling: nil,
// 		NextSibling: nil,
// 		Type:        html.ElementNode,
// 		DataAtom:    0,
// 		Data:        "a",
// 		Namespace:   "",
// 		Attr:        attrs,
// 	}
// }

func Test_TraverseText_WhenGivenPlainText_ShouldReturnText(t *testing.T) {
	expected := "I am but text"
	node := createNode(expected)

	result := dfsText(node)

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}
