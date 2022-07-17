// this is a study project, neither use it in production or even nor in educational purpose!
package main

import (
	"fmt"
	"hash/fnv"
	"math/big"
	"strings"
)

const (
	start_size  = 13
	resize_coef = 0.7
	grow_coef   = 2
)

type node struct {
	key   string
	value int
}

type Hashtable struct {
	data [][]node
	size uint64
}

func NewHashtable() *Hashtable {
	return &Hashtable{make([][]node, start_size, start_size), 0}
}

func (h *Hashtable) Set(key string, val int) {
	if float64(h.size)/float64(cap(h.data)) > resize_coef {
		h.resize()
	}

	idx := h.hash(key)
	h.data[idx] = append(h.data[idx], node{key, val})
	h.size++
}

func (h *Hashtable) Get(key string) (int, error) {
	slice := h.data[h.hash(key)]
	idx := sliceIndex(slice, key)
	if idx == -1 {
		return 0, fmt.Errorf("lmao")
	}

	return slice[idx].value, nil
}

func (h *Hashtable) Remove(key string) error {
	khash := h.hash(key)
	slice := h.data[khash]
	idx := sliceIndex(slice, key)
	if idx == -1 {
		return fmt.Errorf("oh my gah~")
	}

	slice[idx] = slice[len(slice)-1]
	h.data[khash] = slice[:len(slice)-1]

	return nil
}

func (h *Hashtable) Len() uint64 {
	return h.size
}

func (h *Hashtable) hash(s string) uint64 {
	bs := []byte(s)
	hash := fnv.New64()
	hash.Write(bs)

	return hash.Sum64() % uint64(cap(h.data))
}

func (h *Hashtable) resize() {
	fmt.Printf("Before Resizing: size(%v), len(%v), cap(%v)\n", h.size, len(h.data), cap(h.data))
	nextSize := findNextPrime(int64(cap(h.data) * grow_coef))
	old_data := h.data
	h.data = make([][]node, nextSize, nextSize)
	h.size = 0
	for _, slice := range old_data {
		for _, node := range slice {
			h.Set(node.key, node.value)
		}
	}
	fmt.Printf("After Resizing: size(%v), len(%v), cap(%v)\n========================\n", h.size, len(h.data), cap(h.data))
}

func sliceIndex(haystack []node, needle string) int {
	for i, v := range haystack {
		if v.key == needle {
			return i
		}
	}

	return -1
}

// potentialy inf loop
func findNextPrime(from int64) int64 {
	for {
		if big.NewInt(from).ProbablyPrime(0) {
			return from
		}
		from++
	}
}

func main() {
	ht := NewHashtable()
	idx := 0
	for i := 1; i <= 1000; i++ {
		for r := 'a'; r <= 'z'; r++ {
			ht.Set(strings.Repeat(string(r), i), idx)
			idx++
		}
	}
	fmt.Println(ht.Get("aa"))
	fmt.Println(ht.Remove("aa"))
	fmt.Println(ht.Get("aa"))
	fmt.Println(ht.Remove("aa"))
	fmt.Println(ht.Get("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"))
}
