package main

import (
	"fmt"

	"github.com/fighterlyt/permutation"
)

func main() {
	work()

}
func work() {
	fmt.Println("a = {1, 2, 3, 4}")
	a := []int{1, 2, 3, 4}

	p, err := permutation.NewPerm(a, nil)
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
