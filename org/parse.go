package org

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type node struct {
	level   int
	todo    bool
	content string
}

func newNode(line string) node {
	todo, _ := regexp.Match("\\*+ TODO", []byte(line))
	return node{
		level:   strings.LastIndex(line, "*") + 1,
		todo:    todo,
		content: line,
	}
}

// ExtractRemaining extracts remaining tasks and any parent entries then returns
// a syntactically sound org-mode string.
func ExtractRemaining(fname string) string {
	out := ""

	f, err := os.Open(fname)
	if err != nil {
		return out
	}

	var stack []node

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && line[0] == '*' {
			stack = append([]node{newNode(line)}, stack...)
		}
	}

	lastLevel := 0
	for i := range stack {
		if stack[i].todo || stack[i].level < lastLevel {
			out = stack[i].content + "\n" + out
			lastLevel = stack[i].level
		}
	}

	return out
}
