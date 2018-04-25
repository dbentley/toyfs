package kvfs

import (
	"context"
	"fmt"

	"github.com/dbentley/toyfs/fs"
	"github.com/dbentley/toyfs/kv"
)

func NewKVFS(k kv.KVStore) (*KVFS, error) {
	return &KVFS{
		kv: k,
	}, nil
}

type KVFS struct {
	kv kv.KVStore
}

func (f *KVFS) Read(ctx context.Context, p string) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (f *KVFS) Write(ctx context.Context, p string, data []byte) error {
	return fmt.Errorf("not yet implemented")
}

func (f *KVFS) GetDents(ctx context.Context, p string) ([]fs.Dirent, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (f *KVFS) Stat(ctx context.Context, p string) (fs.Stat, error) {
	return fs.Stat{}, fmt.Errorf("not yet implemented")
}

func (f *KVFS) Mkdir(ctx context.Context, p string) error {
	return fmt.Errorf("not yet implemented")
}
