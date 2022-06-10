package intmap

import (
	"testing"
)

func nextPowerOfTwo(n KeyType) KeyType {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return n
}

func filled(n KeyType) IntMap {
	m := New(64)
	for i := KeyType(0); i < n; i++ {
		m.Put(i, ValType(i))
	}
	return m
}

func TestIntMap_Put(t *testing.T) {
	filled(200000)
}

func TestIntMap_Get(t *testing.T) {
	m := filled(200)
	for i := KeyType(0); i < 200; i++ {
		v, ok := m.Get(i)
		if v != ValType(i) || !ok {
			t.Errorf("IntMap.Get() got = %v,%v, want %v,%v", v, ok, i, true)
		}
	}
	v, ok := m.Get(201)
	if v != 0 || ok {
		t.Errorf("IntMap.Get() got = %v,%v, want %v,%v", v, ok, 0, true)
	}
}

func TestIntMap_Delete(t *testing.T) {
	n := KeyType(200)
	m := filled(n)
	for i := KeyType(0); i < n; i++ {
		v, ok := m.Delete(i)
		if v != ValType(i) || !ok {
			t.Errorf("IntMap.Delete() got = %v,%v, want %v,%v", v, ok, i, true)
		}
	}
	for i := KeyType(0); i < n; i++ {
		v, ok := m.Delete(i)
		if v != ValType(0) || ok {
			t.Errorf("IntMap.Delete() got = %v,%v, want %v,%v", v, ok, 0, false)
		}
	}
	v, ok := m.Delete(n + 1)
	if v != 0 || ok {
		t.Errorf("IntMap.Delete() got = %v,%v, want %v,%v", v, ok, 0, false)
	}
}
