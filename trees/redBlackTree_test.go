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
    "testing"
    "github.com/aiwuTech/container"
)

func TestRedBlackTree(t *testing.T) {

    tree := NewRBTree(container.IntCompareFunctionASC)

    // insertions
    tree.Put(5, "e")
    tree.Put(6, "f")
    tree.Put(7, "g")
    tree.Put(3, "c")
    tree.Put(4, "d")
    tree.Put(1, "x")
    tree.Put(2, "b")
    tree.Put(1, "a") //overwrite

    // Test Size()
    if actualValue := tree.Len(); actualValue != 7 {
        t.Errorf("Got %v expected %v", actualValue, 7)
    }

    // test Keys()
    if actualValue, expactedValue := fmt.Sprintf("%d%d%d%d%d%d%d", tree.Keys()...), "1234567"; actualValue != expactedValue {
        t.Errorf("Got %v expected %v", actualValue, expactedValue)
    }

    // test Values()
    if actualValue, expactedValue := fmt.Sprintf("%s%s%s%s%s%s%s", tree.Elements()...), "abcdefg"; actualValue != expactedValue {
        t.Errorf("Got %v expected %v", actualValue, expactedValue)
    }

    // key,expectedValue,expectedFound
    tests1 := [][]interface{}{
        {1, "a", true},
        {2, "b", true},
        {3, "c", true},
        {4, "d", true},
        {5, "e", true},
        {6, "f", true},
        {7, "g", true},
        {8, nil, false},
    }

    for _, test := range tests1 {
        // retrievals
        actualValue, actualFound := tree.Get(test[0])
        if actualValue != test[1] || actualFound != test[2] {
            t.Errorf("Got %v expected %v", actualValue, test[1])
        }
    }

    // removals
    tree.Remove(5)
    tree.Remove(6)
    tree.Remove(7)
    tree.Remove(8)
    tree.Remove(5)

    // Test Keys()
    if actualValue, expactedValue := fmt.Sprintf("%d%d%d%d", tree.Keys()...), "1234"; actualValue != expactedValue {
        t.Errorf("Got %v expected %v", actualValue, expactedValue)
    }

    // test Values()
    if actualValue, expactedValue := fmt.Sprintf("%s%s%s%s", tree.Elements()...), "abcd"; actualValue != expactedValue {
        t.Errorf("Got %v expected %v", actualValue, expactedValue)
    }

    // Test Size()
    if actualValue := tree.Len(); actualValue != 4 {
        t.Errorf("Got %v expected %v", actualValue, 7)
    }

    tests2 := [][]interface{}{
        {1, "a", true},
        {2, "b", true},
        {3, "c", true},
        {4, "d", true},
        {5, nil, false},
        {6, nil, false},
        {7, nil, false},
        {8, nil, false},
    }

    for _, test := range tests2 {
        // retrievals
        actualValue, actualFound := tree.Get(test[0])
        if actualValue != test[1] || actualFound != test[2] {
            t.Errorf("Got %v expected %v", actualValue, test[1])
        }
    }

    // removals
    tree.Remove(1)
    tree.Remove(4)
    tree.Remove(2)
    tree.Remove(3)
    tree.Remove(2)
    tree.Remove(2)

    // Test Keys()
    if actualValue, expactedValue := fmt.Sprintf("%s", tree.Keys()), "[]"; actualValue != expactedValue {
        t.Errorf("Got %v expected %v", actualValue, expactedValue)
    }

    // test Values()
    if actualValue, expactedValue := fmt.Sprintf("%s", tree.Elements()), "[]"; actualValue != expactedValue {
        t.Errorf("Got %v expected %v", actualValue, expactedValue)
    }

    // Test Size()
    if actualValue := tree.Len(); actualValue != 0 {
        t.Errorf("Got %v expected %v", actualValue, 0)
    }

    // Test Empty()
    if actualValue := tree.Empty(); actualValue != true {
        t.Errorf("Got %v expected %v", actualValue, true)
    }

    tree.Put(1, "a")
    tree.Put(2, "b")
    tree.Clear()

    // Test Empty()
    if actualValue := tree.Empty(); actualValue != true {
        t.Errorf("Got %v expected %v", actualValue, true)
    }

}

func BenchmarkRedBlackTree(b *testing.B) {
    for i := 0; i < b.N; i++ {
        tree := NewRBTree(container.IntCompareFunctionASC)
        for n := 0; n < 1000; n++ {
            tree.Put(n, n)
        }
        for n := 0; n < 1000; n++ {
            tree.Remove(n)
        }
    }
}