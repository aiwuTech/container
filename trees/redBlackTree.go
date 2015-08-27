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

// just copied from github.com/emirpasic/gods, and a little recode
package trees

import (
	"github.com/aiwuTech/container"
	"github.com/aiwuTech/container/stacks"
)

type RBTree struct {
	root       *redBlackNode
	size       int
	comparator container.CompareFunction
}

func NewRBTree(comparator container.CompareFunction) *RBTree {
	return &RBTree{
		comparator: comparator,
	}
}

// Inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RBTree) Put(key interface{}, value interface{}) {
	insertedNode := &redBlackNode{key: key, value: value, color: red}
	if tree.root == nil {
		tree.root = insertedNode
	} else {
		node := tree.root
		loop := true
		for loop {
			compare := tree.comparator(key, node.key)
			switch {
			case compare == 0:
				node.value = value
				return
			case compare < 0:
				if node.left == nil {
					node.left = insertedNode
					loop = false
				} else {
					node = node.left
				}
			case compare > 0:
				if node.right == nil {
					node.right = insertedNode
					loop = false
				} else {
					node = node.right
				}
			}
		}
		insertedNode.parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size += 1
}

// Searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RBTree) Get(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.value, true
	}
	return nil, false
}

// Remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RBTree) Remove(key interface{}) {
	var child *redBlackNode
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.left != nil && node.right != nil {
		pred := node.left.maximumNode()
		node.key = pred.key
		node.value = pred.value
		node = pred
	}
	if node.left == nil || node.right == nil {
		if node.right == nil {
			child = node.left
		} else {
			child = node.right
		}
		if node.color == black {
			node.color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.parent == nil && child != nil {
			child.color = black
		}
	}
	tree.size -= 1
}

// Returns true if tree does not contain any nodes
func (tree *RBTree) Empty() bool {
	return tree.Len() == 0
}

// Returns number of nodes in the tree.
func (tree *RBTree) Len() int {
	return tree.size
}

// Returns all keys in-order
func (tree *RBTree) Keys() []interface{} {
	keys := make([]interface{}, tree.size)
	for i, node := range tree.inOrder() {
		keys[i] = node.key
	}
	return keys
}

// Returns all values in-order based on the key.
func (tree *RBTree) Elements() []interface{} {
	values := make([]interface{}, tree.size)
	for i, node := range tree.inOrder() {
		values[i] = node.value
	}
	return values
}

// Removes all nodes from the tree.
func (tree *RBTree) Clear() {
	tree.root = nil
	tree.size = 0
}

func (tree *RBTree) String() string {
	str := "RedBlackTree\n"
	if !tree.Empty() {
		output(tree.root, "", true, &str)
	}
	return str
}

func (tree *RBTree) Contains(elements ...interface{}) bool {
	nodes := tree.inOrder()
	for _, e := range elements {
		if !tree.contain(e, nodes) {
			return false
		}
	}
	return true
}

func (tree *RBTree) contain(element interface{}, nodes []*redBlackNode) bool {
	for _, node := range nodes {
		if element == node.value {
			return true
		}
	}
	return false
}

// Returns all nodes in order
func (tree *RBTree) inOrder() []*redBlackNode {
	nodes := make([]*redBlackNode, tree.size)
	if tree.size > 0 {
		current := tree.root
		stack := stacks.NewLinkedListStack()
		done := false
		count := 0
		for !done {
			if current != nil {
				stack.Push(current)
				current = current.left
			} else {
				if !stack.Empty() {
					currentPop, _ := stack.Pop()
					current = currentPop.(*redBlackNode)
					nodes[count] = current
					count += 1
					current = current.right
				} else {
					done = true
				}
			}
		}
	}
	return nodes
}

func (tree *RBTree) deleteCase1(node *redBlackNode) {
	if node.parent == nil {
		return
	} else {
		tree.deleteCase2(node)
	}
}

func (tree *RBTree) deleteCase2(node *redBlackNode) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.parent.color = red
		sibling.color = black
		if node == node.parent.left {
			tree.rotateLeft(node.parent)
		} else {
			tree.rotateRight(node.parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *RBTree) deleteCase3(node *redBlackNode) {
	sibling := node.sibling()
	if nodeColor(node.parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == black &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		tree.deleteCase1(node.parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RBTree) deleteCase4(node *redBlackNode) {
	sibling := node.sibling()
	if nodeColor(node.parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == black &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		node.parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RBTree) deleteCase5(node *redBlackNode) {
	sibling := node.sibling()
	if node == node.parent.left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == red &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		sibling.left.color = black
		tree.rotateRight(sibling)
	} else if node == node.parent.right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.right) == red &&
		nodeColor(sibling.left) == black {
		sibling.color = red
		sibling.right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *RBTree) deleteCase6(node *redBlackNode) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.parent)
	node.parent.color = black
	if node == node.parent.left && nodeColor(sibling.right) == red {
		sibling.right.color = black
		tree.rotateLeft(node.parent)
	} else if nodeColor(sibling.left) == red {
		sibling.left.color = black
		tree.rotateRight(node.parent)
	}
}

func (tree *RBTree) lookup(key interface{}) *redBlackNode {
	node := tree.root
	for node != nil {
		compare := tree.comparator(key, node.key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.left
		case compare > 0:
			node = node.right
		}
	}
	return nil
}

func (tree *RBTree) insertCase1(node *redBlackNode) {
	if node.parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RBTree) insertCase2(node *redBlackNode) {
	if nodeColor(node.parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *RBTree) insertCase3(node *redBlackNode) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RBTree) insertCase4(node *redBlackNode) {
	grandparent := node.grandparent()
	if node == node.parent.right && node.parent == grandparent.left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandparent.right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.insertCase5(node)
}

func (tree *RBTree) insertCase5(node *redBlackNode) {
	node.parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.parent.left && node.parent == grandparent.left {
		tree.rotateRight(grandparent)
	} else if node == node.parent.right && node.parent == grandparent.right {
		tree.rotateLeft(grandparent)
	}
}

func (tree *RBTree) rotateLeft(node *redBlackNode) {
	right := node.right
	tree.replaceNode(node, right)
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
	node.parent = right
}

func (tree *RBTree) rotateRight(node *redBlackNode) {
	left := node.left
	tree.replaceNode(node, left)
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
	node.parent = left
}

func (tree *RBTree) replaceNode(old *redBlackNode, new *redBlackNode) {
	if old.parent == nil {
		tree.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

func nodeColor(node *redBlackNode) color {
	if node == nil {
		return black
	}
	return node.color
}
