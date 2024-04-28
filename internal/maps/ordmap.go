package maps

type OrderedMap[K comparable, V any] interface {
	Get(K) (V, bool)
	Put(K, V)
	Delete(K)
	Iter(func(K, V) bool) bool
	Slice() []V
}

func NewOrdered[K comparable, V any]() OrderedMap[K, V] {
	return &orderedMap[K, V]{
		order:   []K{},
		mapping: map[K]V{},
	}
}

type orderedMap[K comparable, V any] struct {
	order   []K
	mapping map[K]V
}

func (om *orderedMap[K, V]) Put(key K, value V) {
	if _, ok := om.mapping[key]; !ok {
		om.order = append(om.order, key)
	}
	om.mapping[key] = value
}

func (om *orderedMap[K, V]) Get(key K) (V, bool) {
	c, ok := om.mapping[key]
	return c, ok
}

func (om *orderedMap[K, V]) Iter(yield func(K, V) bool) bool {
	for _, k := range om.order {
		v := om.mapping[k]
		if ok := yield(k, v); !ok {
			return ok
		}
	}
	return true
}

func (om *orderedMap[K, V]) Delete(key K) {
	odr := []K{}
	for i, k := range om.order {
		if k == key {
			continue
		}
		odr = append(odr, om.order[i])
	}
	om.order = odr

	delete(om.mapping, key)
}

func (om *orderedMap[K, V]) Slice() []V {
	ret := make([]V, len(om.order))
	for i, k := range om.order {
		v := om.mapping[k]
		ret[i] = v
	}
	return ret
}

type ChainMap[K comparable, V any] interface {
	Push(leyer map[K]V)
	Pop() (map[K]V, bool)
	Get(K) (V, bool)
	ToMap() map[K]V
	Copy() ChainMap[K, V]
}

type chainmap[K comparable, V any] struct {
	chain []map[K]V
}

func (cm *chainmap[K, V]) Push(layer map[K]V) {
	cm.chain = append(cm.chain, layer)
}

func (cm *chainmap[K, V]) Pop() (map[K]V, bool) {
	l := len(cm.chain)
	if l == 0 {
		return map[K]V{}, false
	}
	newChain, ret := cm.chain[:l-1], cm.chain[l-1]
	cm.chain = newChain
	return ret, true
}

func (cm *chainmap[K, V]) Get(key K) (V, bool) {
	for i := len(cm.chain) - 1; 0 <= i; i -= 1 {
		c, ok := cm.chain[i][key]
		if ok {
			return c, true
		}
	}
	return *new(V), false
}

func (cm *chainmap[K, V]) ToMap() map[K]V {
	m := map[K]V{}
	for i := range cm.chain {
		c := cm.chain[i]
		for k := range c {
			m[k] = c[k]
		}
	}
	return m
}

func (cm *chainmap[K, V]) Copy() ChainMap[K, V] {
	chain := make([]map[K]V, len(cm.chain))
	copy(chain, cm.chain)
	return &chainmap[K, V]{chain: chain}
}

func NewChain[K comparable, V any]() ChainMap[K, V] {
	return &chainmap[K, V]{}
}
