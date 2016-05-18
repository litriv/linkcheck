package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var in io.Reader

	switch len(os.Args[1:]) {
	case 1:
		f, err := os.Open(os.Args[1])
		if err != nil {
			fatal(err)
		}
		defer f.Close()
		in = f
	case 0:
		in = os.Stdin
	default:
		fatal(fmt.Errorf("wrong number of arguments"))
	}

	links := collect(in)

	rch := make(chan result)

	for _, l := range links {
		// check link, sending results on results channel
		go check(l, rch)
	}

	for i := 0; i < len(links); i++ {
		r := <-rch
		fmt.Println(r)
	}

}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// result hold the details for the result of a link check
type result struct {
}

func (r *result) String() string {
	fatal(fmt.Errorf("TODO:implement this function"))
	return fmt.Sprintln(r)
}

// check checks each link and sends the result on ch
func check(link string, ch chan<- result) {
	fatal(fmt.Errorf("TODO:implement this function"))
}

// collect collects the links from the input
func collect(in io.Reader) []string {
	fatal(fmt.Errorf("TODO:implement this function"))
	return []string{}
}
