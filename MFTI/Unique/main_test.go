package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testData = `1
2
3
3
4
5`

var testDataResult = `1
2
3
4
5`

func TestUnique(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testData))
	out := new(bytes.Buffer)

	err := Unique(in, out)
	if err != nil {
		t.Errorf("Test failed")
	}
	if out.String() != testDataResult {
		t.Errorf("Test failed - results not match")
	}
}
