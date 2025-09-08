package kvmap

import (
	sm "github.com/chloyka/sync-map-generic"
)

func NewKVMap[K comparable, V any](key K, subVals map[K]V) sm.KVMap[K, sm.KVMap[K, V]] {
	inner := sm.KVMap[K, V]{}
	for k, v := range subVals {
		inner.Store(k, &v)
	}
	outer := sm.KVMap[K, sm.KVMap[K, V]]{}
	outer.Store(key, &inner)
	return outer
}
