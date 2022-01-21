package main

import (
    "os"
    "fmt"
    "log"
    "flag"
    "bufio"
	"github.com/raptor72/glink"
)

func readFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    return file, err
}

func main() {
    filename := flag.String("filename", "exhamples/ex1.html", "a string filename to parse a html")
    flag.Parse()

    file, err := readFile(*filename)
    if err != nil {
        log.Fatal(err)
    }

    rd := bufio.NewReader(file)

	links, err := glink.Parse(rd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
	// fmt.Println(links)
}

