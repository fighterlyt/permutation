package permutation

import "testing"

func TestFactorial(t *testing.T) {
	var tests = []struct {
		Name       string
		In, Result int
	}{
		{
			Name:   "0",
			In:     0,
			Result: 1,
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
