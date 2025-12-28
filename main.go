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

	content, err := io.ReadAll(os.Stdin)
	if len(content) == 0 {
		fmt.Println("No file and nothing in Stdin")
	}
	return content, err

}

func main() {

	l := flag.Bool("l", false, "Print the newline counts")
	c := flag.Bool("c", false, "Print the byte counts")
	m := flag.Bool("m", false, "Print the character counts")
	w := flag.Bool("w", false, "Print the word counts")
	flag.Parse()

	filePath := os.Args[len(os.Args)-1]

	var content []byte
	var err error
	if filePath != "l" || filePath != "c" || filePath != "m" || filePath != "w" {
		content, err = readFromFile(filePath)
		if err != nil {
			content, err = readFromStdin()
			if err != nil {
				fmt.Println("Error ", err)
			}
		}
	}
	if !*l && !*c && !*m && !*w {
		*l = true
		*c = true
		*w = true
	}

	var lineCount, byteCount, wordCount, charCount int
	if *l {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineCount++
		}
	}
	if *c {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			byteCount++
		}
	}
	if *m {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			charCount++
		}
	}
	if *w {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordCount++
		}
	}

	if lineCount != 0 {
		fmt.Printf("Lines: \t%d\n", lineCount)
	}
	if wordCount != 0 {
		fmt.Printf("Words: \t%d\n", wordCount)
	}
	if byteCount != 0 {
		fmt.Printf("Bytes: \t%d\n", byteCount)
	}
	if charCount != 0 {
		fmt.Printf("Chars: \t%d\n", charCount)
	}
}
