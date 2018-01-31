package permutation

import "log"

type usable interface {
	Equal(usable) bool
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

	case []usable:
		refArr := a.([][]usable)
		resArr := b.([]usable)
		for i, w := range resArr {
			if !refArr[piv][i].Equal(w) {
				log.Fatal(i, "element failed at ", piv)
			}
		}

	default:
		log.Fatalf("Unknown type %T\n", a)

	}
	return true
}
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

	case [][]usable:
		refArr := a.([][]usable)
		resArr := b.([][]usable)
		leng = len(resArr)
		sz = len(resArr[0])
		for piv := range resArr {
			for i, w := range resArr[piv] {
				if !refArr[piv][i].Equal(w) {
					log.Fatal(i, "element failed at ", piv)
				}
			}
		}

	default:
		log.Fatalf("Unknown type %T\n", a)

	}
	good = true
	return
}
