// Package set provides implementation of unordered collection of unique values.
package set

type Set struct {
	set map[int]struct{}
}

// Build new set.
func NewSet(values ...int) *Set { return nil }

// Convert set to string.
func (s *Set) String() string { return "" }

// Get number of items in set.
func (s *Set) Len() int { return 0 }

// Get rid of all values from set.
func (s *Set) Clear() *Set { return nil }

// Make copy of set.
func (s *Set) Copy() *Set { return nil }

// Add new value to a set.
func (s *Set) Add(value int) *Set { return nil }

// Check if set contains value.
func (s *Set) IsIn(value int) bool { return false }

// Remove a value from a set.
func (s *Set) Remove(value int) error { return nil }

// Remove and return a random set element.
func (s *Set) Pop() (int, error) { return 0, nil }

// Remove all elements of another set from this set.
func (s *Set) Difference(other *Set) *Set { return nil }

// Update a set with the intersection of itself and another.
func (s *Set) Intersection(other *Set) *Set { return nil }

// Update a set with the symmetric difference of itself and another.
func (s *Set) SymmetricDifference(other *Set) *Set { return nil }

// Update a set with the union of itself and others.
func (s *Set) Union(other *Set) *Set { return nil }

// Return True if two sets have a null intersection.
func (s *Set) IsDisjoint(other *Set) bool { return false }

// Report whether another set contains this set.
func (s *Set) IsSubset(other *Set) bool { return false }

// Report whether this set contains another set.
func (s *Set) IsSuperset(other *Set) bool { return false }
