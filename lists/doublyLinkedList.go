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
package lists

import (
	"fmt"
	"github.com/aiwuTech/container"
	"sort"
	"strings"
)

type doublyLinkedElement struct {
	value interface{}
	prev  *doublyLinkedElement
	next  *doublyLinkedElement
}

type DoublyLinkedList struct {
	first       *doublyLinkedElement
	last        *doublyLinkedElement
	size        int
	compareFunc container.CompareFunction
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Appends a value (one or more) at the end of the list (same as Append())
func (list *DoublyLinkedList) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &doublyLinkedElement{
			value: value,
			prev:  list.last,
		}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Appends a value (one or more) at the end of the list (same as Add())
func (list *DoublyLinkedList) Append(values ...interface{}) {
	list.Add(values...)
}

// Prepends a values (or more)
func (list *DoublyLinkedList) Prepend(values ...interface{}) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &doublyLinkedElement{value: values[v], next: list.first}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}

// Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *DoublyLinkedList) Get(index int) (interface{}, bool) {

	if !list.inRange(index) {
		return nil, false
	}

	// determine traveral direction, last to first or first to last
	if list.size-index < index {
		element := list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
		return element.value, true
	} else {
		element := list.first
		for e := 0; e != index; e, element = e+1, element.next {
		}
		return element.value, true
	}
}

// Removes one or more elements from the list with the supplied indices.
func (list *DoublyLinkedList) Remove(index int) {

	if !list.inRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var element *doublyLinkedElement
	// determine traversal direction, last to first or first to last
	if list.size-index < index {
		element = list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
	} else {
		element = list.first
		for e := 0; e != index; e, element = e+1, element.next {
		}
	}

	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = element.prev
	}
	if element.prev != nil {
		element.prev.next = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}

	element = nil

	list.size--
}

// Check if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *DoublyLinkedList) Contains(values ...interface{}) bool {

	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		if !list.contain(value) {
			return false
		}
	}
	return true
}

// Returns all elements in the list.
func (list *DoublyLinkedList) Elements() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

// Returns true if list does not contain any elements.
func (list *DoublyLinkedList) Empty() bool {
	return list.size == 0
}

// Returns number of elements within the list.
func (list *DoublyLinkedList) Len() int {
	return list.size
}

// Removes all elements from the list.
func (list *DoublyLinkedList) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Sorts values
func (list *DoublyLinkedList) Sort(comparators ...container.CompareFunction) {
	if len(comparators) == 0 {
		sort.Sort(list)
	}

	comparator := comparators[0]
	list.compareFunc = comparator

	sort.Sort(list)
}

// Swaps values of two elements at the given indices.
func (list *DoublyLinkedList) Swap(i, j int) {
	if list.inRange(i) && list.inRange(j) && i != j {
		var element1, element2 *doublyLinkedElement
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

func (list *DoublyLinkedList) Less(i, j int) bool {
	if list.compareFunc == nil {
		return false
	}
	iVal, _ := list.Get(i)
	jVal, _ := list.Get(j)
	return list.compareFunc(iVal, jVal) < 0
}

func (list *DoublyLinkedList) String() string {
	str := "DoublyLinkedList{ "
	values := []string{}
	for element := list.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	str += " }"
	return str
}

func (list *DoublyLinkedList) contain(value interface{}) bool {
	for element := list.first; element != nil; element = element.next {
		if element.value == value {
			return true
		}
	}
	return false
}

// Check that the index is withing bounds of the list
func (list *DoublyLinkedList) inRange(index int) bool {
	return index >= 0 && index < list.size && list.size != 0
}
