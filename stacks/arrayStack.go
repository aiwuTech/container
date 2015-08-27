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
package stacks

import (
	"fmt"
	"github.com/aiwuTech/container/lists"
	"strings"
)

type ArrayStack struct {
	list *lists.ArrayList
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		list: lists.NewArrayList(),
	}
}

// Pushes a value onto the top of the stack
func (stack *ArrayStack) Push(value interface{}) {
	stack.list.Add(value)
}

// Pops (removes) top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *ArrayStack) Pop() (value interface{}, ok bool) {
	value, ok = stack.list.Get(stack.list.Len() - 1)
	stack.list.Remove(stack.list.Len() - 1)
	return
}

// Returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *ArrayStack) Peek() (value interface{}, ok bool) {
	return stack.list.Get(stack.list.Len() - 1)
}

// Returns true if stack does not contain any elements.
func (stack *ArrayStack) Empty() bool {
	return stack.list.Empty()
}

// Returns number of elements within the stack.
func (stack *ArrayStack) Len() int {
	return stack.list.Len()
}

// Removes all elements from the stack.
func (stack *ArrayStack) Clear() {
	stack.list.Clear()
}

// Returns all elements in the stack (LIFO order).
func (stack *ArrayStack) Elements() []interface{} {
	size := stack.list.Len()
	elements := make([]interface{}, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1) // in reverse (LIFO)
	}
	return elements
}

func (stack *ArrayStack) String() string {
	str := "ArrayStack{ "
	values := []string{}
	for _, value := range stack.list.Elements() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	str += " }"
	return str
}

// not important, just for interface{}
func (stack *ArrayStack) Contains(elements ...interface{}) bool {
	return stack.list.Contains(elements...)
}
