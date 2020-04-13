// https://www.hackerrank.com/challenges/weighted-uniform-string/problem

package main

import (
	"fmt"
	"regexp"
)

// Integer set
type IntSet struct {
	items []int
}

func (iSet *IntSet) Contains(element int) bool {
	for _, elem := range iSet.items {
		if elem == element {
			return true
		}
	}

	return false
}

func (iSet *IntSet) Add(element int) {
	if !iSet.Contains(element) {
		iSet.items = append(iSet.items, element)
	}
}

func NewIntSet() *IntSet {
	intSet := new(IntSet)
	intSet.items = make([]int, 0)

	return intSet
}

// Integer set

// Rune set
type RuneSet struct {
	items []rune
}

func (iSet *RuneSet) Contains(element rune) bool {
	for _, elem := range iSet.items {
		if elem == element {
			return true
		}
	}

	return false
}

func (iSet *RuneSet) Add(element rune) {
	if !iSet.Contains(element) {
		iSet.items = append(iSet.items, element)
	}
}

func (iSet *RuneSet) Items() []rune {
	var destination = make([]rune, len(iSet.items))
	copy(destination, iSet.items)

	return destination
}

func NewRuneSet() *RuneSet {
	runeSet := new(RuneSet)
	runeSet.items = make([]rune, 0)

	return runeSet
}

// Rune set

func distinctCharsInString(s string) []rune {
	var runeSet = NewRuneSet()

	for _, chr := range s {
		runeSet.Add(chr)
	}

	return runeSet.Items()
}

func weightedUniformStrings(s string, queries []int) []string {
	var scoreSet = NewIntSet()
	var answers = make([]string, len(queries))

	// Calculate all the possible values
	for _, chr := range distinctCharsInString(s) {
		var charScore = int(chr - 'a' + 1)
		var searchRegex, _ = regexp.Compile(string(chr) + "{2,}")

		// Search the longest contiguous character
		var matches = searchRegex.FindAllString(s, -1)
		var maxLength = 1

		for _, match := range matches {
			if len(match) > maxLength {
				maxLength = len(match)
			}
		}

		for i := 1; i <= maxLength; i++ {
			scoreSet.Add(charScore * i)
		}
	}

	// Now check all the queries
	for idx, query := range queries {
		if scoreSet.Contains(query) {
			answers[idx] = "Yes"
		} else {
			answers[idx] = "No"
		}
	}

	return answers
}

func main() {
	fmt.Println(weightedUniformStrings("abccddde", []int{1, 3, 12, 5, 9, 10}))
	fmt.Println(weightedUniformStrings("aaabbbbcccddd", []int{9, 7, 8, 12, 5}))
}
