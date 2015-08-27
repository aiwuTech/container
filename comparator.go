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

// compareFunc的结果值：
//   小于0: 第一个参数小于第二个参数
//   等于0: 第一个参数等于第二个参数
//   大于1: 第一个参数大于第二个参数
type CompareFunction func(interface{}, interface{}) int8

func Float64CompareFunctionASC(e1, e2 interface{}) int8 {
	k1 := e1.(float64)
	k2 := e2.(float64)

	if k1 < k2 {
		return -1
	} else if k1 > k2 {
		return 1
	} else {
		return 0
	}
}
func Float64CompareFunctionDESC(e1, e2 interface{}) int8 {
	return -Float64CompareFunctionASC(e1, e2)
}

func Uint64CompareFunctionASC(e1, e2 interface{}) int8 {
	k1 := e1.(uint64)
	k2 := e2.(uint64)

	if k1 < k2 {
		return -1
	} else if k1 > k2 {
		return 1
	} else {
		return 0
	}
}

func Uint64CompareFunctionDESC(e1, e2 interface{}) int8 {
	return -Uint64CompareFunctionASC(e1, e2)
}

func Int64CompareFunctionASC(e1, e2 interface{}) int8 {
	k1 := e1.(int64)
	k2 := e2.(int64)

	if k1 < k2 {
		return -1
	} else if k1 > k2 {
		return 1
	} else {
		return 0
	}
}

func Int64CompareFunctionDESC(e1, e2 interface{}) int8 {
	return -Int64CompareFunctionASC(e1, e2)
}
