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

type Int int

func (i Int) Equal(to Comparable) bool {
	if i != to.(Int) {
		return false
	}
	return true
}

type String string

func (s String) Equal(to Comparable) bool {
	if s != to.(String) {
		return false
	}
	return true
}

type EqualSetAddCase struct {
	EqualSet
	Comparable
	Out EqualSet
}

var EqualSetAddCases = []EqualSetAddCase{
	{
		EqualSet:   EqualSet{Int(0), Int(1)},
		Comparable: Int(2),
		Out:        EqualSet{Int(1), Int(2), Int(0)},
	},
	{
		EqualSet:   EqualSet{String("hello"), String("world")},
		Comparable: String(","),
		Out:        EqualSet{String("hello"), String(","), String("world")},
	},
}

func TestEqualSetAdd(t *testing.T) {
	for i, c := range EqualSetAddCases {
		if c.EqualSet.Add(c.Comparable).Equal(c.Out) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

type EqualSetCombineCase struct {
	A   EqualSet
	B   EqualSet
	C   EqualSet
	Out EqualSet
}

var EqualSetCombineCases = []EqualSetCombineCase{
	{
		A:   EqualSet{Int(1)},
		B:   EqualSet{Int(2)},
		C:   EqualSet{Int(3)},
		Out: EqualSet{Int(2), Int(1), Int(3)},
	},
	{
		A:   EqualSet{String("hello")},
		B:   EqualSet{String("world")},
		C:   EqualSet{String(","), String("!")},
		Out: EqualSet{String("hello"), String(","), String("world"), String("!")},
	},
}

func TestEqualSetCombine(t *testing.T) {
	for i, c := range EqualSetCombineCases {
		if c.A.Combine(c.B, c.C).Equal(c.Out) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

type EqualSetRemoveCase struct {
	EqualSet
	Comparable
	Out EqualSet
}

var EqualSetRemoveCases = []EqualSetRemoveCase{
	{
		EqualSet:   EqualSet{Int(1), Int(2), Int(3)},
		Comparable: Int(2),
		Out:        EqualSet{Int(3), Int(1)},
	},
	{
		EqualSet:   EqualSet{String("hello"), String(","), String("world")},
		Comparable: String(","),
		Out:        EqualSet{String("hello"), String("world")},
	},
}

func TestEqualSetRemove(t *testing.T) {
	for i, c := range EqualSetRemoveCases {
		if c.EqualSet.Remove(c.Comparable).Equal(c.Out) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

var EqualSetRemoveAllCases = []EqualSetRemoveCase{
	{
		EqualSet:   EqualSet{Int(1), Int(2), Int(3), Int(2), Int(2)},
		Comparable: Int(2),
		Out:        EqualSet{Int(3), Int(1)},
	},
	{
		EqualSet:   EqualSet{String("hello"), String(","), String("world"), String(",")},
		Comparable: String(","),
		Out:        EqualSet{String("hello"), String("world")},
	},
}

func TestEqualSetRemoveAll(t *testing.T) {
	for i, c := range EqualSetRemoveAllCases {
		if c.EqualSet.RemoveAll(c.Comparable).Equal(c.Out) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

type EqualSetReduceCase struct {
	EqualSet
	Out EqualSet
}

var EqualSetReduceCases = []EqualSetReduceCase{
	{
		EqualSet: EqualSet{Int(1), Int(2), Int(3), Int(2), Int(2)},
		Out:      EqualSet{Int(3), Int(1), Int(2)},
	},
	{
		EqualSet: EqualSet{String("hello"), String(","), String("world"), String(",")},
		Out:      EqualSet{String("hello"), String(","), String("world")},
	},
}

func TestEqualSetReduce(t *testing.T) {
	for i, c := range EqualSetReduceCases {
		if c.EqualSet.Reduce().Equal(c.Out) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

type EqualSetHasCase struct {
	EqualSet
	Comparable
	Has bool
}

var EqualSetHasCases = []EqualSetHasCase{
	{
		EqualSet:   EqualSet{Int(1), Int(2), Int(3), Int(2), Int(2)},
		Comparable: Int(1),
		Has:        true,
	},
	{
		EqualSet:   EqualSet{Int(1), Int(2), Int(3), Int(2), Int(2)},
		Comparable: Int(4),
		Has:        false,
	},
	{
		EqualSet:   EqualSet{String("hello"), String(","), String("world"), String(",")},
		Comparable: String("hello"),
		Has:        true,
	},
	{
		EqualSet:   EqualSet{String("hello"), String(","), String("world"), String(",")},
		Comparable: String("hello, world"),
		Has:        false,
	},
}

func TestEqualSetHas(t *testing.T) {
	for i, c := range EqualSetHasCases {
		if c.EqualSet.Has(c.Comparable) != c.Has {
			t.Fatalf("%v failed", i)
		}
	}
}

type EqualSetEqualCase struct {
	A     EqualSet
	B     EqualSet
	Equal bool
}

var EqualSetEqualCases = []EqualSetEqualCase{
	{
		A:     EqualSet{Int(1), Int(2), Int(3), Int(2), Int(2)},
		B:     EqualSet{Int(2), Int(2), Int(3), Int(1), Int(2)},
		Equal: true,
	},
	{
		A:     EqualSet{Int(2), Int(3), Int(2), Int(2)},
		B:     EqualSet{Int(2), Int(2), Int(3), Int(1), Int(2)},
		Equal: false,
	},
	{
		A:     EqualSet{String("hello"), String(","), String("world"), String(",")},
		B:     EqualSet{String(","), String("hello"), String("world"), String(",")},
		Equal: true,
	},
	{
		A:     EqualSet{String("hello"), String(","), String("world"), String(",")},
		B:     EqualSet{String(","), String("hello"), String("world"), String("!")},
		Equal: false,
	},
}

func TestEqualSetEqual(t *testing.T) {
	for i, c := range EqualSetEqualCases {
		if c.A.Equal(c.B) != c.Equal {
			t.Fatalf("%v failed", i)
		}
	}
}

type EqualSetDiffCase struct {
	A    EqualSet
	B    EqualSet
	Diff EqualSet
}

var EqualSetDiffCases = []EqualSetDiffCase{
	{
		A:    EqualSet{Int(1), Int(2), Int(3), Int(2), Int(2)},
		B:    EqualSet{Int(2), Int(2), Int(3), Int(1), Int(2)},
		Diff: EqualSet{},
	},
	{
		A:    EqualSet{Int(2), Int(3), Int(2), Int(2)},
		B:    EqualSet{Int(2), Int(2), Int(3), Int(1), Int(2)},
		Diff: EqualSet{Int(1)},
	},
	{
		A:    EqualSet{String("hello"), String(","), String("world"), String(",")},
		B:    EqualSet{String(","), String("hello"), String("world"), String(",")},
		Diff: EqualSet{},
	},
	{
		A:    EqualSet{String("hello"), String(","), String("world"), String(",")},
		B:    EqualSet{String(","), String("hello"), String("world"), String("!")},
		Diff: EqualSet{String("!")},
	},
}

func TestEqualSetDiff(t *testing.T) {
	for i, c := range EqualSetDiffCases {
		if c.A.Diff(c.B).Equal(c.Diff) == false {
			t.Fatalf("%v failed", i)
		}
	}
}
