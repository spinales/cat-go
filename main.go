package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	u := flag.Bool("u", false, "Write bytes from the input file to the standard output without delay as each is read.")
	ver := flag.Bool("version", false, "output version information and exit.")
	help := flag.Bool("help", false, "display this help and exit.")
	flag.Parse()

	if *ver {
		version()
	}

	if *help {
		options()
	}

	for _, f := range flag.Args() {
		cat(f, *u)
	}
}

func options() {
	flag.PrintDefaults()
	os.Exit(0)
}

func version() {
	fmt.Println("0.0.1")
	os.Exit(0)
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
