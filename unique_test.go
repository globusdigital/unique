package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	for i, tc := range []struct {
		data, want Interface
	}{
		{
			IntSlice{&[]int{}},
			IntSlice{&[]int{}},
		},
		{
			IntSlice{&[]int{1}},
			IntSlice{&[]int{1}},
		},
		{
			IntSlice{&[]int{1, 2}},
			IntSlice{&[]int{1, 2}},
		},
		{
			IntSlice{&[]int{1, 1}},
			IntSlice{&[]int{1}},
		},
		{
			IntSlice{&[]int{1, 2, 2}},
			IntSlice{&[]int{1, 2}},
		},
		{
			IntSlice{&[]int{1, 1, 2}},
			IntSlice{&[]int{1, 2}},
		},
		{
			GenericSlice[uint]{&[]uint{1, 1, 2, 3, 3}},
			GenericSlice[uint]{&[]uint{1, 2, 3}},
		},
	} {
		Unique(tc.data)
		if !reflect.DeepEqual(tc.data, tc.want) {
			t.Errorf("%d: got %#v; want %#v", i, tc.data, tc.want)
		}
	}
}

func TestGenericsAreUnique(t *testing.T) {
	have := GenericSlice[int8]{&[]int8{5, 4, 1, 1, -1, 2, 3, 3}}
	want := GenericSlice[int8]{&[]int8{-1, 1, 2, 3, 4, 5}}
	Sort(have)
	if !reflect.DeepEqual(have, want) {
		t.Errorf("got %#v; want %#v", want, want)
	}
}
