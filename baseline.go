package main

import (
	"container/list"
	"fmt"
)

// StaticList strict
type StaticList struct {
	*list.List
	start int
	next  int
}

// NewLinkedList new
func NewLinkedList(cap int) *StaticList {
	l := list.New()
	for i := 0; i < cap; i++ {
		l.PushBack(0)
	}
	return &StaticList{List: l}
}

// Append new
func (s *StaticList) Append(val int) {
	fmt.Println("hello")
}
