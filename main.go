package sift4

import (
	"math"
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

type offset struct {
	c1    int
	c2    int
	trans bool
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

// Methods

func (s *Sift4) Distance(s1, s2 string) float64 {

	// Tokenize strings
	t1 := s.tokenizer(s1)
	t2 := s.tokenizer(s2)

	// Length of token sets
	l1 := len(t1)
	l2 := len(t2)

	if l1 == 0 {
		return float64(l2)
	}
	if l2 == 0 {
		return float64(l1)
	}

	c1 := 0         // cursor for string 1
	c2 := 0         // cursor for string 2
	lcss := 0.0     // largest common subsequence
	local_cs := 0.0 // local common substring
	trans := 0.0    // number of transpositions

	offset_arr := []offset{}

	for (c1 < l1) && (c2 < l2) {
		if s.tokenMatcher(t1[c1], t2[c2]) {
			is_trans := false

			local_cs += s.matchingEvaluator(t1[c1], t2[c2])

			// check if current pair is a transposition
			i := 0
			for i < len(offset_arr) {
				ofs := offset_arr[i]
				if c1 <= ofs.c1 || c2 <= ofs.c2 {
					is_trans = math.Abs(float64(c2-c1)) >= math.Abs(float64(ofs.c2-ofs.c1))
					if is_trans {
						trans += s.transpositionCostEvaluator(c1, c2)
					} else if !ofs.trans {
						ofs.trans = true
						trans += s.transpositionCostEvaluator(ofs.c1, ofs.c2)
					}
					break
				} else {
					if c1 > ofs.c2 && c2 > ofs.c1 {
						offset_arr = append(offset_arr[:i], offset_arr[i+1:]...)
					} else {
						i++
					}
				}
			}
			offset_arr = append(offset_arr, offset{c1, c2, is_trans})
		} else {
			lcss += s.localLengthEvaluator(local_cs)
			local_cs = 0
			if c1 != c2 {
				c1 = int(math.Min(float64(c1), float64(c2)))
				c2 = c1
			}

			for i := 0; i < s.maxOffset && (c1+i < l1 || c2+i < l2); i++ {
				if c1+i < l1 && s.tokenMatcher(t1[c1+i], t2[c2]) {
					c1 += i - 1
					c2--
					break
				}
				if c2+i < l2 && s.tokenMatcher(t1[c1], t2[c2+i]) {
					c1--
					c2 += i - 1
					break
				}
			}
		}

		c1++
		c2++

		if s.maxDistance > 0 {
			tmp_dist := s.localLengthEvaluator(math.Max(float64(c1), float64(c2))) - s.transpositionsEvaluator(lcss, trans)
			if tmp_dist >= float64(s.maxDistance) {
				return math.Floor(tmp_dist + .5)
			}
		}

		if c1 >= l1 || c2 >= l2 {
			lcss += s.localLengthEvaluator(local_cs)
			local_cs = 0
			c1 = int(math.Min(float64(c1), float64(c2)))
			c2 = c1
		}
	}

	lcss += s.localLengthEvaluator(local_cs)
	return math.Floor(s.localLengthEvaluator(math.Max(float64(l1), float64(l2))) - s.transpositionsEvaluator(lcss, trans) + .5)
}
