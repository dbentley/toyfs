package main

import (
	"context"
	"flag"
	"log"

	"github.com/dbentley/toyfs/cp"
	"github.com/dbentley/toyfs/fs/kvfs"
	"github.com/dbentley/toyfs/fs/real"
	"github.com/dbentley/toyfs/kv"
)

var dirP = flag.String("dir", "", "directory to scan")
var memory = flag.Bool("memory", false, "write to kvfs in the middle")

func main() {
	flag.Parse()
	ctx := context.Background()

	if *dirP == "" {
		log.Fatalf("--dir must be non-empty")
	}

	dir, err := real.NewRealFS(*dirP)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("running du against real")

	if err := DU(ctx, dir); err != nil {
		log.Fatal(err)
	}

	if !*memory {
		return
	}

	k := kv.NewMemoryKVStore()
	m, err := kvfs.NewKVFS(k)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("copying into kvfs")

	if err := cp.CP(ctx, dir, m); err != nil {
		log.Fatal(err)
	}

	log.Printf("running du against kvfs")

	if err := DU(ctx, m); err != nil {
		log.Fatal(err)
	}

	log.Printf("done")
}
