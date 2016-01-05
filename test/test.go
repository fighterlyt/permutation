package main
import (
	"../../permutation"
	"fmt"
)
func main(){
	work()
}

func work(){
    sample := []int { 1, 1, 3, 4 }

    fmt.Println("Generating new permutator.")
	perm, _ := permutation.NewPerm(sample, nil)

    fmt.Println("Iterating permutations.")
    
	for result, ok := perm.Next(); ok; result, ok = perm.Next(){
		fmt.Println(perm.Index(), result.([]int))
	}
    
	perm.Reset()
    fmt.Println("Reset permutation index.")    
    fmt.Println(perm.Index())

    fmt.Println("Iterating 30 permutations.")
    
	result := perm.NextN(30)
    
	iterations := result.([][]int)
    
	for _, iteration := range iterations {
		fmt.Println(iteration)
	}

	perm.Reset()
    fmt.Println("Reset permutation index.")    
	fmt.Println(perm.Index())
    
	result = perm.NextN(1)
    
	iterations = result.([][]int)
    
	for _, iteration := range iterations {
		fmt.Println(iteration)
	}
}