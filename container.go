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
package container

import "reflect"

type ContainerInterface interface {
	Empty() bool
	Len() int
	Contains(elements ...interface{}) bool
	Clear()
	Elements() []interface{}
	String() string
}

func Contains(element interface{}, target interface{}) bool {
	valElement := reflect.ValueOf(element)
	typSlice := reflect.TypeOf(target)
	valSlice := reflect.ValueOf(target)

	if !valElement.IsValid() || !valSlice.IsValid() {
		return false
	}

	switch typSlice.Kind() {
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < valSlice.Len(); idx++ {
			val := valSlice.Index(idx)
			if !val.IsValid() {
				continue
			}

			if val.Interface() == valElement.Interface() {
				return true
			}
		}
	case reflect.Map:
		if valSlice.MapIndex(valElement).IsValid() {
			return true
		}
	}

	return false
}

func Delete(element interface{}, slice interface{}) bool {
	valElement := reflect.ValueOf(element)
	typSlice := reflect.TypeOf(slice)
	valSlice := reflect.ValueOf(slice)

	if !valElement.IsValid() || !valSlice.IsValid() {
		return false
	}

	switch typSlice.Kind() {
	case reflect.Slice:
		sliceLen := valSlice.Len()
		for idx := 0; idx < sliceLen; idx++ {
			val := valSlice.Index(idx)
			if !val.IsValid() {
				continue
			}

			if val.Interface() == valElement.Interface() {
				if idx == sliceLen-1 {
					valSlice = valSlice.Slice(0, idx)
				} else {
					valSlice = reflect.AppendSlice(valSlice.Slice(0, idx), valSlice.Slice(idx+1, sliceLen-1))
				}
			}
		}
	case reflect.Map:
	}

	return false
}

func Index(element interface{}, target interface{}) int {
	valElement := reflect.ValueOf(element)
	typSlice := reflect.TypeOf(target)
	valSlice := reflect.ValueOf(target)

	if !valElement.IsValid() || !valSlice.IsValid() {
		return -1
	}

	switch typSlice.Kind() {
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < valSlice.Len(); idx++ {
			val := valSlice.Index(idx)
			if !val.IsValid() {
				continue
			}

			if val.Interface() == valElement.Interface() {
				return idx
			}
		}
	}

	return -1
}
