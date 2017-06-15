package permutation

import (
	"math"
	"reflect"
	"testing"
)

func Test123(t *testing.T) {
	var k = []int{1, 2, 3}
	var nexts = map[int][]int{
		1: {1, 2, 3},
		2: {1, 3, 2},
		3: {2, 1, 3},
		4: {2, 3, 1},
		5: {3, 1, 2},
		6: {3, 2, 1},
	}

	p, err := NewPerm(k, nil)
	if err != nil {
		t.Errorf("Error creating permutator: '%s'\n", err)
	}
	if p.amount != 6 {
		t.Errorf("Error setting permutator amount: got %d\n", p.amount)
	}

	i := 1
	for next, err := p.Next(); err == nil; next, err = p.Next() {
		if p.Index() != i {
			t.Errorf("Index mismatch: i = %d; p.Index = %d\n", i, p.Index())
		}
		if reflect.TypeOf(next) != reflect.TypeOf(nexts[i]) {
			t.Errorf("Type mismatch at index %d: expected %v, got %v\n",
				i, reflect.TypeOf(nexts[i]), reflect.TypeOf(next))
		}
		for j, elem := range nexts[i] {
			if v := int(reflect.ValueOf(next).Index(j).Int()); v != elem {
				t.Errorf("Mismatch at next %d, element %d: expectd %d got %d\n",
					i, j, nexts[i][j], v)
			}
		}
		i++
	}
}

func TestABC(t *testing.T) {
	var k = []string{"A", "B", "C"}
	var nexts = map[int][]string{
		1: {"A", "B", "C"},
		2: {"A", "C", "B"},
		3: {"B", "A", "C"},
		4: {"B", "C", "A"},
		5: {"C", "A", "B"},
		6: {"C", "B", "A"},
	}

	p, err := NewPerm(k, nil)
	if err != nil {
		t.Errorf("Error creating permutator: '%s'\n", err)
	}
	if p.amount != 6 {
		t.Errorf("Error setting permutator amount: got %d\n", p.amount)
	}

	i := 1
	for next, err := p.Next(); err == nil; next, err = p.Next() {
		if p.Index() != i {
			t.Errorf("Index mismatch: i = %d; p.Index = %d\n", i, p.Index())
		}
		if reflect.TypeOf(next) != reflect.TypeOf(nexts[i]) {
			t.Errorf("Type mismatch at index %d: expected %v, got %v\n",
				i, reflect.TypeOf(nexts[i]), reflect.TypeOf(next))
		}
		for j, elem := range nexts[i] {
			if v := reflect.ValueOf(next).Index(j).String(); v != elem {
				t.Errorf("Mismatch at next %d, element %d: expectd %s got %s\n",
					i, j, nexts[i][j], v)
			}
		}
		i++
	}
}

func TestCustomLess(t *testing.T) {
	var k = []int{-1, -2, -3}
	var nexts = map[int][]int{
		1: {-1, -2, -3},
		2: {-1, -3, -2},
		3: {-2, -1, -3},
		4: {-2, -3, -1},
		5: {-3, -1, -2},
		6: {-3, -2, -1},
	}
	var l func(interface{}, interface{}) bool
	l = func(i, j interface{}) bool {
		vi := reflect.ValueOf(i).Int()
		vj := reflect.ValueOf(j).Int()
		return math.Abs(float64(vi)) < math.Abs(float64(vj))
	}

	p, err := NewPerm(k, l)
	if err != nil {
		t.Errorf("Error creating permutator: '%s'\n", err)
	}
	if p.amount != 6 {
		t.Errorf("Error setting permutator amount: got %d\n", p.amount)
	}

	i := 1
	for next, err := p.Next(); err == nil; next, err = p.Next() {
		if p.Index() != i {
			t.Errorf("Index mismatch: i = %d; p.Index = %d\n", i, p.Index())
		}
		if reflect.TypeOf(next) != reflect.TypeOf(nexts[i]) {
			t.Errorf("Type mismatch at index %d: expected %v, got %v\n",
				i, reflect.TypeOf(nexts[i]), reflect.TypeOf(next))
		}
		for j, elem := range nexts[i] {
			if v := int(reflect.ValueOf(next).Index(j).Int()); v != elem {
				t.Errorf("Mismatch at next %d, element %d: expectd %d got %d\n",
					i, j, nexts[i][j], v)
			}
		}
		i++
	}
}
