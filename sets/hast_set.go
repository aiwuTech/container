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
	"sync"
)

type HashSet struct {
	m    map[interface{}]bool
	lock *sync.Mutex
}

var _ Set = &HashSet{}

func NewHashSet() *HashSet {
	return &HashSet{
		m:    make(map[interface{}]bool),
		lock: &sync.Mutex{},
	}
}

func (set *HashSet) Add(elements ...interface{}) {
	set.lock.Lock()
	for _, e := range elements {
		if !set.m[e] {
			set.m[e] = true
		}
	}
	set.lock.Unlock()
}

func (set *HashSet) Remove(elements ...interface{}) {
	set.lock.Lock()
	for _, e := range elements {
		delete(set.m, e)
	}
	set.lock.Unlock()
}

// whether all the elements are in the set
// return true if all in, or false
func (set *HashSet) Contains(elements ...interface{}) bool {
	set.lock.Lock()
	for _, e := range elements {
		if !set.m[e] {
			set.lock.Unlock()
			return false
		}
	}
	set.lock.Unlock()
	return true
}

func (set *HashSet) Clear() {
	set.lock.Lock()
	set.m = make(map[interface{}]bool)
	set.lock.Unlock()
}

func (set *HashSet) Len() int {
	set.lock.Lock()
	len := len(set.m)
	set.lock.Unlock()
	return len
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
	set.lock.Lock()
	for key := range set.m {
		snapshot = append(snapshot, key)
	}
	set.lock.Unlock()

	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("HashSet{ ")
	first := true
	for _, key := range set.Elements() {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString(" }")

	return buf.String()
}
