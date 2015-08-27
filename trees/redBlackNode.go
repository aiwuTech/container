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
import "fmt"


type color bool

const (
    black, red color = true, false
)

type redBlackNode struct {
    key    interface{}
    value  interface{}
    color  color
    left   *redBlackNode
    right  *redBlackNode
    parent *redBlackNode
}

func (node *redBlackNode) maximumNode() *redBlackNode {
    for node.right != nil {
        node = node.right
    }
    return node
}


func (node *redBlackNode) grandparent() *redBlackNode {
    if node != nil && node.parent != nil {
        return node.parent.parent
    }
    return nil
}

func (node *redBlackNode) uncle() *redBlackNode {
    if node == nil || node.parent == nil || node.parent.parent == nil {
        return nil
    }
    return node.parent.sibling()
}

func (node *redBlackNode) sibling() *redBlackNode {
    if node == nil || node.parent == nil {
        return nil
    }
    if node == node.parent.left {
        return node.parent.right
    } else {
        return node.parent.left
    }
}

func (node *redBlackNode) String() string {
    return fmt.Sprintf("%v", node.key)
}

func output(node *redBlackNode, prefix string, isTail bool, str *string) {
    if node.right != nil {
        newPrefix := prefix
        if isTail {
            newPrefix += "│   "
        } else {
            newPrefix += "    "
        }
        output(node.right, newPrefix, false, str)
    }
    *str += prefix
    if isTail {
        *str += "└── "
    } else {
        *str += "┌── "
    }
    *str += node.String() + "\n"
    if node.left != nil {
        newPrefix := prefix
        if isTail {
            newPrefix += "    "
        } else {
            newPrefix += "│   "
        }
        output(node.left, newPrefix, true, str)
    }
}