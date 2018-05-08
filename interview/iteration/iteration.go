package main

import (
	"fmt"
	"sync"
)

func main() {
	testSet()
}

func (tset *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		tset.Lock.RLock()

		for elem := range tset.m {
			ch <- elem
		}

		close(ch)
		tset.Lock.RUnlock()

	}()
	return ch
}

func New() *threadSafeSet {
	return &threadSafeSet{
		m: make(map[string]bool),
	}
}
func testSet() {
	// Initialize our threadSafeSet
	s := New()
	// Add example items
	s.Add("item1")
	s.Add("item1")
	// duplicate item
	s.Add("item2")
	fmt.Printf("%d items\n", s.Len())
	// Clear all items
	s.Clear()
	if s.IsEmpty() {
		fmt.Printf("0 items\n")
	}
	s.Add("item2")
	s.Add("item3")
	s.Add("item4")
	// Check for existence
	if s.Has("item2") {
		fmt.Println("item2 does exist")
	}
	// Remove some of our items
	s.Remove("item2")
	s.Remove("item4")
	s.Add("item5")
	fmt.Println("----use List()--")
	fmt.Println("list of all items:", s.List())
	fmt.Println("----use Iter()--")
	for item := range s.Iter() {
		fmt.Print(item.(string) + ",")
	}
}

type threadSafeSet struct {
	m    map[string]bool
	Lock sync.RWMutex
}

// Add add
func (s *threadSafeSet) Add(item string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.m[item] = true
}

// Remove deletes the specified item from the map
func (s *threadSafeSet) Remove(item string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.m, item)
}

// Has looks for the existence of an item
func (s *threadSafeSet) Has(item string) bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Len returns the number of items in a set.
func (s *threadSafeSet) Len() int {
	return len(s.List())
}

// Clear removes all items from the set
func (s *threadSafeSet) Clear() {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.m = make(map[string]bool)
}

// IsEmpty checks for emptiness
func (s *threadSafeSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// Set returns a slice of all items
func (s *threadSafeSet) List() []string {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	list := make([]string, 0)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
