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
			t.Errorf("NewSet(%v).set == %v, want %v", set, get, want)
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
			t.Errorf("NewSet(%v).Len() == %v, want %v", set, get, want)
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
			t.Errorf("NewSet(%v).Clear().set == %v, want %v", set, get, want)
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
			t.Errorf("NewSet(%v).Copy() = %v, want %v", set, s2.set, s1.set)
		}
		if addrS1 == addrS2 {
			t.Errorf("addr of NewSet(%v).Copy() == %v is the same as original", set, addrS1)
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
			t.Errorf("NewSet(%v).String() == %v, want %v", set, get, want)
		}
	}
}

func TestSet_Add(t *testing.T) {}

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
