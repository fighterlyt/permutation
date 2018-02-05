package permutation

import (
	"errors"
	"reflect"
	"sort"
	"sync"
)

var (
	NotASliceError         = errors.New("argument must be a slice")
	InvalidCollectionError = errors.New("argument must not be nil")
	EmptyCollectionError   = errors.New("argument must not be empty")
)

type sortable struct {
	value reflect.Value
	less  Less
}

func (s sortable) Len() int {
	return s.value.Len()
}
func (s sortable) Less(i, j int) bool {
	return s.less(s.value.Index(i).Interface(), s.value.Index(j).Interface())
}
func (s sortable) Swap(i, j int) {
	temp := reflect.ValueOf(s.value.Index(i).Interface())
	s.value.Index(i).Set(s.value.Index(j))
	s.value.Index(j).Set(temp)
}

// Permutator is a class that one can itterate through
// in order to get the sucessive permutations of the set
type Permutator struct {
	sync.Mutex
	value  reflect.Value
	less   Less
	length int
	index  int
	amount int
}

//Reset the Permutator, next time invoke p.Next() will return the first permutation in lexicalorder
func (p *Permutator) Reset() {
	p.Lock()
	defer p.Unlock()

	sort.Sort(sortable{p.value, p.less})
	p.index = 1
}

// NextN returns the next n permuations, if n>p.Left(),return all the left permuations
// if all permutaions generated or n is illegal(n<=0),return a empty slice
func (p *Permutator) NextN(n int) interface{} {
	p.Lock()
	defer p.Unlock()
	if n <= 0 || p.left() == 0 {
		return reflect.MakeSlice(reflect.SliceOf(p.value.Type()), 0, 0).Interface()
	}

	cap := p.left()
	if cap > n {
		cap = n
	}

	result := reflect.MakeSlice(reflect.SliceOf(p.value.Type()), cap, cap)

	length := 0
	for index := 0; index < cap; index++ {
		p.Unlock()
		if _, err := p.Next(); err == nil {
			length++
			list := p.copySliceValue()
			result.Index(index).Set(list)
		}
		p.Lock()
	}

	list := reflect.MakeSlice(result.Type(), length, length)
	reflect.Copy(list, result)

	return list.Interface()
}

// Index returns the index of last permutation, which start from 1 to n! (n is the length of slice)
func (p *Permutator) Index() int {
	p.Lock()
	defer p.Unlock()

	j := p.index - 1
	return j
}

// ErrUnordered occurs when you have a slice in an unordered state
var ErrUnordered = errors.New("the element type of slice is not ordered, you must provide a function")

// NewPerm generate a New Permuatator,
// the argument k must be a non-nil slice,
// and the less argument must be a Less function that implements compare functionality of k's element type
// if k's element is ordered,less argument can be nil
// for ordered in Golang, visit http://golang.org/ref/spec#Comparison_operators
// After generating a Permutator, the argument k can be modified and deleted,Permutator store a copy of k internel.Rght now, a Permutator can  be used concurrently
func NewPerm(k interface{}, less Less) (*Permutator, error) {
	value := reflect.ValueOf(k)

	//check to see if i is a slice
	if value.Kind() != reflect.Slice {
		return nil, NotASliceError
	}

	if value.IsValid() != true {
		return nil, InvalidCollectionError
	}

	if value.Len() == 0 {
		return nil, EmptyCollectionError
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

	s := &Permutator{value: value, less: less, length: length, index: 1, amount: factorial(length)}

	return s, nil
}

//Next the next permuation in lexcial order,if all permutations generated,return an error
func (p *Permutator) Next() (interface{}, error) {
	p.Lock()
	defer p.Unlock()
	//check to see if all permutations generated
	if p.left() <= 0 {
		return nil, errors.New("all Permutations generated")
	}

	var i, j int
	//the first permuation is just p.value
	if p.index == 1 {
		p.index++
		l := reflect.MakeSlice(p.value.Type(), p.length, p.length)
		reflect.Copy(l, p.value)
		return l.Interface(), nil
	}

	//when we arrive here, there must be some permutations to generate

	for i = p.length - 2; i > 0; i-- {
		if p.less(p.value.Index(i).Interface(), p.value.Index(i+1).Interface()) {
			break
		}
	}
	for j = p.length - 1; j > 0; j-- {
		if p.less(p.value.Index(i).Interface(), p.value.Index(j).Interface()) {
			break
		}
	}
	//swap
	temp := reflect.ValueOf(p.value.Index(i).Interface())
	p.value.Index(i).Set(p.value.Index(j))
	p.value.Index(j).Set(temp)
	//reverse
	reverse(p.value, i+1, p.length-1)

	//increase the counter
	p.index++
	l := reflect.MakeSlice(p.value.Type(), p.length, p.length)
	reflect.Copy(l, p.value)
	return l.Interface(), nil
}

//Left returns the left permutation that can be generated
func (p *Permutator) Left() int {
	p.Lock()
	defer p.Unlock()
	j := p.left()
	return j
}
func (p *Permutator) copySliceValue() reflect.Value {
	list := reflect.MakeSlice(p.value.Type(), p.length, p.length)
	reflect.Copy(list, p.value)
	return list
}

//because we use left inside some methods,so we need a non-block version
func (p *Permutator) left() int {
	return p.amount - p.index + 1
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
		if !less(value.Index(index).Interface(), value.Index(index+1).Interface()) {
			break
		}
	}

	if index != lastIndex {
		sort.Sort(sortable{value, less})
	}
}

//reverse the slice v[i:j]
func reverse(v reflect.Value, i, j int) {
	length := j - i + 1

	if length < 2 {
		return
	}

	for length >= 2 {
		temp := reflect.ValueOf(v.Index(j).Interface())
		v.Index(j).Set(v.Index(i))
		v.Index(i).Set(temp)

		length -= 2
		i++
		j--
	}
}
