package glink

import (
	"strings"
	"testing"
	"fmt"

	"golang.org/x/net/html"
)

func createNode(content string) (n *html.Node) {
	n, err := html.Parse(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
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
	expected := "I am just text"
	node := createNode(expected)
	result := dfsText(node)

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}

func Test_dfsText_from_html_Node(t *testing.T) {
	node_string := `<a href="/other-page">A link to another page</a>`
    node := createNode(node_string)   
    expected := "A link to another page"
    result := dfsText(node)
	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
} 

func Test_dfsLink_from_simple_node(t *testing.T) {
	var l []Link
	exp2 := `<body>
                  <h1>Hello!</h1>
                  <a href="/other-page">A link to another page</a>
             </body>`
	node := createNode(exp2)
	result := dfsLink(node, &l)
	expected := "A link to another page"
    expected_printed_type := "[]glink.Link"
	expected_length := 1

	got_type := fmt.Sprintf("%T", result)
    got_length := len(result)
	if got_type != expected_printed_type {
		t.Errorf("Expected type [%s], but got [%s]", expected_printed_type, got_type)
	}
    if got_length != expected_length {
		t.Errorf("Expected length [%d], but got [%d]", expected_length, got_length)
	}
	if result[0].Text != expected {
		t.Errorf("Expected [%s], but got [%s]", expected, result)
	}
}

