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

type LinkedListStack struct {
	list *lists.SinglyLinkedList
}

// Instantiates a new empty stack
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		list: lists.NewSinglyLinkedList(),
	}
}

// Pushes a value onto the top of the stack
func (stack *LinkedListStack) Push(value interface{}) {
	stack.list.Prepend(value)
}

// Pops (removes) top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *LinkedListStack) Pop() (value interface{}, ok bool) {
	value, ok = stack.list.Get(0)
	stack.list.Remove(0)
	return
}

// Returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *LinkedListStack) Peek() (value interface{}, ok bool) {
	return stack.list.Get(0)
}

// Returns true if stack does not contain any elements.
func (stack *LinkedListStack) Empty() bool {
	return stack.list.Empty()
}

// Returns number of elements within the stack.
func (stack *LinkedListStack) Len() int {
	return stack.list.Len()
}

// Removes all elements from the stack.
func (stack *LinkedListStack) Clear() {
	stack.list.Clear()
}

// Returns all elements in the stack (LIFO order).
func (stack *LinkedListStack) Elements() []interface{} {
	return stack.list.Elements()
}

func (stack *LinkedListStack) String() string {
	str := "LinkedListStack{ "
	values := []string{}
	for _, value := range stack.list.Elements() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
    str += " }"
	return str
}

// not important, just for interface{}
func (stack *LinkedListStack) Contains(elements ...interface{}) bool {
    return stack.list.Contains(elements...)
}
