package main

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Get(key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

func main() {
	m1 := Map[string, int]{"foo": 1, "bar": 2}
	m2 := Map[string, string]{"foo": "bar", "bar": "baz"}

	println(m1.Get("foo"))
	println(m2.Get("foo"))

}
