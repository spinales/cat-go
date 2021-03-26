package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	v bool // -v, --show-nonprinting
	t bool // -T, --show-tabs
)

func main() {
	// -T, --show-tabs
	flag.BoolVar(&t, "T", false, "display TAB characters as ^I")
	flag.BoolVar(&t, "show-tabs", false, "display TAB characters as ^I")
	// -v, --show-nonprinting
	flag.BoolVar(&v, "v", false, "use ^ and M- notation, except for LFD and TAB")
	flag.BoolVar(&v, "show-nonprinting", false, "use ^ and M- notation, except for LFD and TAB")
	// u / byte output
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
		cat(f, *u, v, t)
	}
}

func cat(filename string, u bool, nonprinting bool, tabs bool) {
	result := openFile(filename)
	result = flags(result, nonprinting, tabs)
	if u {
		fmt.Println(result)
	} else {
		fmt.Println(string(result))
	}
}

func flags(arr []byte, nonprinting bool, tabs bool) []byte {
	if nonprinting == true {
		arr = invisibleChar(arr)
		nonprinting = false
		flags(arr, nonprinting, tabs)
	} else if tabs == true {
		arr = replaceTabs(arr)
		tabs = false
		flags(arr, nonprinting, tabs)
	}
	return arr
}

// replace tabs
func replaceTabs(arr []byte) []byte {
	fmt.Println(arr, []byte{9}, []byte("^I"))
	arr = bytes.ReplaceAll(arr, []byte{9}, []byte("^I"))
	fmt.Println(arr, []byte{9}, []byte("^I"))
	return arr
}

// replace invisible characters
func invisibleChar(arr []byte) []byte {
	for v := range invisible {
		arr = bytes.ReplaceAll(arr, []byte{v}, []byte(invisible[v]))
	}
	return arr
}

func openFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
}

func options() {
	flag.PrintDefaults()
	os.Exit(0)
}

func version() {
	fmt.Println("0.0.1")
	os.Exit(0)
}
