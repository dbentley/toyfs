package kv

import (
	"fmt"
)

type Inode uint64

const RootInode Inode = 0

type KVStore interface {
	Get(i Inode) ([]byte, error)
	Size(i Inode) (int, error)
	Set(i Inode, data []byte) error
	NewInode() Inode

	// TODO(dbentley): delete, if we ever want to reclaim space
}

type MemoryKVStore struct {
	inodes    map[Inode][]byte
	nextInode Inode
}

func NewMemoryKVStore() *MemoryKVStore {
	return &MemoryKVStore{
		inodes: make(map[Inode][]byte),
	}
}

func (s *MemoryKVStore) Get(i Inode) ([]byte, error) {
	d, ok := s.inodes[i]
	if !ok {
		return nil, fmt.Errorf("no such inode: %d", i)
	}

	r := make([]byte, len(d))
	copy(r, d)

	return r, nil
}

func (s *MemoryKVStore) Size(i Inode) (int, error) {
	d, ok := s.inodes[i]
	if !ok {
		return 0, fmt.Errorf("no such inode: %d", i)
	}

	return len(d), nil
}

func (s *MemoryKVStore) Set(i Inode, data []byte) error {
	c := make([]byte, len(data))
	copy(c, data)
	s.inodes[i] = c
	return nil
}

func (s *MemoryKVStore) NewInode() Inode {
	s.nextInode++

	return s.nextInode
}
