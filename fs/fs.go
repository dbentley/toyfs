package fs

import (
	"context"

	"github.com/dbentley/toyfs/kv"
)

type FD int

type FileType int

const (
	FILE FileType = iota
	DIRECTORY
)

type Dirent struct {
	Name  string
	Type  FileType
	Inode kv.Inode
}

type Stat struct {
	Type  FileType
	Inode kv.Inode
	Size  uint32
}

type FS interface {
	Read(ctx context.Context, path string) ([]byte, error)

	Write(ctx context.Context, path string, data []byte) error

	GetDents(ctx context.Context, path string) ([]Dirent, error)

	Stat(ctx context.Context, path string) (Stat, error)

	Mkdir(ctx context.Context, path string) error
}
