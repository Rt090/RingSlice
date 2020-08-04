package main

import (
	"errors"
	"fmt"
)

// Slice struct
type Slice struct {
	values []int
	used   int
	start  int
	end    int
	debug  bool
}

// NewSlice creates an empty ring with the capacity provided
func NewSlice(capacity int, debug bool) *Slice {
	return &Slice{values: make([]int, capacity), debug: debug}
}

// Append pushes a value in next open position on ring
func (s *Slice) Append(value int) error {
	if s.used == cap(s.values) {
		return errors.New("Index is full cannot append")
	}
	ind := s.trueIndex(s.start, s.used) // next index is same as num written
	if s.debug {
		fmt.Println("ind to write to ", ind)
	}
	s.values[ind] = value
	s.used++
	return nil
}

// Fetch returns the value at the position
// todo return err?
func (s *Slice) Fetch(index int) int {
	return s.values[index]
}

// DeleteBounds removes all the values [start,end]
// start to end-1 will wipe slice completely
// start to start will wipe slice completely
func (s *Slice) DeleteBounds(start, end int) {
	stop := s.next(end)
	if stop == start || end == start {
		s.values = make([]int, len(s.values))
	}
	for i := start; i != stop; {
		s.values[i] = 0
		i = s.next(i)
		s.used-- // TODO: could speed this up with calculation
	}
	s.start = stop
}

// DeleteLen deletes 'length' number of entries in the ring starting at start
// anything >= capacity of slice allocates a new slice
func (s *Slice) DeleteLen(start, length int) {
	ind := start
	if length >= len(s.values) {
		s.values = make([]int, len(s.values))
		return
	}
	for i := 0; i < length; i++ {
		s.values[ind] = 0
		ind = s.next(ind)
	}
	s.used -= length
	s.start = ind
}

// returns index in slice corresponding to next index in the ring
func (s *Slice) next(cur int) int {
	if cur == cap(s.values)-1 {
		return 0
	}
	return cur + 1
}

// returns index of length AWAY from start taking into account wrap around
// start,0 == start
func (s *Slice) trueIndex(start, length int) int {
	if s.debug {
		fmt.Println("true index of start", start, "length", length, "cap", cap(s.values))
	}
	return (start + length) % cap(s.values)
}

func (s *Slice) firstFree() int {
	for i := 0; i < len(s.values); i++ {
		if s.values[i] == 0 { // TODO obviously this is a poor check, need an allocation variable?
			return i
		}
	}
	return -1
}
