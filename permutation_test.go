package permutation

import (
	"log"
	"math"
	"reflect"
	"testing"
)

func TestPermBasicInt(t *testing.T) {
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
	runTestGeneric(a, e, nil)
}
func TestPermCustomLess(t *testing.T) {
	a := []int{1, 2, 3}
	e := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	l := func(i, j interface{}) bool {
		return i.(int) < j.(int)
	}
	runTestGeneric(a, e, l)
	log.Println("Forward less worked")
	a = []int{3, 2, 1}
	e = [][]int{
		{3, 2, 1},
		{3, 1, 2},
		{2, 3, 1},
		{2, 1, 3},
		{1, 3, 2},
		{1, 2, 3},
	}
	l = func(i, j interface{}) bool {
		return j.(int) < i.(int)
	}
	runTestGeneric(a, e, l)
	log.Println("Backwards less worked")
}
func TestPermDuplicates(t *testing.T) {
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
	runTestGeneric(a, e, nil)
}
func TestPermString0(t *testing.T) {
	a := []string{"one", "Two"}
	e := [][]string{
		{"Two", "one"},
		{"one", "Two"},
	}
	runTestGeneric(a, e, nil)
}
func TestPermString1(t *testing.T) {
	a := []string{"one", "Two", "three"}
	e := [][]string{
		{"Two", "one", "three"},
		{"Two", "three", "one"},
		{"one", "Two", "three"},
		{"one", "three", "Two"},
		{"three", "Two", "one"},
		{"three", "one", "Two"},
	}
	runTestGeneric(a, e, nil)
}

type tmpType int

func (tt tmpType) Equal(nt Useable) bool {
	nti := nt.(tmpType)
	return tt == nti
}
func TestPermCustomType0(t *testing.T) {
	var bob tmpType
	var steve Useable
	bob = tmpType(1)
	steve = Useable(bob)
	if !steve.Equal(bob) {
		log.Fatal("Type problem")
	}
	a := []tmpType{1, 2}
	e := [][]tmpType{
		{1, 2},
		{2, 1},
	}
	runTestGeneric(a, e, nil)
}

func runTestGeneric(a interface{}, e interface{}, l Less) {
	p, err := NewPerm(a, l)
	if err != nil {
		log.Fatal(err)
	}
	for result, err := p.Next(); err == nil; result, err = p.Next() {
		if true {
			log.Println(result)
		} else {
			equalSliceGen(e, result, p.Index()-1)
		}
	}

	p.Reset()
	result := p.NextN(30)
	//checkSliceInt(e, j)

	_, leng, sz := equalSliceSliceGen(e, result)

	if factorial(sz) != p.Index() {
		log.Fatal("Length is wrong", leng, p.Index())
	}

	p.Reset()
	result = p.NextN(1)
	_, leng, _ = equalSliceSliceGen(e, result)
	if leng != 1 {
		log.Fatal("p.NextN not returned expected length", leng, result)
	}
}

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

func Test_MoveIndex(t *testing.T) {
	testintdata := []int{1, 2, 3, 4}

	p, err := NewPerm(testintdata, nil)
	if err != nil {
		t.Errorf("Error creating permutator: '%s'\n", err)
	}

	newindex, err := p.MoveIndex(2)
	if err != nil {
		t.Errorf("Error moving index: '%s'\n", err)
	}
	if newindex != 2 {
		t.Errorf("Expected indext 2, is: %d\n", newindex)
	}

	// Test that an error occurs if we go beyond the end of the index
	newindex, err = p.MoveIndex(p.Amount())
	if err != nil {
		t.Error("Error moving index to last permutation")
	}
	if newindex != p.Amount() {
		t.Errorf("Expected index %d, was: %d\n", p.Amount(), newindex)
	}

	// Test that an error occurs if we go beyond the end of the index
	oldidx := p.Index()
	newindex, err = p.MoveIndex(p.Amount() + 1)
	if err == nil {
		t.Error("Expected an error moving index beyond end of permutations")
	}
	if oldidx != newindex {
		t.Errorf("Index should not have moved from %d, to: %d\n", oldidx, newindex)
	}

	// Test that an error occurs if we specify a negative index
	newindex, err = p.MoveIndex(-1)
	if err == nil {
		t.Error("Expected an error moving index negative")
	}
	if oldidx != newindex {
		t.Errorf("Index should not have moved from %d, to: %d\n", oldidx, newindex)
	}
}

func Test_Amount(t *testing.T) {
	testintdata := []int{1, 2}
	p, err := NewPerm(testintdata, nil)
	if err != nil {
		t.Errorf("Error creating permutator: '%s'\n", err)
	}

	if p.Amount() != 2 { // 2!
		t.Errorf("Incorrect value for Amount. Expected 2 was %d'\n", p.Amount())
	}

	testintdata = []int{1, 2, 3, 4}
	p, err = NewPerm(testintdata, nil)
	if err != nil {
		t.Errorf("Error creating permutator: '%s'\n", err)
	}

	if p.Amount() != 24 { // 4!
		t.Errorf("Incorrect value for Amount. Expected 24 was %d'\n", p.Amount())
	}

	testintdata = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p, err = NewPerm(testintdata, nil)
	if err != nil {
		t.Errorf("Error creating permutator: '%s'\n", err)
	}

	if p.Amount() != 3628800 { // 10!
		t.Errorf("Incorrect value for Amount. Expected 3628800 was %d'\n", p.Amount())
	}
}
