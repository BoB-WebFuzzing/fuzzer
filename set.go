package main

import (
)

type Set struct {
    data map[string]struct{}
}

func NewSet() *Set {
    s := &Set{data: make(map[string]struct{}),}

    return s
}

func (s *Set) Add(element string) {
    s.data[element] = struct{}{}
}

func (s *Set) Remove(element string) {
    delete(s.data, element)
}

func (s *Set) Contains(element string) bool {
    _, exists := s.data[element]

    return exists
}

func (s *Set) Size() int {
    return len(s.data)
}

// func main() {
//     s := NewSet()
//     s.Add("password")
//     s.Add("password")
//     s.Add("password")
//     s.Add("id")
//
//     fmt.Println(s, s.Contains("password"))
//
//     s.Remove("password")
//
//     fmt.Println(s, s.Contains("password"))
// }
