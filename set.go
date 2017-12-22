// Package set provides implementation of unordered collection of unique values.
package set

import (
	"fmt"
	"strings"
	"sort"
	"errors"
	"reflect"
)

type Set struct {
	set map[int]struct{}
}

// Build new set.
func New(values ...int) *Set {
	set := make(map[int]struct{})
	for _, v := range values {
		set[v] = struct{}{}
	}
	s := Set{set}
	return &s
}

// Get number of items in set.
func (s *Set) Len() int {
	return len(s.set)
}

// Check if set is empty.
func (s *Set) Empty() bool {
	return s.Len() == 0
}

// Get rid of all set from set.
func (s *Set) Clear() *Set {
	s.set = make(map[int]struct{})
	return s
}

// Make copy of set.
func (s *Set) Copy() *Set {
	cp := New()
	for i := range s.set {
		cp.Add(i)
	}
	return cp
}

// Check if sets contains the same elements.
func (s *Set) Equal(other *Set) bool {
	isEqual := reflect.DeepEqual(s.set, other.set)
	return isEqual
}

// Convert set to string.
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

// Add new value to a set.
func (s *Set) Add(value int) *Set {
	s.set[value] = struct{}{}
	return s
}

// Check if set contains value.
func (s *Set) IsIn(value int) bool {
	_, ok := s.set[value]
	return ok
}

// Remove a value from a set.
func (s *Set) Remove(value int) error {
	if !s.IsIn(value) {
		return errors.New("no such value in set")
	}
	delete(s.set, value)
	return nil
}

// Remove and return a random set element.
func (s *Set) Pop() (int, error) {
	if s.Len() == 0 {
		return 0, errors.New("the map is empty")
	}
	var value int
	for i := range s.set {
		value = i
		break
	}
	s.Remove(value)
	return value, nil
}

// Remove all elements of another set from this set.
func (s *Set) Difference(other *Set) *Set {
	for i := range other.set {
		s.Remove(i)
	}
	return s
}

// Update a set with the intersection of itself and another.
func (s *Set) Intersection(other *Set) *Set {
	for i := range s.set {
		if !other.IsIn(i) {
			s.Remove(i)
		}
	}
	return s
}

// Update a set with the union of itself and others.
func (s *Set) Union(other *Set) *Set {
	for i := range other.set {
		s.Add(i)
	}
	return s
}

// Update a set with the symmetric difference of itself and another.
func (s *Set) SymmetricDifference(other *Set) *Set {
	intersection := s.Copy().Intersection(other)
	s.Union(other).Difference(intersection)
	return s
}

// Return True if two sets have a null intersection.
func (s *Set) IsDisjoint(other *Set) bool {
	intersection := s.Copy().Intersection(other)
	isDisjoint := intersection.Empty()
	return isDisjoint
}

// Report whether another set contains this set.
func (s *Set) IsSubset(other *Set) bool {
	difference := s.Copy().Difference(other)
	isSubset := difference.Empty()
	return isSubset
}

// Report whether this set contains another set.
func (s *Set) IsSuperset(other *Set) bool {
	difference := other.Copy().Difference(s)
	isSuperset := difference.Empty()
	return isSuperset
}
