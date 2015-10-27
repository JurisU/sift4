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

// Getters & Setters

func (s *Sift4) SetTokenizer(fn func(string) []string) *Sift4 {
	s.tokenizer = fn
	return s
}
func (s *Sift4) GetTokenizer() func(string) []string {
	return s.tokenizer
}

func (s *Sift4) SetTokenMatcher(fn func(string, string) bool) *Sift4 {
	s.tokenMatcher = s
	return s
}
func (s *Sift4) GetTokenMatcher() func(string, string) bool {
	return s.tokenMatcher
}

func (s *Sift4) SetMatchingEvaluator(fn func(string, string) int) *Sift4 {
	s.matchingEvaluator = fn
	return s
}
func (s *Sift4) GetMatchingEvaluator() func(string, string) int {
	return s.matchingEvaluator
}

func (s *Sift4) SetLocalLengthEvaluator(fn func(int) int) *Sift4 {
	s.localLengthEvaluator = fn
	return s
}
func (s *Sift4) GetLocalLengthEvaluator() func(int) int {
	return s.localLengthEvaluator
}

func (s *Sift4) SetTranspositionCostEvaluator(fn func(string, string) int) *Sift4 {
	s.transpositionCostEvaluator = fn
	return s
}
func (s *Sift4) GetTranspositionCostEvaluator() func(string, string) int {
	return s.transpositionCostEvaluator
}

func (s *Sift4) SetTranspositionsEvaluator(fn func(int, int) int) *Sift4 {
	s.transpositionsEvaluator = fn
	return s
}
func (s *Sift4) GetTranspositionsEvaluator() func(int, int) int {
	return s.transpositionsEvaluator
}
