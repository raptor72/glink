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
	node_string := `<body>
                  <h1>Hello!</h1>
                  <a href="/other-page">A link to another page</a>
             </body>`
	node := createNode(node_string)
	result := dfsLink(node, &l)
	expected_text := "A link to another page"
	expected_href := "/other-page"
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
	if result[0].Text != expected_text {
		t.Errorf("Expected text [%s], but got [%s]", expected_text, result[0].Text)
	}
    if result[0].Href != expected_href {
		t.Errorf("Expected href [%s], but got [%s]", expected_href, result[0].Href)
	}
}

func Test_dfsLink_from_node_with_text_in_tags(t *testing.T) {
	var l []Link
	node_string := `<html>
	<body>
	  <h1>Hello!</h1>
	  <a href="/page-one">
		  A link to first page
		  <span> some span </span>
	  </a>
	</body>
	</html>`
	node := createNode(node_string)
	result := dfsLink(node, &l)
	expected_text := "A link to first page some span"
	expected_href := "/page-one"
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
	if result[0].Text != expected_text {
		t.Errorf("Expected [%s], but got [%s]", expected_text, result[0].Text)
	}
    if result[0].Href != expected_href {
		t.Errorf("Expected href [%s], but got [%s]", expected_href, result[0].Href)
	}
}


func Test_dfsLink_from_node_with_coment(t *testing.T) {
	var l []Link
	node_string := `<html>
	<body>
	  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
	</body>
	</html>`
	node := createNode(node_string)
	result := dfsLink(node, &l)
	expected_text := "dog cat"
	expected_href := "/dog-cat"
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
	if result[0].Text != expected_text {
		t.Errorf("Expected [%s], but got [%s]", expected_text, result[0].Text)
	}
    if result[0].Href != expected_href {
		t.Errorf("Expected href [%s], but got [%s]", expected_href, result[0].Href)
	}
}

func Test_dfsLink_from_node_with_many_subnodes_same_depth(t *testing.T) {
	var l []Link
	node_string := `<html>
	<head>
	</head>
	<body>
		<a href="https://www.twitter.com/joncalhoun">
		  Check me out on twitter
		</a>
		<a href="https://github.com/gophercises">
		  Gophercises is on Github!
		</a>
	</body>
	</html>`
	node := createNode(node_string)
	result := dfsLink(node, &l)
	expected_1_text := "Check me out on twitter"
	expected_1_href := "https://www.twitter.com/joncalhoun"
	expected_2_text := "Gophercises is on Github!"
	expected_2_href := "https://github.com/gophercises"
	expected_printed_type := "[]glink.Link"
	expected_length := 2

	got_type := fmt.Sprintf("%T", result)
    got_length := len(result)
	if got_type != expected_printed_type {
		t.Errorf("Expected type [%s], but got [%s]", expected_printed_type, got_type)
	}
    if got_length != expected_length {
		t.Errorf("Expected length [%d], but got [%d]", expected_length, got_length)
	}
	if result[0].Text != expected_1_text {
		t.Errorf("Expected [%s], but got [%s]", expected_1_text, result[0].Text)
	}
    if result[0].Href != expected_1_href {
		t.Errorf("Expected href [%s], but got [%s]", expected_1_href, result[0].Href)
	}
	if result[1].Text != expected_2_text {
		t.Errorf("Expected [%s], but got [%s]", expected_2_text, result[0].Text)
	}
    if result[1].Href != expected_2_href {
		t.Errorf("Expected href [%s], but got [%s]", expected_2_href, result[0].Href)
	}
}


func Test_dfsLink_from_node_with_many_subnodes_different_depth(t *testing.T) {
	var l []Link
	node_string := `<html>
	<head>
	</head>
	<body>
    	<div>
            <div>
	            <a href="https://www.twitter.com/joncalhoun">
		            Check me out on twitter
            	</a>
            </div>
    	</div>
	    <a href="https://github.com/gophercises">
		    Gophercises is on Github!
    	</a>
	</body>
	</html>`
	node := createNode(node_string)
	result := dfsLink(node, &l)
	expected_1_text := "Check me out on twitter"
	expected_1_href := "https://www.twitter.com/joncalhoun"
	expected_2_text := "Gophercises is on Github!"
	expected_2_href := "https://github.com/gophercises"
	expected_printed_type := "[]glink.Link"
	expected_length := 2

	got_type := fmt.Sprintf("%T", result)
    got_length := len(result)
	if got_type != expected_printed_type {
		t.Errorf("Expected type [%s], but got [%s]", expected_printed_type, got_type)
	}
    if got_length != expected_length {
		t.Errorf("Expected length [%d], but got [%d]", expected_length, got_length)
	}
	if result[0].Text != expected_1_text {
		t.Errorf("Expected [%s], but got [%s]", expected_1_text, result[0].Text)
	}
    if result[0].Href != expected_1_href {
		t.Errorf("Expected href [%s], but got [%s]", expected_1_href, result[0].Href)
	}
	if result[1].Text != expected_2_text {
		t.Errorf("Expected [%s], but got [%s]", expected_2_text, result[0].Text)
	}
    if result[1].Href != expected_2_href {
		t.Errorf("Expected href [%s], but got [%s]", expected_2_href, result[0].Href)
	}
}
