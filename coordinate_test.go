// Copyright 2017 Matthew Juran
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package unordered

import (
	"testing"
)

type Coordinate struct {
	X int
	Y int
}

func (the Coordinate) Equal(to Comparable) bool {
	if the != to.(Coordinate) {
		return false
	}
	return true
}

// defining a type makes using the generic Equal Set type safe except for iteration where a type assertion is required
type CoordinateSet EqualSet

func coordinateSetSliceToEqualSetSlice(a []CoordinateSet) []EqualSet {
	out := make([]EqualSet, len(a))
	for i, set := range a {
		out[i] = EqualSet(set)
	}
	return out
}

func (the CoordinateSet) Add(a Coordinate) CoordinateSet {
	return CoordinateSet(EqualSet(the).Add(a))
}

func (the CoordinateSet) Combine(with ...CoordinateSet) CoordinateSet {
	return CoordinateSet(EqualSet(the).Combine(coordinateSetSliceToEqualSetSlice(with)...))
}

func (the CoordinateSet) Remove(a Coordinate) CoordinateSet {
	return CoordinateSet(EqualSet(the).Remove(a))
}

func (the CoordinateSet) RemoveAll(a Coordinate) CoordinateSet {
	return CoordinateSet(EqualSet(the).RemoveAll(a))
}

func (the CoordinateSet) Reduce() CoordinateSet {
	return CoordinateSet(EqualSet(the).Reduce())
}

func (the CoordinateSet) Has(a Coordinate) bool {
	return EqualSet(the).Has(a)
}

func (the CoordinateSet) Equal(to CoordinateSet) bool {
	return EqualSet(the).Equal(EqualSet(to))
}

func (the CoordinateSet) Diff(from CoordinateSet) CoordinateSet {
	return CoordinateSet(EqualSet(the).Diff(EqualSet(from)))
}

func TestCoordinateSetOverview(t *testing.T) {
	set := make(CoordinateSet, 0, 8)
	set = set.Add(Coordinate{0, 0})
	set = set.Add(Coordinate{1, 1})
	set = set.Add(Coordinate{2, 2})
	set = set.Add(Coordinate{3, 3})
	set = set.Add(Coordinate{0, 0})
	if set.Equal(CoordinateSet{Coordinate{3, 3}, Coordinate{1, 1}, Coordinate{0, 0}, Coordinate{0, 0}, Coordinate{2, 2}}) == false {
		t.Fatal("Coordinate Equal failed")
	}
	if set.Equal(CoordinateSet{Coordinate{1, 1}, Coordinate{3, 3}, Coordinate{2, 2}, Coordinate{0, 0}}) != false {
		t.Fatalf("Coordinate Equal false failed")
	}
	set = set.Combine(CoordinateSet{Coordinate{5, 5}}, CoordinateSet{Coordinate{6, 6}, Coordinate{7, 7}})
	if set.Equal(CoordinateSet{Coordinate{3, 3}, Coordinate{1, 1}, Coordinate{0, 0}, Coordinate{0, 0}, Coordinate{2, 2}}) != false {
		t.Fatal("Coordinate Equal false failed")
	}
	if set.Equal(CoordinateSet{Coordinate{3, 3}, Coordinate{1, 1}, Coordinate{2, 2}, Coordinate{7, 7}, Coordinate{5, 5}, Coordinate{6, 6}, Coordinate{0, 0}, Coordinate{0, 0}}) == false {
		t.Fatal("Coordinate Equal failed")
	}
	if set.Has(Coordinate{0, 0}) == false {
		t.Fatal("Coordinate Has failed")
	}
	if set.Has(Coordinate{1, 2}) == true {
		t.Fatal("Coordinate Has false failed")
	}
	count := make(map[Coordinate]uint)
	for _, coord := range set {
		_, has := count[coord.(Coordinate)]
		if has == false {
			count[coord.(Coordinate)] = 1
			continue
		}
		count[coord.(Coordinate)]++
	}
	if count[Coordinate{0, 0}] != 2 {
		t.Fatal("0,0 not 2")
	}
	if count[Coordinate{1, 1}] != 1 {
		t.Fatal("1,1 not 1")
	}
	if count[Coordinate{2, 2}] != 1 {
		t.Fatal("2,2 not 1")
	}
	if count[Coordinate{3, 3}] != 1 {
		t.Fatal("3,3 not 1")
	}
	if count[Coordinate{5, 5}] != 1 {
		t.Fatal("5,5 not 1")
	}
	if count[Coordinate{6, 6}] != 1 {
		t.Fatal("6,6 not 1")
	}
	if count[Coordinate{7, 7}] != 1 {
		t.Fatal("7,7 not 1")
	}
	if len(count) != 7 {
		t.Fatal("count not 7")
	}
	if set.Remove(Coordinate{0, 0}).Equal(CoordinateSet{Coordinate{1, 1}, Coordinate{2, 2}, Coordinate{3, 3}, Coordinate{5, 5}, Coordinate{7, 7}, Coordinate{6, 6}, Coordinate{0, 0}}) == false {
		t.Fatal("Remove set not equal")
	}
	if set.Reduce().Equal(CoordinateSet{Coordinate{1, 1}, Coordinate{0, 0}, Coordinate{2, 2}, Coordinate{3, 3}, Coordinate{5, 5}, Coordinate{7, 7}, Coordinate{6, 6}}) == false {
		t.Fatal("Reduce set not equal")
	}
	set = set.RemoveAll(Coordinate{0, 0})
	if set.Equal(CoordinateSet{Coordinate{1, 1}, Coordinate{2, 2}, Coordinate{3, 3}, Coordinate{5, 5}, Coordinate{7, 7}, Coordinate{6, 6}}) == false {
		t.Fatal("RemoveAll set not equal")
	}
	if set.Diff(CoordinateSet{Coordinate{1, 1}, Coordinate{2, 2}, Coordinate{3, 3}}).Equal(CoordinateSet{Coordinate{5, 5}, Coordinate{6, 6}, Coordinate{7, 7}}) == false {
		t.Fatal("Diff set not equal")
	}
}
