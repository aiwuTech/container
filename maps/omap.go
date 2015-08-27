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
package maps

import (
	"bytes"
	"fmt"
	"reflect"
)

// 有序map接口
type OrderedMapper interface {
	MapInterface
	// 获取key对应的第一个值
	GetFirst(key interface {}) interface {}
	// 获取key对应的所有值
	GetAll(key interface {}) []interface {}
	// 获取第一个键值，没有返回nil
	FirstKey() interface{}
	// 获取最后一个键值，没有返回nil
	LastKey() interface{}
	// 获取 < toKey的键值的OrderedMapper
	Head(toKey interface{}) OrderedMapper
	// 获取 [fromKey, toKey)区间的OrderMapper
	Sub(fromKey, toKey interface{}) OrderedMapper
	// 获取 > fromKey的键值的OrderMapper
	Tail(fromKey interface{}) OrderedMapper
}

type omap struct {
	keys     Keys
	elemType reflect.Type
	m        map[interface{}][]interface{}
}

func (m *omap) Get(key interface{}) interface{} {
	return m.m[key]
}

func (m *omap) GetFirst(key interface {}) interface {} {
	return m.m[key][0]
}

func (m *omap) GetAll(key interface {}) []interface {} {
	return m.m[key]
}

func (m *omap) isAcceptableElem(elem interface{}) bool {
	if elem == nil {
		return false
	}

	if reflect.TypeOf(elem) != m.ElemType() {
		return false
	}

	return true
}

func (m *omap) Put(key interface{}, elem interface{}) (interface{}, bool) {
	if !m.isAcceptableElem(elem) {
		return nil, false
	}

	oldElems, ok := m.m[key]
	if ok {
		m.m[key] = append(m.m[key], elem)
	} else {
		m.keys.Add(key)
		m.m[key] = make([]interface {}, 0)
		m.m[key] = append(m.m[key], elem)
	}

	return oldElems, true
}

func (m *omap) Remove(key interface{}) interface{} {
	oldElem, ok := m.m[key]
	delete(m.m, key)
	if ok {
		m.keys.Remove(key)
	}

	return oldElem
}

func (m *omap) Clear() {
	m.m = make(map[interface{}][]interface{})
	m.keys.Clear()
}

func (m *omap) Len() int {
	length := 0
	for _, val := range m.m {
		length += len(val)
	}

	return length
}

func (m *omap) Contains(keys ...interface{}) bool {
	for key := range keys {
		if _, ok := m.m[key]; !ok {
			return false
		}
	}
	return true
}

func (m *omap) Empty() bool {
	return m.Len() == 0
}

func (m *omap) FirstKey() interface{} {
	if m.Len() == 0 {
		return nil
	}

	return m.keys.Get(0)
}

func (m *omap) LastKey() interface{} {
	len := m.Len()
	if len == 0 {
		return nil
	}

	return m.keys.Get(len - 1)
}

func (m *omap) Sub(fromKey, toKey interface{}) OrderedMapper {
	newOmap := NewOrderMap(NewKeys(m.keys.CompareFunc(), m.keys.ElemType()), m.ElemType())
	omapLen := m.Len()
	if omapLen == 0 {
		return newOmap
	}

	beginIndex, contains := m.keys.Search(fromKey)
	if !contains {
		beginIndex = 0
	}
	endIndex, contains := m.keys.Search(toKey)
	if !contains {
		endIndex = omapLen
	}

	var key interface{}
	var elems []interface {}
	for i := beginIndex; i < endIndex; i++ {
		key = m.keys.Get(i)
		elems = m.GetAll(key)

		for _, elem := range elems {
			newOmap.Put(key, elem)
		}
	}

	return newOmap
}

func (m *omap) Head(toKey interface{}) OrderedMapper {
	return m.Sub(nil, toKey)
}

func (m *omap) Tail(fromKey interface{}) OrderedMapper {
	return m.Sub(fromKey, nil)
}

func (m *omap) Keys() []interface{} {
	return m.keys.GetAll()
}

func (m *omap) Elements() []interface{} {
	elems := make([]interface{}, 0)
	for _, key := range m.Keys() {
		elems = append(elems, m.GetAll(key)...)
	}

	return elems
}

func (m *omap) ToMap() map[interface{}]interface{} {
	replica := make(map[interface{}]interface{})
	for k, v := range m.m {
		replica[k] = v
	}

	return replica
}

func (m *omap) KeyType() reflect.Type {
	return m.keys.ElemType()
}

func (m *omap) ElemType() reflect.Type {
	return m.elemType
}

func (m *omap) String() string {
	var buf bytes.Buffer
	buf.WriteString("OrderedMap<")
	buf.WriteString(m.KeyType().Kind().String())
	buf.WriteString(",")
	buf.WriteString(m.ElemType().Kind().String())
	buf.WriteString(">{")
	first := true
	omapLen := m.Len()
	for i := 0; i < omapLen; i++ {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}

		key := m.keys.Get(i)
		buf.WriteString(fmt.Sprintf("%v", key))
		buf.WriteString(":")
		buf.WriteString(fmt.Sprintf("%v", m.Get(key)))
	}
	buf.WriteString("}")

	return buf.String()
}

func NewOrderMap(keys Keys, elemType reflect.Type) OrderedMapper {
	return &omap{
		keys:     keys,
		elemType: elemType,
		m:        make(map[interface{}][]interface{}),
	}
}