package sift4

import (
	"strings"
)

// Models

type Sift4 struct {
	maxOffset                  int
	maxDistance                int
	tokenizer                  func(string) []string
	tokenMatcher               func(string, string) bool
	matchingEvaluator          func(string, string) float64
	localLengthEvaluator       func(float64) float64
	transpositionCostEvaluator func(int, int) float64
	transpositionsEvaluator    func(float64, float64) float64
}

// Initialization

func New() *Sift4 {
	return &Sift4{
		maxOffset:                  0,
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

func defaultMatchingEvaluator(t1, t2 string) float64 {
	return 1.0
}

func defaultLocalLengthEvaluator(i float64) float64 {
	return i
}

func defaultTranspositionCostEvaluator(c1, c2 int) float64 {
	return 1.0
}

func defaultTranspositionsEvaluator(lcss, trans float64) float64 {
	return lcss - trans
}

// Getters & Setters
func (s *Sift4) SetMaxOffset(i int) *Sift4 {
	s.maxOffset = i
	return s
}
func (s *Sift4) GetMaxOffset() int {
	return s.maxOffset
}

func (s *Sift4) SetTokenizer(fn func(string) []string) *Sift4 {
	s.tokenizer = fn
	return s
}
func (s *Sift4) GetTokenizer() func(string) []string {
	return s.tokenizer
}

func (s *Sift4) SetTokenMatcher(fn func(string, string) bool) *Sift4 {
	s.tokenMatcher = fn
	return s
}
func (s *Sift4) GetTokenMatcher() func(string, string) bool {
	return s.tokenMatcher
}

func (s *Sift4) SetMatchingEvaluator(fn func(string, string) float64) *Sift4 {
	s.matchingEvaluator = fn
	return s
}
func (s *Sift4) GetMatchingEvaluator() func(string, string) float64 {
	return s.matchingEvaluator
}

func (s *Sift4) SetLocalLengthEvaluator(fn func(float64) float64) *Sift4 {
	s.localLengthEvaluator = fn
	return s
}
func (s *Sift4) GetLocalLengthEvaluator() func(float64) float64 {
	return s.localLengthEvaluator
}

func (s *Sift4) SetTranspositionCostEvaluator(fn func(int, int) float64) *Sift4 {
	s.transpositionCostEvaluator = fn
	return s
}
func (s *Sift4) GetTranspositionCostEvaluator() func(int, int) float64 {
	return s.transpositionCostEvaluator
}

func (s *Sift4) SetTranspositionsEvaluator(fn func(float64, float64) float64) *Sift4 {
	s.transpositionsEvaluator = fn
	return s
}
func (s *Sift4) GetTranspositionsEvaluator() func(float64, float64) float64 {
	return s.transpositionsEvaluator
}
