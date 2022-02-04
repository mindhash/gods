// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

import "github.com/emirpasic/gods/containers"

func assertIteratorImplementation() {
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

// Iterator holding the iterator's state
type Iterator struct {
	list    *List
	index   int
	element *Element
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1, element: nil}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == 0 {
		iterator.element = iterator.list.First
	} else {
		iterator.element = iterator.element.Next
	}
	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	return iterator.element.Value
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

func (iterator *Iterator) End() bool {
	element := iterator.list.Last
	if element != nil {
		iterator.element = element 
		iterator.index = iterator.list.size - 1
	}
	return false
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}
