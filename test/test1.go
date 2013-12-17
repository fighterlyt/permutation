package main
import (
	"../../permutation"
	"fmt"
)
func main(){
	a:=[]int{1,2,3,4}

	p,err:=permutation.NewPerm(a,nil)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for result,err:=p.Next();err==nil;result,err=p.Next(){
		fmt.Println(p.Index()-1,result.([]int))
	}
	p.Reset()

	result:=p.NextN(30)
	j:=result.([][]int)
	for _,i:=range j{
		fmt.Println(i)
	}
	fmt.Println(p.Index()-1)
}