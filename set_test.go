package set

import (
	"testing"
	"reflect"
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

func TestSet_String(t *testing.T) {}

func TestSet_Len(t *testing.T) {}

func TestSet_Clear(t *testing.T) {}

func TestSet_Copy(t *testing.T) {}

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
