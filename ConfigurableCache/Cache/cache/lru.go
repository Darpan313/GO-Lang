package cache

import "container/list"

// ---------------------------------------------------------------------------
// LRU EVICTION POLICY
// ---------------------------------------------------------------------------

// LRU = Least Recently Used. A linked list keeps keys in usage order
// (front = most recent, back = least recent). The map lets us find a key's
// node in the list instantly instead of walking the list.
type LRU[K comparable] struct {
	order *list.List
	nodes map[K]*list.Element
}

// Compile-time check: fails to build if LRU stops satisfying EvictionPolicy.
var _ EvictionPolicy[string] = (*LRU[string])(nil)

func NewLRU[K comparable]() *LRU[K] {
	return &LRU[K]{order: list.New(), nodes: make(map[K]*list.Element)}
}

func (p *LRU[K]) OnInsert(key K) {
	p.nodes[key] = p.order.PushFront(key)
}

func (p *LRU[K]) OnAccess(key K) {
	if node, found := p.nodes[key]; found {
		p.order.MoveToFront(node)
	}
}

func (p *LRU[K]) OnRemove(key K) {
	if node, found := p.nodes[key]; found {
		p.order.Remove(node)
		delete(p.nodes, key)
	}
}

func (p *LRU[K]) Evict() (K, bool) {
	oldest := p.order.Back()
	if oldest == nil {
		var zero K
		return zero, false
	}
	key := oldest.Value.(K)
	p.OnRemove(key)
	return key, true
}
