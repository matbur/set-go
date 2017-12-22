package set

import (
	"testing"
	"reflect"
	"fmt"
	"errors"
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

func TestSet_Equal(t *testing.T) {
	for _, v := range []struct {
		set1    []int
		set2    []int
		isEqual bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1}, []int{7}, false},
		{[]int{1, -2}, []int{1}, false},
		{[]int{10, 2, -1}, []int{-1, 7}, false},
		{[]int{2, -10}, []int{-10, 2}, true},
	} {
		set1, set2, want := v.set1, v.set2, v.isEqual
		s := New(set1...)
		get := s.Equal(New(set2...))
		if get != want {
			t.Errorf("New(%v).Equal(New(%v)) == %v, want %v", set1, set2, get, want)
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

func TestSet_IsIn(t *testing.T) {
	for _, v := range []struct {
		set   []int
		value int
		isIn  bool
	}{
		{[]int{}, 7, false},
		{[]int{1}, 7, false},
		{[]int{1, -2}, 1, true},
		{[]int{10, 2, -1}, 7, false},
		{[]int{10, 2, -1, -10}, 2, true},
	} {
		set, value, want := v.set, v.value, v.isIn
		s := New(set...)
		get := s.IsIn(value)
		if get != want {
			t.Errorf("New(%v).IsIn(%v) == %v, want %v", set, value, get, want)
		}
	}
}

func TestSet_Remove(t *testing.T) {
	err := "no such value in set"
	for _, v := range []struct {
		set   []int
		value int
		err   error
	}{
		{[]int{}, 7, errors.New(err)},
		{[]int{1}, 7, errors.New(err)},
		{[]int{1, -2}, 1, nil},
		{[]int{10, 2, -1}, 7, errors.New(err)},
		{[]int{10, 2, -1, -10}, 2, nil},
	} {
		set, value, want := v.set, v.value, v.err
		s := New(set...)
		get := s.Remove(value)
		if (get == nil || want == nil ) && get != want ||
			get != nil && want != nil && get.Error() != want.Error() {
			t.Errorf(`New(%v).Remove(%v) == "%v", want "%v"`, set, value, get, want)
		}
	}
}

func TestSet_Pop(t *testing.T) {
	err := "the map is empty"
	for _, v := range []struct {
		set []int
		err error
	}{
		{[]int{}, errors.New(err)},
		{[]int{1}, nil},
		{[]int{1, -2}, nil},
		{[]int{10, 2, -1}, nil},
	} {
		set, want := v.set, v.err
		s := New(set...)
		v, get := s.Pop()
		if (get == nil || want == nil ) && get != want ||
			get != nil && want != nil && get.Error() != want.Error() {
			t.Errorf(`New(%v).Pop() == (%v, "%v"), want (x, "%v")`, set, v, get, want)
		}
	}
}

func TestSet_Difference(t *testing.T) {
	for _, v := range []struct {
		set1 []int
		set2 []int
		str  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{1}, "{}"},
		{[]int{1}, []int{7}, "{1}"},
		{[]int{1, -2}, []int{1}, "{-2}"},
		{[]int{10, 2, -1}, []int{-1, 7}, "{2, 10}"},
		{[]int{10, 2, -1, -10}, []int{-10, 2, 3}, "{-1, 10}"},
	} {
		set1, set2, want := v.set1, v.set2, v.str
		s := New(set1...)
		get := s.Difference(New(set2...)).String()
		if get != want {
			t.Errorf("New(%v).Difference(New(%v)) == %v, want %v", set1, set2, get, want)
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	for _, v := range []struct {
		set1 []int
		set2 []int
		str  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{1}, "{1}"},
		{[]int{1}, []int{7}, "{}"},
		{[]int{1, -2}, []int{1}, "{1}"},
		{[]int{10, 2, -1}, []int{-1, 7}, "{-1}"},
		{[]int{10, 2, -1, -10}, []int{-10, 2, 3}, "{-10, 2}"},
	} {
		set1, set2, want := v.set1, v.set2, v.str
		s := New(set1...)
		get := s.Intersection(New(set2...)).String()
		if get != want {
			t.Errorf("New(%v).Intersection(New(%v)) == %v, want %v", set1, set2, get, want)
		}
	}
}

func TestSet_Union(t *testing.T) {
	for _, v := range []struct {
		set1 []int
		set2 []int
		str  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{1}, "{1}"},
		{[]int{1}, []int{7}, "{1, 7}"},
		{[]int{1, -2}, []int{1}, "{-2, 1}"},
		{[]int{10, 2, -1}, []int{-1, 7}, "{-1, 2, 7, 10}"},
		{[]int{10, 2, -1, -10}, []int{-10, 2, 3}, "{-10, -1, 2, 3, 10}"},
	} {
		set1, set2, want := v.set1, v.set2, v.str
		s := New(set1...)
		get := s.Union(New(set2...)).String()
		if get != want {
			t.Errorf("New(%v).Union(New(%v)) == %v, want %v", set1, set2, get, want)
		}
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	for _, v := range []struct {
		set1 []int
		set2 []int
		str  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{1}, "{}"},
		{[]int{1}, []int{7}, "{1, 7}"},
		{[]int{1, -2}, []int{1}, "{-2}"},
		{[]int{10, 2, -1}, []int{-1, 7}, "{2, 7, 10}"},
		{[]int{10, 2, -1, -10}, []int{-10, 2, 3}, "{-1, 3, 10}"},
	} {
		set1, set2, want := v.set1, v.set2, v.str
		s := New(set1...)
		get := s.SymmetricDifference(New(set2...)).String()
		if get != want {
			t.Errorf("New(%v).SymmetricDifference(New(%v)) == %v, want %v", set1, set2, get, want)
		}
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	for _, v := range []struct {
		set1       []int
		set2       []int
		isDisjoint bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, false},
		{[]int{1}, []int{7}, true},
		{[]int{1, -2}, []int{1}, false},
		{[]int{10, 2, -1}, []int{1, 7}, true},
		{[]int{10, 2, -1, -10}, []int{-10, 2, 3}, false},
	} {
		set1, set2, want := v.set1, v.set2, v.isDisjoint
		s := New(set1...)
		get := s.IsDisjoint(New(set2...))
		if get != want {
			t.Errorf("New(%v).SymmetricDifference(New(%v)) == %v, want %v", set1, set2, get, want)
		}
	}
}

func TestSet_IsSubset(t *testing.T) {
	for _, v := range []struct {
		set1     []int
		set2     []int
		isSubset bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1}, []int{7}, false},
		{[]int{1, -2}, []int{1}, false},
		{[]int{10, 2, -1}, []int{1, 7}, false},
		{[]int{-10, 2}, []int{10, 2, -1, -10}, true},
	} {
		set1, set2, want := v.set1, v.set2, v.isSubset
		s := New(set1...)
		get := s.IsSubset(New(set2...))
		if get != want {
			t.Errorf("New(%v).IsSubset(New(%v)) == %v, want %v", set1, set2, get, want)
		}
	}
}

func TestSet_IsSuperset(t *testing.T) {}
