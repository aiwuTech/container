// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package trees

import (
	"github.com/aiwuTech/container"
	"math/rand"
	"testing"
)

func TestBinaryHeap(t *testing.T) {

	heap := NewBinaryHeap(container.IntCompareFunctionASC)

	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	// insertions
	heap.Push(3)
	// [3]
	heap.Push(2)
	// [2,3]
	heap.Push(1)
	// [1,3,2](2 swapped with 1, hence last)

	if actualValue := heap.Elements(); actualValue[0].(int) != 1 || actualValue[1].(int) != 3 || actualValue[2].(int) != 2 {
		t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
	}

	if actualValue := heap.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := heap.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := heap.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	heap.Pop()

	if actualValue, ok := heap.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := heap.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := heap.Pop(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := heap.Pop(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := heap.Elements(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}

	rand.Seed(3)
	for i := 0; i < 10000; i++ {
		r := int(rand.Int31n(30))
		heap.Push(r)
	}

	prev, _ := heap.Pop()
	for !heap.Empty() {
		curr, _ := heap.Pop()
		if prev.(int) > curr.(int) {
			t.Errorf("Heap property invalidated. prev: %v current: %v", prev, curr)
		}
		prev = curr
	}

}

func BenchmarkBinaryHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heap := NewBinaryHeap(container.IntCompareFunctionASC)
		for n := 0; n < 1000; n++ {
			heap.Push(i)
		}
		for !heap.Empty() {
			heap.Pop()
		}
	}

}
