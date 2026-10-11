package main
import "fmt"
/*
	Interfaces

	Build: a Store interface with Get(key string) (string, bool) and Set(key, value string). Implement MapStore, 
	and write func Fill(s Store) that uses it.

	Try: add a second implementation, SliceStore, that keeps pairs in a slice, and pass it to Fill without changing Fill.

	You’ve got it when: you see that nothing says “implements”, and you can write the var _ Store = (*MapStore)(nil) 
	check and explain it.
*/

type Store interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

// -----------------------------------------------------------------------------
// MapStore: Map-backed implementation
// -----------------------------------------------------------------------------
type MapStore struct {
	data map[string]string
}

func (m *MapStore) Get(key string) (string, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *MapStore) Set(key, value string) {
	m.data[key] = value
}

func NewMapStore() *MapStore {
	return &MapStore{data : make(map[string]string)}
}

var _ Store = (*MapStore)(nil)

/*
	why did we add above line vat_ Store = (*MapStore)(nil)

	(*MapStore)(nil) casts nil pointer into a typed *MapStore
	pointer without allocating memory at runtime

	var _ Store = ... assigns that typed nill pointer to a variable of type Store.
	The blank identified _ tells the compiler to discard the variable so it doesn't
	trigger an unused variable error

	Why do this? Interface satisfaction is usually verified when passing a type into
	a function (like Fill(ms)). Adding this top-level line forces the compiler to veriify 
	that *MapStore satisfies Store at the exact line where the type is defined. If MapStore
	misses a method or changes a signature, compilation fails immediately at definition rather
	than down the call chain.
*/

// -----------------------------------------------------------------------------
// SliceStore: Slice-backed implementation
// -----------------------------------------------------------------------------

type entry struct {
	key string
	value string
}

type SliceStore struct {
	entries []entry 
}

func (s *SliceStore) Get(key string) (string, bool) {
	for _, e := range s.entries {
		if e.key == key {
			return e.value, true
		}
	}
	return "", false
}

func (s *SliceStore) Set(key, value string) {
	for i, e := range s.entries {
		if e.key == key {
			s.entries[i].value = value
			return
		}
	}

	s.entries = append(s.entries, entry{key: key, value: value})
}

var _ Store = (*SliceStore)(nil)

// -----------------------------------------------------------------------------
// Decoupled Helper & Demonstration
// -----------------------------------------------------------------------------

func Fill(s Store) {
	s.Set("language", "Go")
	s.Set("paradigm", "Implicit Interface")
}

func main() {
	// 1. Using MapStorew
	ms := NewMapStore()
	Fill(ms)
	lang, _ := ms.Get("language")
	fmt.Printf("MapStore language: %s\n", lang)

	// 2. Using SliceStore
	ss := &SliceStore{}
	Fill(ss)
	paradigm, _ := ss.Get("paradigm")
	fmt.Printf("SliceStore paradigm: %s\n", paradigm)
}


