package real

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"github.com/dbentley/toyfs/fs"
	"github.com/dbentley/toyfs/kv"
	"github.com/dbentley/toyfs/path"
)

func NewRealFS(root string) (*FS, error) {
	_, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	return &FS{
		root: root,
	}, nil
}

type FS struct {
	root string
}

func (f *FS) realPath(p string) (string, error) {
	if err := path.MustBeAbs(p); err != nil {
		return "", err
	}

	p = p[1:] // strip off /
	return path.Join(f.root, p)
}

func (f *FS) Read(ctx context.Context, p string) ([]byte, error) {
	p, err := f.realPath(p)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(p)
}

func (f *FS) Write(ctx context.Context, p string, data []byte) error {
	p, err := f.realPath(p)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(p, data, os.FileMode(0755))
}

func (f *FS) GetDents(ctx context.Context, p string) ([]fs.Dirent, error) {
	p, err := f.realPath(p)
	if err != nil {
		return nil, err
	}

	fis, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	r := make([]fs.Dirent, len(fis))

	for i, fi := range fis {
		t := fs.FILE
		if fi.Mode().IsDir() {
			t = fs.DIRECTORY
		}
		stat, ok := fi.Sys().(*syscall.Stat_t)
		if !ok {
			return nil, fmt.Errorf("stat is not a Stat_t: %T %v", fi.Sys(), fi.Sys())
		}

		r[i] = fs.Dirent{
			Name:  fi.Name(),
			Type:  t,
			Inode: kv.Inode(stat.Ino),
		}
	}

	return r, nil
}

func (f *FS) Stat(ctx context.Context, p string) (fs.Stat, error) {
	p, err := f.realPath(p)
	if err != nil {
		return fs.Stat{}, err
	}

	fi, err := os.Stat(p)
	if err != nil {
		return fs.Stat{}, err
	}

	t := fs.FILE
	if fi.Mode().IsDir() {
		t = fs.DIRECTORY
	}
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return fs.Stat{}, fmt.Errorf("stat is not a Stat_t: %T %v", fi.Sys(), fi.Sys())
	}
	return fs.Stat{
		Type:  t,
		Inode: kv.Inode(stat.Ino),
		Size:  uint32(fi.Size()),
	}, nil
}

func (f *FS) Mkdir(ctx context.Context, p string) error {
	p, err := f.realPath(p)
	if err != nil {
		return err
	}

	return os.Mkdir(p, os.FileMode(0755))
}
