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
	"fmt"
	"github.com/aiwuTech/container"
	"github.com/aiwuTech/container/lists"
	"strings"
)

type BinaryHeap struct {
	list       *lists.ArrayList
	comparator container.CompareFunction
}

// Returns true if heap does not contain any elements.
func (heap *BinaryHeap) Empty() bool {
	return heap.list.Empty()
}

// Returns number of elements within the heap.
func (heap *BinaryHeap) Len() int {
	return heap.list.Len()
}

// Removes all elements from the heap.
func (heap *BinaryHeap) Clear() {
	heap.list.Clear()
}

// check if the elements are in the heap
func (heap *BinaryHeap) Contains(elements ...interface{}) bool {
	return heap.list.Contains(elements...)
}

// Returns all elements in the heap.
func (heap *BinaryHeap) Elements() []interface{} {
	return heap.list.Elements()
}

func (heap *BinaryHeap) String() string {
	str := "BinaryHeap{ "
	values := []string{}
	for _, value := range heap.Elements() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	str += " }"
	return str
}

func NewBinaryHeap(comparator container.CompareFunction) *BinaryHeap {
	return &BinaryHeap{
		list:       lists.NewArrayList(),
		comparator: comparator,
	}
}

func (heap *BinaryHeap) Push(val interface{}) {
	heap.list.Add(val)
	heap.bubbleUp()
}

func (heap *BinaryHeap) Pop() (val interface{}, ok bool) {
	val, ok = heap.list.Get(0)
	if !ok {
		return
	}
	lastIndex := heap.list.Len() - 1
	heap.list.Swap(0, lastIndex)
	heap.list.Remove(lastIndex)

	heap.bubbleDown()
	return
}

// Returns top element on the heap without removing it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to peek.
func (heap *BinaryHeap) Peek() (val interface{}, ok bool) {
	return heap.list.Get(0)
}

// Performs the "bubble down" operation. This is to place the element that is at the
// root of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *BinaryHeap) bubbleDown() {
	index := 0
	size := heap.list.Len()
	for leftIndex := index<<1 + 1; leftIndex < size; leftIndex = index<<1 + 1 {
		rightIndex := index<<1 + 2
		smallerIndex := leftIndex
		leftValue, _ := heap.list.Get(leftIndex)
		rightValue, _ := heap.list.Get(rightIndex)
		if rightIndex < size && heap.comparator(leftValue, rightValue) > 0 {
			smallerIndex = rightIndex
		}
		indexValue, _ := heap.list.Get(index)
		smallerValue, _ := heap.list.Get(smallerIndex)
		if heap.comparator(indexValue, smallerValue) > 0 {
			heap.list.Swap(index, smallerIndex)
		} else {
			break
		}
		index = smallerIndex
	}
}

// Performs the "bubble up" operation. This is to place a newly inserted
// element (i.e. last element in the list) in its correct place so that
// the heap maintains the min/max-heap order property.
func (heap *BinaryHeap) bubbleUp() {
	index := heap.list.Len() - 1
	for parentIndex := (index - 1) >> 1; index > 0; parentIndex = (index - 1) >> 1 {
		indexValue, _ := heap.list.Get(index)
		parentValue, _ := heap.list.Get(parentIndex)
		if heap.comparator(parentValue, indexValue) <= 0 {
			break
		}
		heap.list.Swap(index, parentIndex)
		index = parentIndex
	}
}
