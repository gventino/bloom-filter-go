package bloomfilter

import (
	"hash"
	"hash/fnv"
	"sync"
)

type BloomFilterInterface interface {
	Add(item []byte)
	Exists(item []byte) bool
	MultipleAdd(items [][]byte)
	MultipleExists(items [][]byte) []bool
}

type BloomFilter struct {
	Bits     []bool
	Hasher   hash.Hash64
	BitsLen  uint64
	ChainLen uint16
	mu       sync.RWMutex
}

func NewBloomFilterWithFNVa(bitsLen uint64, hashersLen uint16) *BloomFilter {
	return &BloomFilter{
		Bits:     make([]bool, bitsLen),
		Hasher:   fnv.New64a(),
		BitsLen:  bitsLen,
		ChainLen: hashersLen,
	}
}

func (bf *BloomFilter) Add(item []byte) {
	bf.mu.Lock()

	var index uint64
	defer bf.mu.Unlock()
	defer bf.Hasher.Reset()

	for range bf.ChainLen {
		bf.Hasher.Write(item)
		index = bf.Hasher.Sum64() % bf.BitsLen
		bf.Bits[index] = true
	}
}

func (bf *BloomFilter) Exists(item []byte) bool {
	bf.mu.Lock()

	var i uint64
	defer bf.mu.Unlock()
	defer bf.Hasher.Reset()

	for range bf.ChainLen {
		bf.Hasher.Write(item)
		i = bf.Hasher.Sum64() % bf.BitsLen
		if !bf.Bits[i] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) MultipleExists(items [][]byte) []bool {
	exists := make([]bool, len(items))
	for i, item := range items {
		exists[i] = bf.Exists(item)
	}
	return exists
}

func (bf *BloomFilter) MultipleAdd(items [][]byte) {
	for _, item := range items {
		bf.Add(item)
	}
}
