// Package set provides implementation of unordered collection of unique values.
package set

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var (
	ErrValueNotFound = errors.New("no such value in set")
	ErrEmptySet      = errors.New("the map is empty")
)

type Set struct {
	set map[int]struct{}
}

// New builds a new set.
func New(values ...int) *Set {
	set := make(map[int]struct{})
	for _, v := range values {
		set[v] = struct{}{}
	}
	s := Set{set}
	return &s
}

// Len gets number of items in set.
func (s *Set) Len() int {
	return len(s.set)
}

// Empty checks if set contains no elements.
func (s *Set) Empty() bool {
	return s.Len() == 0
}

// Clear gets rid of all set from set.
func (s *Set) Clear() *Set {
	s.set = make(map[int]struct{})
	return s
}

// Copy creates a new set with the same elements.
func (s *Set) Copy() *Set {
	cp := New()
	for i := range s.set {
		cp.Add(i)
	}
	return cp
}

// Equal checks if two sets contain the same elements.
func (s *Set) Equal(other *Set) bool {
	isEqual := reflect.DeepEqual(s.set, other.set)
	return isEqual
}

// String converts set to a string.
func (s *Set) String() string {
	var tabInt []int
	for i := range s.set {
		tabInt = append(tabInt, i)
	}
	sort.Ints(tabInt)
	var tabStr []string
	for _, i := range tabInt {
		tabStr = append(tabStr, fmt.Sprint(i))
	}
	inner := strings.Join(tabStr, ", ")
	str := fmt.Sprintf("{%s}", inner)
	return str
}

// Add inserts new value to a set.
func (s *Set) Add(value int) *Set {
	s.set[value] = struct{}{}
	return s
}

// IsIn checks if set contains value.
func (s *Set) IsIn(value int) bool {
	_, ok := s.set[value]
	return ok
}

// Remove deletes a value from a set.
func (s *Set) Remove(value int) error {
	if !s.IsIn(value) {
		return ErrValueNotFound
	}
	delete(s.set, value)
	return nil
}

// Pop removes and returns some element from a set.
func (s *Set) Pop() (int, error) {
	if s.Len() == 0 {
		return 0, ErrEmptySet
	}
	var value int
	for i := range s.set {
		value = i
		break
	}
	s.Remove(value)
	return value, nil
}

// Difference removes all elements of another set from this set.
func (s *Set) Difference(other *Set) *Set {
	for i := range other.set {
		s.Remove(i)
	}
	return s
}

// Intersection updates set with the intersection of itself and another.
func (s *Set) Intersection(other *Set) *Set {
	for i := range s.set {
		if !other.IsIn(i) {
			s.Remove(i)
		}
	}
	return s
}

// Union updates set with the union of itself and another.
func (s *Set) Union(other *Set) *Set {
	for i := range other.set {
		s.Add(i)
	}
	return s
}

// SymmetricDifference updates set with the symmetric difference of itself and another.
func (s *Set) SymmetricDifference(other *Set) *Set {
	intersection := s.Copy().Intersection(other)
	s.Union(other).Difference(intersection)
	return s
}

// IsDisjoint checks if two sets have a null intersection.
func (s *Set) IsDisjoint(other *Set) bool {
	intersection := s.Copy().Intersection(other)
	isDisjoint := intersection.Empty()
	return isDisjoint
}

// IsSubset checks if another set contains this set.
func (s *Set) IsSubset(other *Set) bool {
	difference := s.Copy().Difference(other)
	isSubset := difference.Empty()
	return isSubset
}

// IsSuperset checks if this set contains another.
func (s *Set) IsSuperset(other *Set) bool {
	difference := other.Copy().Difference(s)
	isSuperset := difference.Empty()
	return isSuperset
}
