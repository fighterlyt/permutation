package main

import (
	"fmt"
	"math"
	"reflect"

	"github.com/wayeast/permutation"
)

func main() {
	work()

}
func work() {
	fmt.Println("a = {-1, -2, -3, -4}")
	a := []int{-1, -2, -3, -4}

	fmt.Println("less = math.Abs(x) < math.Abs(y)")
	var l func(interface{}, interface{}) bool
	l = func(i, j interface{}) bool {
		vi := reflect.ValueOf(i).Int()
		vj := reflect.ValueOf(j).Int()
		return math.Abs(float64(vi)) < math.Abs(float64(vj))
	}

	p, err := permutation.NewPerm(a, l)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p.Index() -> p.Next()")
	for result, err := p.Next(); err == nil; result, err = p.Next() {
		fmt.Println(p.Index(), result.([]int))
	}
	p.Reset()

	fmt.Println("Results of p.NextN(30)")
	result := p.NextN(30)
	j := result.([][]int)
	for _, i := range j {
		fmt.Println(i)
	}
	fmt.Println(p.Index())
	p.Reset()

	fmt.Println("Results of p.NextN(1)")
	result = p.NextN(1)
	j = result.([][]int)
	for _, i := range j {
		fmt.Println(i)
	}
}
