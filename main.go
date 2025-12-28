package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func readFromStdin() ([]byte, error) {

	return io.ReadAll(os.Stdin)

}

func whichInput(filePath string) []byte {
	var content []byte
	var err error

	if strings.Contains(filePath, "wc") {
		content, err = readFromStdin()
		if err != nil {
			fmt.Println("Error ", err)
		}
	} else if len(filePath) > 1 {
		content, err = readFromFile(filePath)
		if err != nil {
			content, err = readFromStdin()
			if err != nil {
				fmt.Println("Error ", err)
			}
		}
	}
	return content
}

func count(content []byte, l bool, c bool, m bool, w bool, p bool) {

	if !l && !c && !m && !w {
		l = true
		m = true
		w = true
	}

	var lineCount, byteCount, wordCount, charCount int

	if l {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineCount++
		}
	}
	if c {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			byteCount++
		}
	}
	if m {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			charCount++
		}
	}
	if w {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordCount++
		}
	}
	printOutput(lineCount, wordCount, charCount, byteCount, p)

}

func printOutput(lineCount int, wordCount int, charCount int, byteCount int, p bool) {

	var l string
	var n string
	if p {
		n = "\n"
	}
	if lineCount != 0 {
		if p {
			l = "Lines: "
		}
		fmt.Printf("%s\t%d%s", l, lineCount, n)
	}
	if wordCount != 0 {
		if p {
			l = "Words: "
		}
		fmt.Printf("%s\t%d%s", l, wordCount, n)
	}
	if charCount != 0 {
		if p {
			l = "Chars: "
		}
		fmt.Printf("%s\t%d%s", l, charCount, n)
	}
	if byteCount != 0 {
		if p {
			l = "Bytes: "
		}
		fmt.Printf("%s\t%d%s", l, byteCount, n)
	}

	if !p {
		fmt.Println("")
	}
}

func main() {

	l := flag.Bool("l", false, "Print the newline counts")
	c := flag.Bool("c", false, "Print the byte counts")
	m := flag.Bool("m", false, "Print the character counts")
	w := flag.Bool("w", false, "Print the word counts")
	p := flag.Bool("p", false, "Print number of and titles on single line")
	flag.Parse()

	filePath := os.Args[len(os.Args)-1]

	count(whichInput(filePath), *l, *c, *m, *w, *p)

}
