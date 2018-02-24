// Package set provides implementation of unordered collection of unique values.
package set

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// ErrValueNotFound is returned when set does not contain given value.
var ErrValueNotFound = errors.New("no such value in set")

// ErrEmptySet is returned during trying pop from empty set.
var ErrEmptySet = errors.New("the map is empty")

// Set is wrapper for map with int as key.
type Set map[int]struct{}

// New builds a new set.
func New(values ...int) *Set {
	s := make(Set)
	s.Add(values...)
	return &s
}

// Len gets number of items in set.
func (s *Set) Len() int {
	i := len(*s)
	return i
}

// Empty checks if set contains no elements.
func (s *Set) Empty() bool {
	b := s.Len() == 0
	return b
}

// Clear gets rid of all set from set.
func (s *Set) Clear() *Set {
	*s = *New()
	return s
}

// Copy creates a new set with the same elements.
func (s *Set) Copy() *Set {
	slice := s.ToSlice()
	cp := New(slice...)
	return cp
}

// Equal checks if two sets contain the same elements.
func (s *Set) Equal(other *Set) bool {
	isEqual := reflect.DeepEqual(s, other)
	return isEqual
}

// ToSlice converts set to a slice of ints.
func (s *Set) ToSlice() []int {
	tabInt := make([]int, 0, s.Len())
	for i := range *s {
		tabInt = append(tabInt, i)
	}
	return tabInt
}

// String converts set to a string.
func (s *Set) String() string {
	tabInt := s.ToSlice()
	sort.Ints(tabInt)
	tabStr := make([]string, 0, s.Len())
	for _, i := range tabInt {
		tabStr = append(tabStr, fmt.Sprint(i))
	}
	inner := strings.Join(tabStr, ", ")
	str := fmt.Sprintf("{%s}", inner)
	return str
}

// Add inserts new values to a set.
func (s *Set) Add(values ...int) *Set {
	for _, v := range values {
		(*s)[v] = struct{}{}
	}
	return s
}

// IsIn checks if set contains value.
func (s *Set) IsIn(value int) bool {
	_, ok := (*s)[value]
	return ok
}

// Remove deletes a value from a set.
func (s *Set) Remove(value int) error {
	if !s.IsIn(value) {
		return ErrValueNotFound
	}
	delete(*s, value)
	return nil
}

// Pop removes and returns some element from a set.
func (s *Set) Pop() (int, error) {
	if s.Len() == 0 {
		return 0, ErrEmptySet
	}
	var value int
	for i := range *s {
		value = i
		break
	}
	s.Remove(value)
	return value, nil
}

// Difference removes all elements of another set from this set.
func (s *Set) Difference(other *Set) *Set {
	for i := range *other {
		s.Remove(i)
	}
	return s
}

// Intersection updates set with the intersection of itself and another.
func (s *Set) Intersection(other *Set) *Set {
	for i := range *s {
		if !other.IsIn(i) {
			s.Remove(i)
		}
	}
	return s
}

// Union updates set with the union of itself and another.
func (s *Set) Union(other *Set) *Set {
	slice := other.ToSlice()
	s.Add(slice...)
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
