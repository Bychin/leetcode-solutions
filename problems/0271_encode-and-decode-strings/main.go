package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Solution: use prefix to encode info about len of string
// and use separator to separate prefix from the string.

const separator = "#"

type Codec struct{}

// Encode encodes a list of strings to a single string.
func (c *Codec) Encode(strs []string) string {
	buf := &bytes.Buffer{}
	for _, s := range strs {
		buf.WriteString(strconv.Itoa(len(s)) + separator)
		buf.WriteString(s)
	}
	return buf.String()
}

// Decode decodes a single string to a list of strings.
func (c *Codec) Decode(strs string) []string {
	result := []string{}

	for i, n := 0, len(strs); i < n; {
		sepPos := strings.Index(strs[i:], separator)
		if sepPos == -1 || sepPos >= len(strs) {
			panic("unexpected encoding")
		}

		strLen, err := strconv.Atoi(strs[i : i+sepPos])
		if err != nil {
			panic("unexpected encoding")
		}

		result = append(result, strs[i+sepPos+1:i+sepPos+1+strLen])
		i += sepPos + 1 + strLen
	}

	return result
}

var testCases = [][]string{
	{"123", "456", "789"},
	{"hello", "wor#ld"},
	{"hello"},
	{},
}

func main() {
	var codec Codec
	for _, tc := range testCases {
		result := codec.Decode(codec.Encode(tc))
		fmt.Println(result, reflect.DeepEqual(tc, result))
	}
}
