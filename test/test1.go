package main

import (
	"fmt"

	"../../permutation"
)

func main() {
	work()

}
func work() {
	a := []int{1, 2, 3, 4}

	p, err := permutation.NewPerm(a, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for result, err := p.Next(); err == nil; result, err = p.Next() {
		fmt.Println(p.Index(), result.([]int))
	}
	p.Reset()

	result := p.NextN(30)
	j := result.([][]int)
	for _, i := range j {
		fmt.Println(i)
	}
	fmt.Println(p.Index())

	p.Reset()
	result = p.NextN(1)
	j = result.([][]int)
	for _, i := range j {
		fmt.Println(i)
	}
}
