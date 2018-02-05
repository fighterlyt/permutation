package permutation

import (
	"log"
	"reflect"
)

type Useable interface {
	Equal(Useable) bool
}

func calcFactorial(in int) int {
	// n! = n * (n-1)!
	if in > 1 {
		return calcFactorial(in-1) * in
	} else if in == 1 {
		return 1
	}
	return 0
}
func checkSliceInt(e, j [][]int) {
	for i, v := range j {
		if !equalSliceInt(e[i], v) {
			log.Fatal("array:", e[i], "!=", v)
		}
	}
}
func equalSliceInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func equalSliceString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalSliceGen(a, b interface{}, piv int) bool {
	switch b.(type) {
	case []int:
		refArr := a.([][]int)
		resArr := b.([]int)
		if !equalSliceInt(refArr[piv], resArr) {
			log.Fatal("array:", refArr[piv], "!=", resArr)
		}
	case []string:
		refArr := a.([][]string)
		resArr := b.([]string)
		if !equalSliceString(refArr[piv], resArr) {
			log.Fatal(piv, "String array:", refArr[piv], "!=", resArr)
		}

	case []Useable:
		refArr := a.([][]Useable)
		resArr := b.([]Useable)
		for i, w := range resArr {
			if !refArr[piv][i].Equal(w) {
				log.Fatal(i, "element failed at ", piv)
			}
		}

	default:
		log.Fatalf("Unknown type %T, %t", a)

	}
	return true
}

// leng is the length of outer array
// sz if the size of inner element
func equalSliceSliceGen(a, b interface{}) (good bool, leng, sz int) {
	switch b.(type) {
	case [][]int:
		refArr := a.([][]int)
		resArr := b.([][]int)
		leng = len(resArr)
		sz = len(resArr[0])
		for piv := range resArr {
			if !equalSliceInt(refArr[piv], resArr[piv]) {
				log.Fatal("array:", refArr[piv], "!=", resArr)
			}
		}
		good = true
	case [][]string:
		refArr := a.([][]string)
		resArr := b.([][]string)
		leng = len(resArr)
		sz = len(resArr[0])
		for piv := range resArr {
			if !equalSliceString(refArr[piv], resArr[piv]) {
				log.Fatal(piv, "String array:", refArr[piv], "!=", resArr)
			}
		}
		good = true
	default:
		_, ok := a.([][]Useable)
		if ok {
			refArr := a.([][]Useable)
			resArr := b.([][]Useable)
			leng = len(resArr)
			sz = len(resArr[0])
			for piv := range resArr {
				for i, w := range resArr[piv] {
					if !refArr[piv][i].Equal(w) {
						log.Fatal(i, "element failed at ", piv)
					}
				}
			}
			good = true
		} else {
			val := reflect.ValueOf(b)
			if val.Kind() == reflect.Slice {
				leng = val.Len() // outer length
				for i := 0; i < leng; i++ {
					v := val.Index(i)
					if v.Kind() == reflect.Slice {
						sz = v.Len() // inner length
						for j := 0; j < sz; j++ {
							vv := v.Index(j)
							// b implements Usable
							if m, ok := vv.Interface().(Useable); ok {
								n, ok := reflect.ValueOf(a).Index(i).Index(j).Interface().(Useable)
								if !ok {
									log.Fatal("a is not the same type as b")
								}
								good = m.Equal(n)
							} else {
								log.Fatal("Candidate cannot be converted to a Usable type")
							}
						}
					}
				}
			}
		}
	}
	//good = true
	return
}
