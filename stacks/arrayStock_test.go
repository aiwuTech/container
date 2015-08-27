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
	"testing"
)

func TestArrayStack(t *testing.T) {

	stack := NewArrayStack()

	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	// insertions
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if actualValue := stack.Elements(); actualValue[0].(int) != 3 || actualValue[1].(int) != 2 || actualValue[2].(int) != 1 {
		t.Errorf("Got %v expected %v", actualValue, "[3,2,1]")
	}

	if actualValue := stack.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := stack.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := stack.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	stack.Pop()

	if actualValue, ok := stack.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := stack.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := stack.Pop(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := stack.Pop(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := stack.Elements(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}

}

func BenchmarkArrayStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := NewArrayStack()
		for n := 0; n < 1000; n++ {
			stack.Push(i)
		}
		for !stack.Empty() {
			stack.Pop()
		}
	}

}
