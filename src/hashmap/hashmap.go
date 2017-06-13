package hashmap

import (
	"linkedlist"
	"fmt"
)

type hashedMap struct {
	size, entries, alpha int
	array                []linkedlist.List
}

const alpha_default int = 2

var primes = []int{3, 7, 13, 31, 61, 127, 251, 509, 1021, 2039,
				   4093, 8191, 16381, 32749, 65521, 131071,
				   262139, 524287, 1048573, 2097143, 4194301,
				   8388593, 16777213, 33554393, 67108859,
				   134217689, 268435399, 536870909, 1073741789,
				   2147483647}

func NewHashMap(n int) *hashedMap {
	return &hashedMap{size: n, entries: 0, array: make([]linkedlist.List, n), alpha: alpha_default}
}

func (m *hashedMap) Find(key string) interface{} {
	hash := m.hashCode(key)
	list := m.array[hash]
	return list.Find(key)
}

func (m *hashedMap) Insert(key string, value interface{}) {
	if m.entries > m.size*m.alpha {
		m.rehash()
		m.Insert(key, value)
		return
	}
	hash := m.hashCode(key)
	list := m.array[hash]
	lLen := list.Length()
	list.Insert(key, value)
	m.array[hash] = list
	if list.Length() > lLen {
		m.entries++
	}
}

func (m *hashedMap) Delete(key string) bool {
	hash := m.hashCode(key)
	if hash < m.size {
		result := m.array[hash].Delete(key)
		if result {
			m.entries--
		}
		return result
	} else {
		return false
	}
}

func (m *hashedMap) rehash() {
	arrCopy := make([]linkedlist.List, m.size)
	copy(arrCopy, m.array)
	m.array = make([]linkedlist.List, m.nextPrime())
	m.entries = 0
	for _, e := range arrCopy {
		e.ForEach(func(key interface{}, value interface{}) {
			if v, ok := key.(string); ok {
				m.Insert(v, value)
			} else {
				println("Error")
			}
		})
	}
}

func (m *hashedMap) nextPrime() int {
	next := m.size
	for e := range primes {
		if e > next {
			return e
		}
	}
	return next
}

func (m *hashedMap) hashCode(s string) int {
	if s == "" {
		return 0
	} else {
		return (int(s[0]) + 128*m.hashCode(s[1:])) % m.size
	}
}

func (m *hashedMap) ForEach(f func(string, interface{})) {
	for _, e := range m.array {
		e.ForEach(func(key interface{}, value interface{}) {
			if v, ok := key.(string); ok {
				f(string(v), value)
			} else {
				println("Error")
			}
		})
	}
}

func (m *hashedMap) String() string {
	result := "{ "
	m.ForEach(func(key string, value interface{}) {
		result += key + "=" + fmt.Sprintf("%v", key) + ","
	})
	return result[:len(result)-1] + "}"
}
