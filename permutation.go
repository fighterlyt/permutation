package permutation

import (
    "sort"
	"errors"
	"reflect"
)

type sortable struct {
	value reflect.Value
	less  Less
}

type Permutator struct {
	idle   chan bool
	value  reflect.Value
	less   Less
	length int
	index  int
	amount int
}

// Generate a New Permuatator, the argument k must be a non-nil slice,and the less argument must be a Less function that implements compare functionality of k's element type
// if k's element is ordered,less argument can be nil
// for ordered in Golang, visit http://golang.org/ref/spec#Comparison_operators
// After generating a Permutator, the argument k can be modified and deleted,Permutator store a copy of k internel.Rght now, a Permutator can  be used concurrently
func NewPerm(input interface{}, less Less) (*Permutator, error) {
	value := reflect.ValueOf(input)
    
	//check to see if i is a slice
	if value.Kind() != reflect.Slice {
		return nil, errors.New("argument must be a slice")
	}
    
	if value.IsValid() != true {
		return nil, errors.New("argument must not be nil")
	}
    
	if value.Len() == 0 {
		return nil, errors.New("argument must not be empty")
	}

	l := reflect.MakeSlice(value.Type(), value.Len(), value.Len())
	
    reflect.Copy(l, value)
    
	value = l
    
	length := value.Len()
    
	if less == nil {
        lessType, err := getLessFunctionByValueType(value)
        if err != nil {
            return nil, err
        }
        less = lessType
	}
    
    sortValues(value, less)

	s := &Permutator { value: value, less: less, length: length, index: 1, amount: factorial(length) }
	s.idle = make(chan bool, 1)
	s.idle <- true
    
	return s, nil
}

func (s sortable) Len() int {
	return s.value.Len()
}

func (s sortable) Less(left, right int) bool {
	return s.less(s.value.Index(left).Interface(), s.value.Index(right).Interface())
}

func (s sortable) Swap(left, right int) {
	value := reflect.ValueOf(s.value.Index(left).Interface())
	s.value.Index(left).Set(s.value.Index(right))
	s.value.Index(right).Set(value)
}

// Reset the Permutator so the next invocation of p.Next() will return the first permutation in lexical order
func (p *Permutator) Reset() {
	<- p.idle
	sort.Sort(sortable{ p.value, p.less })
	p.index = 1
	p.idle <- true
}

// Index will return the index of last permutation, which starts from 1 to n! (n is the length of slice)
func (p Permutator) Index() int {
	<- p.idle
	j := p.index - 1
	p.idle <- true
	return j
}

// return the next n permuations, if n>p.Left(),return all the left permuations
// if all permutaions generated or n is illegal(n<=0),return a empty slice
func (p *Permutator) NextN(count int) interface{} {    
	if count <= 0 || p.left() == 0 {
		return reflect.MakeSlice(reflect.SliceOf(p.value.Type()), 0, 0).Interface()
	}
    
    cap := p.left()
	if cap > count {
		cap = count
	}

    result := reflect.MakeSlice(reflect.SliceOf(p.value.Type()), cap, cap)

    length := 0    
    for index := 0; index < cap; index++ {        
        if _, ok := p.Next(); ok {
            length++
            list := p.copySliceValue()
            result.Index(index).Set(list)
        }
    }

    list := reflect.MakeSlice(result.Type(), length, length)
    reflect.Copy(list, result)
     
    return list.Interface()
}

// Next generates the next permuation in lexcial order. If all permutations were generated an error is returned.
func (p *Permutator) Next() (interface{}, bool) {
	<- p.idle

	// check to see if all permutations generated
	if p.left() <= 0 {
		p.idle <- true
		return nil, false
	}

	// the first permuation is just p.value
	if p.index == 1 {
        p.index++        
        list := p.copySliceValue()   
        p.idle <- true        
        return list.Interface(), true
	}

    var left, right int
	for left = p.length - 2; left >= 0; left-- {
		if p.less(p.value.Index(left).Interface(), p.value.Index(left + 1).Interface()) {
			break
		}
	}
    
    if left == -1 {
        p.idle <- true
        return nil, false
    }
    
	for right = p.length - 1; right >= 0; right-- {
		if p.less(p.value.Index(left).Interface(), p.value.Index(right).Interface()) {
			break
		}
	}
    
    p.swap(left, right)    
    
    left++    
	
    p.reverse(left, p.length - 1)
    p.index++    
    list := p.copySliceValue()	
    
    p.idle <- true
        
	return list.Interface(), true
}

// Left returns the left permutation that can be generated
func (p Permutator) Left() int {
	<- p.idle
	remaining := p.left()
	p.idle <- true
	return remaining
}

func (p *Permutator) copySliceValue() reflect.Value {
    list := reflect.MakeSlice(p.value.Type(), p.length, p.length)
    reflect.Copy(list, p.value)
    return list
}

// because we use left inside some methods we need a non-blocking version
func (p Permutator) left() int {
	return (p.amount - p.index) + 1
}

func (p *Permutator) swap(left, right int) {
    value := reflect.ValueOf(p.value.Index(right).Interface())
	p.value.Index(right).Set(p.value.Index(left))
	p.value.Index(left).Set(value)
}

func sortValues(value reflect.Value, less Less) {
	index := 0
    lastIndex := value.Len() - 1
	for index = 0; index < lastIndex; index++ {
		if !less(value.Index(index).Interface(), value.Index(index + 1).Interface()) {
			break
		}
	}
    
	if index != lastIndex {
		sort.Sort(sortable { value, less })
	}
}

func (p *Permutator) reverse(left, right int) {   
	length := (right - left) + 1
	if length < 2 {
		return
	}

	for length >= 2 {
        length -= 2
        p.swap(left, right)
		left++
		right--
	}
}

func factorial(i int) int {
	result := 1
	for i > 0 {
		result *= i
		i--
	}
	return result
}