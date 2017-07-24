package main

import (
	"testing"
)

func Test_random(t *testing.T) {
	for i := 0; i < 1000; i++ {
		r := random()
		if r < 0 || r > 2.0 {
			t.Error("return outside expected bounds:", r)
		}
	}
}

func Test_initWeights(t *testing.T) {
	var testcases = []struct {
		input int
	}{
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
	}

	for _, j := range testcases {
		w := initWeights(j.input)
		if len(w) != j.input {
			t.Errorf("Expected len %d for %d; got %d", j.input, j.input, len(w))
		}

		if len(w) > 0 {
			totalSame := 0
			first := w[0]
			for i := 1; i < len(w); i++ {
				if w[i] == first {
					totalSame++
				}
			}

			if totalSame == len(w)-1 {
				t.Errorf("expected random values for %d, but all equal", j.input)
			}
		}

	}

}
