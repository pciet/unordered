// Copyright 2017 Matthew Juran
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Provides generic unordered set types. For added type safety you may define a wrapper type:
//     // example for unordered.EqualSet
//     type Coordinate struct {
//         X int
//         Y int
//     }
//
//     // satisfies unordered.Comparable
//     func (a Coordinate) Equal(to unordered.Comparable) bool {
//         if a != to.(Coordinate) {
//             return false
//         }
//         return true
//     }
//
//     type CoordinateSet unordered.EqualSet
//
//     func (a CoordinateSet) Add(the Coordinate) CoordinateSet {
//         return CoordinateSet(unordered.EqualSet(a).Add(the))
//     }
//     ...
// Iteration requires a type assertion even with a wrapper:
//     for _, coord := range set {
//         if coord.(Coordinate).X == 1 {
//             return false
//         }
//     }
// This package is "Writing Solid Code"-style assertion enabled. To disable assertions vendor the package and set asserting = false in assert.go. You may want to do this if profiling shows a significant performance impact, but any panic from an assertion indicates an invalid program state.
package unordered

import (
	"fmt"
	"reflect"
)

// An Item can be stored in an unordered Set by being contained in an empty interface reference.
type Item interface{}

// The Set is created like a slice with:
//     make(unordered.Set, length, capacity)
// or:
//     make(unordered.Set, length&capacity)
// and iterated with:
//     for _, item := range set {
// Iteration order is not guaranteed because the set is unordered. A type assertion is required when operating on the iteration values.
//
// Use the built-in len function to get a count of items in the set:
//
//     // returns 3
//     len(unordered.Set{Int(1), Int(2), Int(1)})
//
//All items in the set are expected to be the same underlying type; a panic will occur if a different type of item is added.
type Set []Item

// Adds an item to the set. Duplicates are allowed.
func (a Set) Add(an Item) Set {
	if asserting {
		if a == nil {
			panic("unordered: Add called on nil set")
		}
		if an == nil {
			panic("unordered: nil Item provided to set.Add")
		}
		t := a.typeof()
		if (t != nil) && (t != reflect.TypeOf(an)) {
			panic(fmt.Sprintf("unordered: set type %v doesn't match new item (%v) type %v", t, an, reflect.TypeOf(an)))
		}
	}
	return append(a, an)
}

// Combines items in the receiver set with items of the argument sets into a new set. Duplicates are not removed.
func (a Set) Combine(with ...Set) Set {
	if asserting {
		if a == nil {
			panic("unordered: Combine called on nil set")
		}
		if len(with) == 0 {
			panic("unordered: Combine called for zero sets")
		}
		t := a.typeof()
		for _, s := range with {
			nt := s.typeof()
			if (t == nil) && (nt != nil) {
				t = nt
				continue
			}
			if t != nt {
				panic(fmt.Sprintf("set type %v doesn't match next set type %v", t, nt))
			}
		}
	}
	l := len(a)
	for _, s := range with {
		l += len(s)
	}
	out := make(Set, l)
	i := 0
	for _, item := range a {
		out[i] = item
		i++
	}
	for _, s := range with {
		for _, item := range s {
			out[i] = item
			i++
		}
	}
	return out
}

func (the Set) equalset() EqualSet {
	if asserting {
		if the == nil {
			panic("unordered: nil set")
		}
	}
	out := make(EqualSet, len(the))
	for i, item := range the {
		out[i] = item.(Comparable)
	}
	return out
}
