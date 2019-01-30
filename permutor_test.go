package permutation

import "testing"

func TestNewPermErrors(t *testing.T) {
	var tests = []struct {
		Name   string
		K      interface{}
		Expect error
	}{
		{
			Name:   "String",
			K:      "not a slice",
			Expect: NotASliceError,
		},
		{
			Name:   "Int",
			K:      1735,
			Expect: NotASliceError,
		},
		{
			Name:   "Struct",
			K:      struct{ Data int }{Data: 21},
			Expect: NotASliceError,
		},
		{
			Name:   "Empty",
			K:      []int{},
			Expect: EmptyCollectionError,
		},
	}

	for _, test := range tests {
		p, err := NewPerm(test.K, nil)
		if p != nil {
			t.Errorf("%s: got non-nil permutator pointer with k=%v\n", test.Name, test.K)
		}
		if err != test.Expect {
			t.Errorf("%s: got wrong error type '%s' from k=%v\n", test.Name, err, test.K)
		}
	}
}

func TestValueIsCopy(t *testing.T) {
	var orig = []int{1, 2, 3, 4}
	p, err := NewPerm(orig, nil)
	if err != nil {
		t.Errorf("Error when creating permutator: '%s'\n", err)
	}

	// the elements of orig and p.value should be the same
	for i, elem := range orig {
		if p.value.Index(i).Int() != int64(elem) {
			t.Errorf("Element mismatch: orig[%d] = %d; p.value[%d] = %d\n",
				i, orig[i], i, p.value.Index(i).Int())
		}
	}

	for i := 0; i < len(orig); i++ {
		orig[i] = orig[i] * 2
	}

	// not the elements of orig should all be twice the corresponding elements of p.value
	for i, elem := range orig {
		if p.value.Index(i).Int() != int64(elem/2) {
			t.Errorf("Element mismatch: orig[%d] = %d; p.value[%d] = %d\n",
				i, orig[i], i, p.value.Index(i).Int())
		}
	}
}
