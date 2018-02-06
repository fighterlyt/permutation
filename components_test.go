package permutation

import (
	"log"
	"testing"
)

func TestFactorial0(t *testing.T) {
	reference := [][]int{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
	}
	for _, v := range reference {
		if factorial(v[0]) != v[1] {
			log.Fatal("Factorial problem", v[0], v[1])
		}
	}
}
func TestFactorial1(t *testing.T) {
	var tests = []struct {
		Name       string
		In, Result int
	}{
		{
			Name:   "0",
			In:     0,
			Result: 0,
		},
		{
			Name:   "1",
			In:     1,
			Result: 1,
		},
		{
			Name:   "2",
			In:     2,
			Result: 2,
		},
		{
			Name:   "3",
			In:     3,
			Result: 6,
		},
		{
			Name:   "4",
			In:     4,
			Result: 24,
		},
	}

	for _, test := range tests {
		if result := factorial(test.In); result != test.Result {
			t.Errorf("%s: result %d doesn't match expected result %d\n",
				result, test.Result)
		}
	}
}
