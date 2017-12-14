// Copyright 2017 Matthew Juran
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package unordered

import ()

// A Comparable can be checked for equality against others of the same underlying type and follows the pattern of Item. You define what a comparable item is.
type Comparable interface {
	Equal(Comparable) bool
}

// An EqualSet follows the same patterns as Set but holds items that are comparable to each other. The comparable constraint expands capabilities of the EqualSet with Remove, Reduce, Has, Equal, and Diff.
type EqualSet []Comparable

// Adds a new item to the set. Duplicates are allowed.
func (an EqualSet) Add(the Comparable) EqualSet {
	return an.set().Add(the.(Item)).equalset()
}

// Combines items in the receiver set with items of the argument sets into a new set. Duplicates are not removed.
func (an EqualSet) Combine(with ...EqualSet) EqualSet {
	return an.set().Combine(setslice(with)...).equalset()
}

// Removes one matching item. Use RemoveAll to remove all matches.
func (an EqualSet) Remove(the Comparable) EqualSet {
	if asserting {
		if an == nil {
			panic("unordered: nil set")
		}
		if the == nil {
			panic("unordered: nil arg")
		}
	}
	out := make(EqualSet, 0, len(an))
	found := false
	for _, item := range an {
		if (found == false) && item.Equal(the) {
			found = true
			continue
		}
		out = out.Add(item)
	}
	return out
}

// Removes all matching items from the set.
func (an EqualSet) RemoveAll(the Comparable) EqualSet {
	if asserting {
		if an == nil {
			panic("unordered: nil set")
		}
		if the == nil {
			panic("unordered: nil arg")
		}
	}
	out := make(EqualSet, 0, len(an))
	for _, item := range an {
		if item.Equal(the) {
			continue
		}
		out = out.Add(item)
	}
	return out
}

// Reduces the set by eliminating all duplicate items.
func (an EqualSet) Reduce() EqualSet {
	if asserting {
		if an == nil {
			panic("unordered: nil set")
		}
	}
	out := make(EqualSet, 0, len(an))
	for _, item := range an {
		if out.Has(item) {
			continue
		}
		out = out.Add(item)
	}
	return out
}

// If the set has the item then true is returned.
func (an EqualSet) Has(the Comparable) bool {
	if asserting {
		if an == nil {
			panic("unordered: nil set")
		}
		if the == nil {
			panic("unordered: nil arg")
		}
	}
	for _, item := range an {
		if the.Equal(item) {
			return true
		}
	}
	return false
}

// If both sets contain an equal count of each item then true is returned.
func (an EqualSet) Equal(to EqualSet) bool {
	if asserting {
		if an == nil {
			panic("unordered: nil set")
		}
		if to == nil {
			panic("unordered: nil arg")
		}
	}
	l := len(an)
	if l != len(to) {
		return false
	}
	if l == 0 {
		return true
	}
	firstCount := make(map[Comparable]uint)
	secondCount := make(map[Comparable]uint)
	for _, item := range an {
		_, has := firstCount[item]
		if has {
			firstCount[item]++
			continue
		}
		firstCount[item] = 1
	}
	for _, item := range to {
		_, has := secondCount[item]
		if has {
			secondCount[item]++
			continue
		}
		secondCount[item] = 1
	}
OUTER:
	for item, count := range firstCount {
		for secondItem, secondCount := range secondCount {
			if secondItem.Equal(item) {
				if count != secondCount {
					return false
				}
				continue OUTER
			}
		}
		return false
	}
	return true
}

// Provides a set of the items not in both sets. Duplicates are not removed.
func (an EqualSet) Diff(from EqualSet) EqualSet {
	if asserting {
		if an == nil {
			panic("unordered: nil set")
		}
		if from == nil {
			panic("unordered: nil arg")
		}
	}
	out := make(EqualSet, 0, len(an))
	for _, item := range an {
		if from.Has(item) == false {
			out = out.Add(item)
		}
	}
	for _, item := range from {
		if an.Has(item) == false {
			out = out.Add(item)
		}
	}
	return out
}

func (an EqualSet) set() Set {
	if asserting {
		if an == nil {
			panic("unordered: nil slice arg")
		}
	}
	out := make(Set, len(an))
	for i, item := range an {
		out[i] = item.(Item)
	}
	return out
}

func setslice(the []EqualSet) []Set {
	if asserting {
		if the == nil {
			panic("unordered: nil slice arg")
		}
	}
	out := make([]Set, len(the))
	for i, set := range the {
		out[i] = set.set()
	}
	return out
}
