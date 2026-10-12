package cache

// ---------------------------------------------------------------------------
// IN-MEMORY STORAGE
// ---------------------------------------------------------------------------

// MemoryStorage is just a map behind the Storage interface.
type MemoryStorage[K comparable, V any] struct {
	data map[K]Item[V]
}

// Compile-time check: fails to build if MemoryStorage stops satisfying Storage.
var _ Storage[string, int] = (*MemoryStorage[string, int])(nil)

func NewMemoryStorage[K comparable, V any]() *MemoryStorage[K, V] {
	return &MemoryStorage[K, V]{data: make(map[K]Item[V])}
}

func (s *MemoryStorage[K, V]) Get(key K) (Item[V], bool, error) {
	item, found := s.data[key]
	return item, found, nil
}

func (s *MemoryStorage[K, V]) Set(key K, value Item[V]) error {
	s.data[key] = value
	return nil
}

func (s *MemoryStorage[K, V]) Delete(key K) error {
	delete(s.data, key)
	return nil
}

func (s *MemoryStorage[K, V]) Len() int {
	return len(s.data)
}
