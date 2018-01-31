package permutation

import (
	"log"
	"testing"
)

func TestFactorial(t *testing.T) {
	reference := [][]int{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
	}
	for _, v := range reference {
		if calcFactorial(v[0]) != v[1] {
			log.Fatal("Factorial problem", v[0], v[1])
		}
	}
}
func TestPerm1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	e := [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 3, 2, 4},
		{1, 3, 4, 2},
		{1, 4, 2, 3},
		{1, 4, 3, 2},
		{2, 1, 3, 4},
		{2, 1, 4, 3},
		{2, 3, 1, 4},
		{2, 3, 4, 1},
		{2, 4, 1, 3},
		{2, 4, 3, 1},
		{3, 1, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 1, 4},
		{3, 2, 4, 1},
		{3, 4, 1, 2},
		{3, 4, 2, 1},
		{4, 1, 2, 3},
		{4, 1, 3, 2},
		{4, 2, 1, 3},
		{4, 2, 3, 1},
		{4, 3, 1, 2},
		{4, 3, 2, 1},
	}
	runTestGeneric(a, e)

}
func TestPerm2(t *testing.T) {
	a := []int{1, 2, 4, 4}
	e := [][]int{
		{1, 2, 4, 4},
		{1, 4, 2, 4},
		{1, 4, 4, 2},
		{2, 1, 4, 4},
		{2, 4, 1, 4},
		{2, 4, 4, 1},
		{4, 1, 2, 4},
		{4, 1, 4, 2},
		{4, 2, 1, 4},
		{4, 2, 4, 1},
		{4, 4, 1, 2},
		{4, 4, 2, 1},
		{4, 1, 2, 4},
		{4, 1, 4, 2},
		{4, 2, 1, 4},
		{4, 2, 4, 1},
		{4, 4, 1, 2},
		{4, 4, 2, 1},
		{4, 1, 2, 4},
		{4, 1, 4, 2},
		{4, 2, 1, 4},
		{4, 2, 4, 1},
		{4, 4, 1, 2},
		{4, 4, 2, 1},
	}
	runTestGeneric(a, e)
}
func TestPerm3(t *testing.T) {
	a := []string{"one", "Two"}
	e := [][]string{
		{"Two", "one"},
		{"one", "Two"},
	}
	runTestGeneric(a, e)
}

func runTestGeneric(a interface{}, e interface{}) {
	p, err := NewPerm(a, nil)
	if err != nil {
		log.Fatal(err)
	}
	for result, err := p.Next(); err == nil; result, err = p.Next() {
		if false {
			log.Println(result)
		} else {
			equalSliceGen(e, result, p.Index()-1)
		}
	}

	p.Reset()
	result := p.NextN(30)
	//checkSliceInt(e, j)

	_, leng, sz := equalSliceSliceGen(e, result)

	if calcFactorial(sz) != p.Index() {
		log.Fatal("Length is wrong", leng, p.Index())
	}

	p.Reset()
	result = p.NextN(1)
	_, leng, _ = equalSliceSliceGen(e, result)
	if leng != 1 {
		log.Fatal("p.NextN not returned expected length", result)
	}
}
