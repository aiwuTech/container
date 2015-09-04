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

const (
	_EXPAND_FACTOR = float64(2.0)
	_SHRINK_FACTOR = float64(0.25)
)

type ArrayList struct {
	elements    []interface{}
	size        int
	compareFunc container.CompareFunction
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

// append elements at the end of the list
func (list *ArrayList) Add(elements ...interface{}) {
	list.expand(len(elements))
	for _, e := range elements {
		list.elements[list.size] = e
		list.size += 1
	}
}

// return the element at idx, if get the element, return element and true, otherwise nil, false
func (list *ArrayList) Get(idx int) (interface{}, bool) {
	if !list.inRange(idx) {
		return nil, false
	}

	return list.elements[idx], true
}

// remove the element form the arraylist
func (list *ArrayList) Remove(idx int) {
	if !list.inRange(idx) {
		return
	}

	list.elements[idx] = nil
	copy(list.elements[idx:], list.elements[idx+1:list.size])
	list.size -= 1

	list.shrink()
	return
}

// check if the elements are in the array list
func (list *ArrayList) Contains(elements ...interface{}) bool {
	for _, e := range elements {
		if !list.contain(e) {
			return false
		}
	}
	return true
}

// return all the elements in the arraylist
func (list *ArrayList) Elements() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// return true if the list's size is zero
func (list *ArrayList) Empty() bool {
	return list.Len() == 0
}

// return list's size
func (list *ArrayList) Len() int {
	return list.size
}

// remove all the elements
func (list *ArrayList) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

func (list *ArrayList) Clone() *ArrayList {
	l := NewArrayList()
	l.Add(list.Elements()...)
	return l
}

// sort the elements by comparator
func (list *ArrayList) Sort(comparators ...container.CompareFunction) {
	if len(comparators) == 0 {
		sort.Sort(list)
	}

	comparator := comparators[0]
	list.compareFunc = comparator

	sort.Sort(list)
}

// out format
func (list *ArrayList) String() string {
	str := "ArrayList{ "
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	str += " }"
	return str
}

// return whether e in the array list
func (list *ArrayList) contain(e interface{}) bool {
	for _, le := range list.elements {
		if e == le {
			return true
		}
	}
	return false
}

// check whether the idx is within bounds of the arraylist
func (list *ArrayList) inRange(idx int) bool {
	return idx >= 0 && idx < list.size && list.size != 0
}

// ReExpand the array list if necessary
func (list *ArrayList) expand(n int) {
	curCap := cap(list.elements)
	if list.size+n >= curCap {
		newCap := int(_EXPAND_FACTOR * float64(curCap+n))
		list.resize(newCap)
	}
}

// Shrink the array list if necessary
func (list *ArrayList) shrink() {
	curCap := cap(list.elements)
	if list.size <= int(float64(curCap)*_SHRINK_FACTOR) {
		list.resize(list.size)
	}
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *ArrayList) Less(i, j int) bool {
	if list.compareFunc == nil {
		return false
	}
	return list.compareFunc(list.elements[i], list.elements[j]) < 0
}

func (list *ArrayList) Swap(i, j int) {
	if list.inRange(i) && list.inRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}
