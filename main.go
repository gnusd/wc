package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type Count struct {
	Lines    int
	Bytes    int
	Chars    int
	Words    int
	MaxWidth int
}

type Sum struct {
	Lines    int
	Bytes    int
	Chars    int
	Words    int
	MaxWidth int
}

type Flags struct {
	l *bool
	c *bool
	m *bool
	w *bool
	L *bool
}

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

func checkFlags(flags Flags) Flags {
	if !*flags.l && !*flags.c && !*flags.m && !*flags.w && !*flags.L {
		*flags.l = true
		*flags.m = true
		*flags.w = true
	}
	return flags
}

func getCount(content []byte, flags Flags) Count {
	fl := checkFlags(flags)
	return Count{
		Lines:    countLines(content, *fl.l),
		Bytes:    countBytes(content, *fl.c),
		Chars:    countChars(content, *fl.m),
		Words:    countWords(content, *fl.w),
		MaxWidth: countMaxWidth(content, *fl.L),
	}
}

func initializeSum() Sum {
	return Sum{
		Lines:    0,
		Bytes:    0,
		Chars:    0,
		Words:    0,
		MaxWidth: 0,
	}
}

func addSum(sum *Sum, count Count) {
	sum.Lines += count.Lines
	sum.Bytes += count.Bytes
	sum.Chars += count.Chars
	sum.Words += count.Words
	if count.MaxWidth > sum.MaxWidth {
		sum.MaxWidth = count.MaxWidth
	}
}

func countLines(content []byte, flags bool) int {
	var lineCount int
	if flags {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineCount++
		}
	}
	return lineCount
}
func countBytes(content []byte, flags bool) int {
	var byteCount int
	if flags {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			byteCount++
		}
	}
	return byteCount
}

func countChars(content []byte, flags bool) int {
	var charCount int
	if flags {
		charCount = utf8.RuneCount(content)
	}
	return charCount
}

func countWords(content []byte, flags bool) int {
	var wordCount int
	if flags {
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordCount++
		}
	}
	return wordCount
}

func countMaxWidth(content []byte, flags bool) int {
	var maxWidth int
	if flags {
		scanner := bufio.NewScanner(bytes.NewReader(content))

		for scanner.Scan() {
			currentWidth := utf8.RuneCountInString(scanner.Text())
			if currentWidth > maxWidth {
				maxWidth = currentWidth
			}
		}
	}
	return maxWidth
}

func (count Count) printOutput(filePath string) {
	if count.Lines != 0 {
		fmt.Printf("\t%d ", count.Lines)
	}
	if count.Words != 0 {
		fmt.Printf("\t%d ", count.Words)
	}
	if count.Chars != 0 {
		fmt.Printf("\t%d ", count.Chars)
	}
	if count.Bytes != 0 {
		fmt.Printf("\t%d ", count.Bytes)
	}
	if count.MaxWidth != 0 {
		fmt.Printf("\t%d ", count.MaxWidth)
	}
	if filePath != "" {
		fmt.Printf("\t%s", filePath)
	}
	fmt.Println("")
}

func (count Count) printEndOutput(sum Sum) {
	if sum.Lines != 0 {
		fmt.Printf("\t%d ", sum.Lines)
	}
	if sum.Words != 0 {
		fmt.Printf("\t%d ", sum.Words)
	}
	if sum.Chars != 0 {
		fmt.Printf("\t%d ", sum.Chars)
	}
	if sum.Bytes != 0 {
		fmt.Printf("\t%d ", sum.Bytes)
	}
	if sum.MaxWidth != 0 {
		fmt.Printf("\t%d ", sum.MaxWidth)
	}
	fmt.Println("\ttotal")
}

func handleFiles(files []string, flags Flags) {
	total := initializeSum()
	for i := range files {
		if len(files) == 1 || i < len(files)-1 {
			content, err := readFromFile(files[i])
			errMsg(err)
			counted := getCount(content, flags)
			counted.printOutput(files[i])
			addSum(&total, counted)
		} else {
			content, err := readFromFile(files[i])
			errMsg(err)
			counted := getCount(content, flags)
			addSum(&total, counted)
			counted.printOutput(files[i])
			counted.printEndOutput(total)
		}
	}
}

func handleStdin(flags Flags) {
	var files string
	content, err := readFromStdin()
	errMsg(err)
	counted := getCount(content, flags)
	counted.printOutput(files)
}

func handleArgs(args []string, flags Flags) {
	var files []string
	for index, value := range args {
		if index != 0 && !strings.HasPrefix(value, "-") {
			files = append(files, value)
		}
	}
	if len(files) >= 1 {
		handleFiles(files, flags)
	} else {
		handleStdin(flags)
	}
}

func (flags *Flags) returnFlags() {
	flags.l = flag.Bool("l", false, "Print the newline counts")
	flags.c = flag.Bool("c", false, "Print the byte counts")
	flags.m = flag.Bool("m", false, "Print the character counts")
	flags.w = flag.Bool("w", false, "Print the word counts")
	flags.L = flag.Bool("L", false, "Print the maximum display width")
}

func main() {
	var flags Flags
	flags.returnFlags()
	flag.Parse()
	handleArgs(os.Args, flags)
}
