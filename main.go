package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	u := flag.Bool("u", false, "Write bytes from the input file to the standard output without delay as each is read.")
	ver := flag.Bool("version", false, "output version information and exit")
	flag.Parse()

	if *ver {
		fmt.Println("0.0.1")
	}

	for _, f := range flag.Args() {
		cat(f, *u)
	}
}

func cat(filename string, u bool) {
	result := openFile(filename)
	if u {
		fmt.Println(result)
	} else {
		fmt.Println(string(result))
	}
}

func openFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}
