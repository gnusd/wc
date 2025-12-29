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

func errMsg(err error) {

	if err != nil {
		fmt.Println("An error has occurd: ", err)
	}
}

func whichInput(filePath string, flags [5]string) ([]byte, bool) {
	var content []byte
	var err error
	var fp bool

	if strings.Contains(filePath, "wc") {
		content, err = readFromStdin()
		errMsg(err)
	} else if len(filePath) == 2 {
		for _, value := range flags {
			if filePath == value {
				content, err = readFromStdin()
				errMsg(err)
			}
		}
	} else {
		content, err = readFromFile(filePath)
		fp = true
		errMsg(err)
	}
	return content, fp
}

func count(content []byte, l bool, c bool, m bool, w bool) (int, int, int, int) {

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
	return lineCount, wordCount, charCount, byteCount

}

func printOutput(lineCount int, wordCount int, charCount int, byteCount int, p bool, filePath string, fp bool) {

	var n, l, c, m, w, f string
	if p {
		n = "\n"
		l = "Lines: "
		c = "Bytes: "
		m = "Chars: "
		w = "Words: "
		f = "File: "
	}
	if lineCount != 0 {
		fmt.Printf("%s\t%d%s", l, lineCount, n)
	}
	if wordCount != 0 {
		fmt.Printf("%s\t%d%s", w, wordCount, n)
	}
	if charCount != 0 {
		fmt.Printf("%s\t%d%s", m, charCount, n)
	}
	if byteCount != 0 {
		fmt.Printf("%s\t%d%s", c, byteCount, n)
	}
	if fp {
		fmt.Printf("%s\t%s%s", f, filePath, n)
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

	var flags [5]string
	flags[0] = "-l"
	flags[1] = "-c"
	flags[2] = "-m"
	flags[3] = "-w"
	flags[4] = "-p"

	filePath := os.Args[len(os.Args)-1]

	content, fp := whichInput(filePath, flags)
	lineCount, wordCount, charCount, byteCount := count(content, *l, *c, *m, *w)
	printOutput(lineCount, wordCount, charCount, byteCount, *p, filePath, fp)

}
