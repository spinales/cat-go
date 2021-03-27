package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	v  bool // -v, --show-nonprinting
	t  bool // -T, --show-tabs
	st bool // -s, --squeeze-blank
	n  bool // -n, --number
	e  bool // -E, --show-ends
	b  bool // -b, --number-nonblank
	vE bool // -e
)

func main() {
	// -e
	flag.BoolVar(&vE, "e", false, "equivalent to -vE")
	// -b, --number-nonblank
	flag.BoolVar(&b, "b", false, "number nonempty output lines, overrides -n")
	flag.BoolVar(&b, "number-nonblank", false, "number nonempty output lines, overrides -n")
	// -E, --show-ends
	flag.BoolVar(&e, "E", false, "display $ at end of each line")
	flag.BoolVar(&e, "show-ends", false, "display $ at end of each line")
	// -n, --number
	flag.BoolVar(&n, "n", false, "number all output lines")
	flag.BoolVar(&n, "number", false, "number all output lines")
	// -s, --squeeze-blank
	flag.BoolVar(&st, "s", false, "suppress repeated empty output lines")
	flag.BoolVar(&st, "squeeze-blank", false, "suppress repeated empty output lines")
	// -t
	vt := flag.Bool("t", false, "equivalent to -vT")
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

	if *vt {
		v, t = true, true
	}

	if vE {
		e, v = true, true
	}

	for _, f := range flag.Args() {
		cat(f, *u, v, t, st, n, e, b)
	}
}

func cat(filename string, u bool, nonprinting bool, tabs bool, suppress bool, number bool, dollar bool, numbernospacesblank bool) {
	result := openFile(filename)
	result = flags(result, nonprinting, tabs, suppress)
	if number {
		numbersLine(string(result))
	} else if u {
		fmt.Println(result)
	} else if dollar {
		dollarLine(string(result))
	} else if numbernospacesblank {
		numberNoSpacesBlank(string(result))
	} else {
		fmt.Println(string(result))
	}
}

func flags(arr []byte, nonprinting bool, tabs bool, suppress bool) []byte {
	if nonprinting {
		arr = invisibleChar(arr)
		nonprinting = false
		flags(arr, nonprinting, tabs, suppress)
	} else if tabs {
		arr = replaceTabs(arr)
		tabs = false
		flags(arr, nonprinting, tabs, suppress)
	} else if suppress {
		arr = suppressEmpty(arr)
		suppress = false
		flags(arr, nonprinting, tabs, suppress)
	}
	return arr
}

// number nonempty output lines
func numberNoSpacesBlank(value string) {
	count := 1
	for _, v := range strings.Split(value, "\n") {
		if v == "" {
			fmt.Println(v)
		} else {
			fmt.Println(count, v)
			count++
		}
	}
}

// number all output lines
func dollarLine(value string) {
	for _, v := range strings.Split(value, "\n") {
		fmt.Printf("%s%s\n", v, "$")
	}
}

// number all output lines
func numbersLine(value string) {
	for i, v := range strings.Split(value, "\n") {
		fmt.Printf("%v	%s\n", (i + 1), v)
	}
}

// suppress repeated empty output lines
func suppressEmpty(arr []byte) []byte {
	arr = bytes.Replace(arr, []byte{10, 10, 10}, []byte{10, 10}, 1)
	if bytes.Contains(arr, []byte{10, 10, 10}) {
		arr = suppressEmpty(arr)
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
