package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func errCh(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

var file string = "testing/the_odysse.txt"

func testFile() []byte {
	content, err := os.ReadFile(file)
	errCh(err)
	return content
}

func UWC(flag string) int {
	cmd := exec.Command("wc", flag, file)
	stdout, err := cmd.Output()
	errCh(err)
	out := strings.Split(string(stdout), " ")
	count, err := strconv.Atoi(out[0])
	errCh(err)
	return count
}

func TestLines(t *testing.T) {
	want := UWC("-l")

	got := countLines(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestBytes(t *testing.T) {
	want := UWC("-c")
	got := countBytes(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestChars(t *testing.T) {
	want := UWC("-m")
	got := countChars(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestWords(t *testing.T) {
	want := UWC("-w")
	got := countWords(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TesMaxWidth(t *testing.T) {
	want := UWC("-L")
	got := countMaxWidth(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
