package main

import "fmt"

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Get(key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

func (cm *CustomMap[K, V]) Insert(key K, value V) error {
	cm.data[key] = value
	return nil
}

type CustomMap[K comparable, V any] struct {
	data Map[K, V]
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{data: make(Map[K, V])}

}

func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("a", 1)
	fmt.Println(m1.data.Get("a"))

	m2 := NewCustomMap[int, string]()
	m2.Insert(1, "a")
	fmt.Println(m2.data.Get(1))

	// m2.Insert("a", 1) // compile error

}
