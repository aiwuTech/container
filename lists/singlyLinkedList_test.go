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
	"github.com/aiwuTech/container"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {

	list := NewSinglyLinkedList()

	list.Sort(container.StringCompareFunction)

	list.Add("e", "f", "g", "a", "b", "c", "d")

	list.Sort(container.StringCompareFunction)
	for i := 1; i < list.Len(); i++ {
		a, _ := list.Get(i - 1)
		b, _ := list.Get(i)
		if a.(string) > b.(string) {
			t.Errorf("Not sorted! %s > %s", a, b)
		}
	}

	list.Clear()

	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := list.Len(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	list.Add("a")
	list.Add("b", "c")

	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := list.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}

	list.Swap(0, 2)
	list.Swap(0, 2)
	list.Swap(0, 1)

	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}

	list.Remove(2)

	if actualValue, ok := list.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	list.Remove(1)
	list.Remove(0)

	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := list.Len(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	list.Add("a", "b", "c")

	if actualValue := list.Contains("a", "b", "c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := list.Contains("a", "b", "c", "d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	list.Clear()

	if actualValue := list.Contains("a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue, ok := list.Get(0); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

}

func BenchmarkSinglyLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := NewSinglyLinkedList()
		for n := 0; n < 1000; n++ {
			list.Add(i)
		}
		for !list.Empty() {
			list.Remove(0)
		}
	}
}
