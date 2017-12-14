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

type SetAddCase struct {
	Set
	Item
	Out Set
}

var SetAddCases = []SetAddCase{
	{
		Set:  Set{0, 0, 1, 2, 3, 4, 5},
		Item: 0,
		Out:  Set{5, 4, 3, 2, 1, 0, 0, 0},
	},
	{
		Set:  Set{"hello", ",", "world"},
		Item: "!",
		Out:  Set{"world", "hello", "!", ","},
	},
}

func TestSetAdd(t *testing.T) {
	for i, c := range SetAddCases {
		out := c.Set.Add(c.Item)
		if len(out) != len(c.Out) {
			t.Fatalf("%v: failed %v vs %v", i, len(out), len(c.Out))
		}
		// TODO: this test requires type switching on the actual type in the test case to compare for equality
	}
}

type SetCombineCase struct {
	A   Set
	B   Set
	C   Set
	Out Set
}

var SetCombineCases = []SetCombineCase{
	{
		A:   Set{0, 1, 2},
		B:   Set{2, 1, 0},
		C:   Set{1, 2, 0},
		Out: Set{0, 0, 1, 1, 2, 2, 0, 1, 2},
	},
	{
		A:   Set{"hello"},
		B:   Set{"world"},
		C:   Set{",", "!"},
		Out: Set{"hello", ",", "world", "!"},
	},
}

func TestSetCombine(t *testing.T) {
	for i, c := range SetCombineCases {
		out := c.A.Combine(c.B, c.C)
		if (len(c.Out) != (len(c.A) + len(c.B) + len(c.C))) || (len(out) != len(c.Out)) {
			t.Fatalf("%v: failed %v vs %v", i, len(c.Out), len(c.A)+len(c.B)+len(c.C))
		}
		// TODO: this test requires type switching on the actual types in the test case to compare for equality
	}
}
