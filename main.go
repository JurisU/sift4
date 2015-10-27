package sift4

import (
	"strings"
)

// Models

type Sift4 struct {
	maxDistance                int
	tokenizer                  func(string) []string
	tokenMatcher               func(string, string) bool
	matchingEvaluator          func(string, string) int
	localLengthEvaluator       func(int) int
	transpositionCostEvaluator func(string, string) int
	transpositionsEvaluator    func(int, int) int
}

// Initialization

func New() *Sift4 {
	return &Sift4{
		maxDistance:                0,
		tokenizer:                  defaultTokenizer,
		tokenMatcher:               defaultTokenMatcher,
		matchingEvaluator:          defaultMatchingEvaluator,
		localLengthEvaluator:       defaultLocalLengthEvaluator,
		transpositionCostEvaluator: defaultTranspositionCostEvaluator,
		transpositionsEvaluator:    defaultTranspositionsEvaluator,
	}
}

// Defaults

func defaultTokenizer(s string) []string {
	return strings.Split(s, "")
}

func defaultTokenMatcher(t1, t2 string) bool {
	return t1 == t2
}

func defaultMatchingEvaluator(t1, t2 string) int {
	return 1
}

func defaultLocalLengthEvaluator(i int) int {
	return i
}

func defaultTranspositionCostEvaluator(c1, c2 int) int {
	return 1
}

func defaultTranspositionsEvaluator(lcss, trans int) int {
	return lcss - trans
}
