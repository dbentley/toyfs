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

var srcP = flag.String("src", "", "source dir")
var destP = flag.String("dest", "", "destination dir")
var memory = flag.Bool("memory", false, "write to kvfs in the middle")

func main() {
	flag.Parse()
	ctx := context.Background()

	if *srcP == "" {
		log.Fatalf("--src must be non-empty")
	}
	if *destP == "" {
		log.Fatalf("--dest must be non-empty")
	}

	src, err := real.NewRealFS(*srcP)
	if err != nil {
		log.Fatal(err)
	}

	dest, err := real.NewRealFS(*destP)
	if err != nil {
		log.Fatal(err)
	}

	if !*memory {
		log.Printf("cp from %v to %v", *srcP, *destP)
		if err := cp.CP(ctx, src, dest); err != nil {
			log.Fatal(err)
		}
		log.Printf("done")
		return
	}

	k := kv.NewMemoryKVStore()
	m, err := kvfs.NewKVFS(k)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("cp from %v to kvfs", *srcP)
	if err := cp.CP(ctx, src, m); err != nil {
		log.Fatal(err)
	}

	log.Printf("cp from kfvs to %v", *destP)

	if err := cp.CP(ctx, m, dest); err != nil {
		log.Fatal(err)
	}

	log.Printf("done")
}
