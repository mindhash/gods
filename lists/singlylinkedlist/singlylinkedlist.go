// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package singlylinkedlist implements the singly-linked list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package singlylinkedlist

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/utils"
)

func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}

// List holds the elements, where each element points to the next element
type List struct {
	First *Element
	Last  *Element
	size  int
} 

type Element struct {
	Value interface{}
	Next  *Element
}
 

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &Element{Value: value}
		if list.size  == 0 {
			list.First = newElement
			list.Last = newElement
		} else {
			list.Last.Next = newElement
			list.Last = newElement
		}
		list.size ++
	}
}

// Append appends a value (one or more) at the end of the list (same as Add())
func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

// Prepend prepends a values (or more)
func (list *List) Prepend(values ...interface{}) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &Element{Value: values[v], Next: list.First}
		list.First = newElement
		if list.size  == 0 {
			list.Last = newElement
		}
		list.size++
	}
}

func (list *List) GetElement(index int) (*Element) {

	if !list.WithinRange(index) {
		return nil
	}

	element := list.First
	for e := 0; e != index; e, element = e+1, element.Next {
	}

	return element
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {

	if !list.WithinRange(index) {
		return nil, false
	}

	element := list.First
	for e := 0; e != index; e, element = e+1, element.Next {
	}

	return element.Value, true
}

// Remove removes the element at the given index from the list.
func (list *List) Remove(index int) {

	if !list.WithinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeElement *Element
	element := list.First
	for e := 0; e != index; e, element = e+1, element.Next {
		beforeElement = element
	}

	if element == list.First {
		list.First = element.Next
	}
	if element == list.Last {
		list.Last = beforeElement
	}
	if beforeElement != nil {
		beforeElement.Next = element.Next
	}

	element = nil

	list.size--
}

// Contains checks if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List) Contains(values ...interface{}) bool {

	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.First; element != nil; element = element.Next {
			if element.Value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all elements in the list.
func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size , list.size)
	for e, element := 0, list.First; element != nil; e, element = e+1, element.Next {
		values[e] = element.Value
	}
	return values
}

func (list *List) ValuesInt() []int {
	values := make([]int, list.size , list.size)
	for e, element := 0, list.First; element != nil; e, element = e+1, element.Next {
		values[e] = element.Value.(int)
	}
	return values
}

func (list *List) ValuesString() []string {
	values := make([]string, list.size , list.size)
	for e, element := 0, list.First; element != nil; e, element = e+1, element.Next {
		values[e] = element.Value.(string)
	}
	return values
}

//IndexOf returns index of provided element
func (list *List) IndexOf(value interface{}) int {
	if list.size  == 0 {
		return -1
	}
	for index, element := range list.Values() {
		if element == value {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements.
func (list *List) Empty() bool {
	return list.size  == 0
}

// Size returns number of elements within the list.
func (list *List) Size() int {
	return list.size 
}

// Clear removes all elements from the list.
func (list *List) Clear() {
	list.size  = 0
	list.First = nil
	list.Last = nil
}

// Sort sort values (in-place) using.
func (list *List) Sort(comparator utils.Comparator) {

	if list.size  < 2 {
		return
	}

	values := list.Values()
	utils.Sort(values, comparator)

	list.Clear()

	list.Add(values...)

}

// Swap swaps values of two elements at the given indices.
func (list *List) Swap(i, j int) {
	if list.WithinRange(i) && list.WithinRange(j) && i != j {
		var element1, element2 *Element
		for e, currentElement := 0, list.First; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.Next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.Value, element2.Value = element2.Value, element1.Value
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List) Insert(index int, values ...interface{}) {

	if !list.WithinRange(index) {
		// Append
		if index == list.size  {
			list.Add(values...)
		}
		return
	}

	list.size += len(values)

	var beforeElement *Element
	foundElement := list.First
	for e := 0; e != index; e, foundElement = e+1, foundElement.Next {
		beforeElement = foundElement
	}

	if foundElement == list.First {
		oldNextElement := list.First
		for i, value := range values {
			newElement := &Element{Value: value}
			if i == 0 {
				list.First = newElement
			} else {
				beforeElement.Next = newElement
			}
			beforeElement = newElement
		}
		beforeElement.Next = oldNextElement
	} else {
		oldNextElement := beforeElement.Next
		for _, value := range values {
			newElement := &Element{Value: value}
			beforeElement.Next = newElement
			beforeElement = newElement
		}
		beforeElement.Next = oldNextElement
	}
}

// Set value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List) Set(index int, value interface{}) {

	if !list.WithinRange(index) {
		// Append
		if index == list.size  {
			list.Add(value)
		}
		return
	}

	foundElement := list.First
	for e := 0; e != index; {
		e, foundElement = e+1, foundElement.Next
	}
	foundElement.Value = value
}

// String returns a string representation of container
func (list *List) String() string {
	str := "SinglyLinkedList\n"
	values := []string{}
	for element := list.First; element != nil; element = element.Next {
		values = append(values, fmt.Sprintf("%v", element.Value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *List) WithinRange(index int) bool {
	return index >= 0 && index < list.size 
}
