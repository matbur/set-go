package set

import (
	"testing"
	"reflect"
	"fmt"
)

func TestNew(t *testing.T) {
	for _, v := range []struct {
		values []int
		set    map[int]struct{}
	}{
		{[]int{}, map[int]struct{}{}},
		{[]int{1}, map[int]struct{}{1: {}}},
		{[]int{1, -2}, map[int]struct{}{1: {}, -2: {}}},
		{[]int{1, 1}, map[int]struct{}{1: {}}},
		{[]int{-1, 2, -1}, map[int]struct{}{-1: {}, 2: {}}},
	} {
		set, want := v.values, v.set
		s := New(set...)
		get := s.set
		if !reflect.DeepEqual(get, want) {
			t.Errorf("New(%v).set == %v, want %v", set, get, want)
		}
	}
}

func TestSet_Len(t *testing.T) {
	for _, v := range []struct {
		set []int
		l   int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, -2}, 2},
		{[]int{1, 1}, 1},
		{[]int{-1, 2, -1}, 2},
	} {
		set, want := v.set, v.l
		s := New(set...)
		get := s.Len()
		if get != want {
			t.Errorf("New(%v).Len() == %v, want %v", set, get, want)
		}
	}
}

func TestSet_Clear(t *testing.T) {
	for _, v := range []struct {
		set []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, -2}},
		{[]int{1, 1}},
		{[]int{-1, 2, -1}},
	} {
		set := v.set
		want := make(map[int]struct{})
		s := New(set...)
		get := s.Clear().set
		if !reflect.DeepEqual(get, want) {
			t.Errorf("New(%v).Clear().set == %v, want %v", set, get, want)
		}
	}
}

func TestSet_Copy(t *testing.T) {
	for _, v := range []struct {
		set []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, -2}},
		{[]int{1, 1}},
		{[]int{-1, 2, -1}},
	} {
		set := v.set
		s1 := New(set...)
		s2 := s1.Copy()
		addrS1 := fmt.Sprintf("%p", s1.set)
		addrS2 := fmt.Sprintf("%p", s2.set)
		if !reflect.DeepEqual(s1.set, s2.set) {
			t.Errorf("New(%v).Copy() = %v, want %v", set, s2.set, s1.set)
		}
		if addrS1 == addrS2 {
			t.Errorf("addr of New(%v).Copy() == %v is the same as original", set, addrS1)
		}
	}
}

func TestSet_String(t *testing.T) {
	for _, v := range []struct {
		set []int
		str string
	}{
		{[]int{}, "{}"},
		{[]int{1}, "{1}"},
		{[]int{1, -2}, "{-2, 1}"},
		{[]int{1, 1}, "{1}"},
		{[]int{-1, 2, -1}, "{-1, 2}"},
		{[]int{10, 2, -1}, "{-1, 2, 10}"},
		{[]int{10, 2, -1, -10}, "{-10, -1, 2, 10}"},
	} {
		set, want := v.set, v.str
		s := New(set...)
		get := s.String()
		if get != want {
			t.Errorf("New(%v).String() == %v, want %v", set, get, want)
		}
	}
}

func TestSet_Add(t *testing.T) {
	for _, v := range []struct {
		set   []int
		value int
		str   string
	}{
		{[]int{}, 7, "{7}"},
		{[]int{1}, 7, "{1, 7}"},
		{[]int{1, -2}, 1, "{-2, 1}"},
		{[]int{10, 2, -1}, 7, "{-1, 2, 7, 10}"},
		{[]int{10, 2, -1, -10}, 2, "{-10, -1, 2, 10}"},
	} {
		set, value, want := v.set, v.value, v.str
		s := New(set...)
		get := s.Copy().Add(value).String()
		if get != want {
			t.Errorf("New(%v).Add(%v) == %v, want %v", set, value, get, want)
		}
	}
}

func TestSet_IsIn(t *testing.T) {}

func TestSet_Remove(t *testing.T) {}

func TestSet_Pop(t *testing.T) {}

func TestSet_Difference(t *testing.T) {}

func TestSet_Intersection(t *testing.T) {}

func TestSet_SymmetricDifference(t *testing.T) {}

func TestSet_Union(t *testing.T) {}

func TestSet_IsDisjoint(t *testing.T) {}

func TestSet_IsSubset(t *testing.T) {}

func TestSet_IsSuperset(t *testing.T) {}
