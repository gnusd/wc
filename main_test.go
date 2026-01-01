package main

import (
	"fmt"
	"os"
	"testing"
)

func errCh(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func testFile() []byte {
	file := "testing/the_odysse.txt"
	content, err := os.ReadFile(file)
	errCh(err)
	return content
}

func TestLines(t *testing.T) {
	want := 12242
	got := countLines(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestBytes(t *testing.T) {
	want := 717826
	got := countBytes(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestChars(t *testing.T) {
	want := 710323
	got := countChars(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestWords(t *testing.T) {
	want := 132604
	got := countWords(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TesMaxWidth(t *testing.T) {
	want := 78
	got := countMaxWidth(testFile())
	if want != got {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
