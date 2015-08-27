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
package sets

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	set := NewHashSet()

	// insert
	set.Add(1)
	set.Add()
	set.Add(2, 4)
	set.Add(5, 3)
	set.Add([]interface{}{5,7,8}...)

	if set.Empty() {
		t.Errorf("Empty error, expected %v", false)
	}

	if set.Len() != 7 {
		t.Errorf("Len error, expected %v", 7)
	}

	if !set.Contains(4, 8, 7) {
		t.Errorf("Contains error, expected true")
	}

	if set.Contains(9) {
		t.Errorf("Contains error, expected false")
	}
}

func BenchmarkHashSet(b *testing.B) {
	for i := 0; i < b.N; i++{
		set := NewHashSet()
		for n := 0; n < 10000; n++{
			set.Add(n)
		}
		for n := 0; n < 10000; n++{
			set.Remove(n)
		}
	}
}