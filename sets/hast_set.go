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
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{
		m: make(map[interface{}]bool),
	}
}

func (set *HashSet) Add(elements ...interface{}) {
	for _, e := range elements {
		if !set.m[e] {
			set.m[e] = true
		}
	}
}

func (set *HashSet) Remove(elements ...interface{}) {
	for _, e := range elements {
		delete(set.m, e)
	}
}

// whether all the elements are in the set
// return true if all in, or false
func (set *HashSet) Contains(elements ...interface{}) bool {
	for _, e := range elements {
		if !set.m[e] {
			return false
		}
	}
	return true
}

func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

func (set *HashSet) Len() int {
	return len(set.m)
}

func (set *HashSet) Empty() bool {
	return set.Len() == 0
}

func (set *HashSet) Same(other Set) bool {
	if other == nil || set.Len() != other.Len() {
		return false
	}

	return other.Contains(set.Elements()...)
}

func (set *HashSet) Elements() []interface{} {
	snapshot := make([]interface{}, 0)
	for key := range set.m {
		snapshot = append(snapshot, key)
	}

	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("HashSet{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")

	return buf.String()
}
