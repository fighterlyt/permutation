permutation
===========

A permutation library for Golang.

  Use NewPerm to generate a Pemutator, the argument k must be a non-nil slice,and the less argument must be a Less function that implements compare functionality of k's element type. If k's element is ordered,less argument can be nil.For ordered in Golang, visit http://golang.org/ref/spec#Comparison_operators
	
	func NewPerm(k interface{}, less Less) (*Permutator, error) 

  After generating a Permutator, the argument k can be modified and deleted,Permutator store a copy of k internel.A Permutator can  be used concurrently
  
  Invoke Permutator.Next() to return the next permutation in lexcial order.If all permutations generated,return an error
	
	func (p *Permutator) Next()(interface{}, error)
	
  The returned interface{} can be modified,it does nothing to do with the Permutator

Invoke Permutator.Left() to return the number of ungenerated permutation

	func (p Permutator) Left() int

Invoke Permutator.Index() to return the index of next permutation, which start from 1 to factorial(length of slice)

	func (p Permutator) Index() int

An example:

	func main() {
		i := []int{4,3,2,1}
		p,err:=NewPerm(i,nil) //generate a Permutator
		if err != nil {
			fmt.Println(err)
			return
		}
		for i,err:=p.Next();err==nil;i,err=p.Next(){
			fmt.Printf("%3d permutation: %v left %d\n",p.Index()-1,i.([]int),p.Left())
		}
	}
--------------------------------------
outputs:

	  1 permutation: [1 2 3] left 5
	  2 permutation: [1 3 2] left 4
	  3 permutation: [2 1 3] left 3
	  4 permutation: [2 3 1] left 2
	  5 permutation: [3 1 2] left 1
	  6 permutation: [3 2 1] left 0
