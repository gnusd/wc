package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var sumLine, sumWord, sumByte, sumChar int

func readFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func readFromStdin() ([]byte, error) {

	return io.ReadAll(os.Stdin)

}

func errMsg(err error) {

	if err != nil {
		fmt.Println("An error has occurd: ", err)
		os.Exit(1)
	}
}

func count(content []byte, flags [8]bool) (int, int, int, int) {

	if !flags[0] && !flags[1] && !flags[2] && !flags[3] {
		flags[0] = true
		flags[2] = true
		flags[3] = true
	}

	var lineCount, byteCount, wordCount, charCount int

	if flags[0] {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineCount++
		}
	}
	if flags[1] {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			byteCount++
		}
	}
	if flags[2] {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			charCount++
		}
	}
	if flags[3] {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordCount++
		}
	}

	if flags[6] {
		sumLine += lineCount
		sumByte += byteCount
		sumChar += charCount
		sumWord += wordCount
	}

	return lineCount, wordCount, charCount, byteCount

}

func printOutput(lineCount int, wordCount int, charCount int, byteCount int, filePath string, flags [8]bool) {

	var n, l, c, m, w, f, t, total string
	total = "total"
	if flags[4] {
		n = "\n"
		l = "Lines: "
		c = "Bytes: "
		m = "Chars: "
		w = "Words: "
		f = "File: "
		t = "Total: "
	}
	if filePath == "" {
		flags[5] = false
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
	if flags[5] {
		fmt.Printf("%s\t%s%s", f, filePath, n)
	}

	if !flags[4] && flags[6] {
		fmt.Println("")
	} else if flags[4] && flags[6] {
		fmt.Println("")
	} else if !flags[4] {
		fmt.Println("")
	}
	if flags[7] {
		if sumLine != 0 {
			fmt.Printf("%s\t%d%s", l, sumLine, n)
		}
		if sumWord != 0 {
			fmt.Printf("%s\t%d%s", c, sumWord, n)
		}
		if sumChar != 0 {
			fmt.Printf("%s\t%d%s", m, sumChar, n)
		}
		if sumByte != 0 {
			fmt.Printf("%s\t%d%s", w, sumByte, n)
		}
		if flags[5] {
			fmt.Printf("%s\t%s%s", t, total, n)
		}
		if !flags[4] {
			fmt.Println("")
		}

	}
}
func handleFiles(files []string, flags [8]bool) {
	for i := range files {
		if i < len(files)-1 {
			content, err := readFromFile(files[i])
			errMsg(err)
			flags[6] = true
			flags[5] = true
			lineCount, wordCount, charCount, byteCount := count(content, flags)
			printOutput(lineCount, wordCount, charCount, byteCount, files[i], flags)
		} else {
			content, err := readFromFile(files[i])
			errMsg(err)
			flags[6] = true
			flags[5] = true
			flags[7] = true
			lineCount, wordCount, charCount, byteCount := count(content, flags)
			printOutput(lineCount, wordCount, charCount, byteCount, files[i], flags)
		}
	}
}

func handleStdin(content []byte, flags [8]bool) {
	var files string
	lineCount, wordCount, charCount, byteCount := count(content, flags)
	printOutput(lineCount, wordCount, charCount, byteCount, files, flags)
}

func handleArgs(args []string, flags [8]bool) {
	var files []string
	var content []byte
	var err error

	for index, value := range args {
		if index != 0 && !strings.HasPrefix(value, "-") {
			files = append(files, value)
		}
	}
	if len(files) >= 1 {
		handleFiles(files, flags)
	} else {
		content, err = readFromStdin()
		errMsg(err)
		handleStdin(content, flags)
	}
}

func main() {

	var flags [8]bool
	l := flag.Bool("l", false, "Print the newline counts")
	c := flag.Bool("c", false, "Print the byte counts")
	m := flag.Bool("m", false, "Print the character counts")
	w := flag.Bool("w", false, "Print the word counts")
	p := flag.Bool("p", false, "Print number of and titles on single line")
	flag.Parse()

	flags[0] = *l
	flags[1] = *c
	flags[2] = *m
	flags[3] = *w
	flags[4] = *p
	handleArgs(os.Args, flags)

}
