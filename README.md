Windmill does a lot with filesystems, so this is a question about filesystems. And because we care about growth, not knowledge, it's meant to guide you to dig into filesystems and learn more.

We've made a go interface (fs/fs.go) that's a bit simpler than Go's os package and optimized for pedagogy.

The assignment is in two parts: first, use this interface; second, implement some of the interface.

(We've included one implementation of this interface (fs/real.RealFS) that wraps Go's os package, so you can do part 1 without worrying about the implmentation).

The first part is to write a tool to display the disk usage of a directory. Check out the Unix tool "du" for an example. For each file or directory, print its (recursive) size. (The size of one file is in the Stat struct returned by the call Stat.)

You'll need to edit cmd/toydu/lib.go to implement the DU function.

Note: your numbers might be very different from du for a few reasons. They round up to a block size, and they report in numbers of blocks. Try du -ah to see a human number.

Before the second part, there's one more function to write that uses the FS interface. You'll need to edit cp/cp.go to implement the CP function that copies files from a source FS to a destination FS. We're asking you to write this because the code is so similar to the du tool. (Warning: trying to refactor these so they share code can be a time sink; write it twice and we can discuss.)

The second part is to implement one layer of the filesystem: the tree structure of the filesystem.

Users see the filesystem as a tree (in "/foo/bar", the file "bar" lives under the directory "foo"). The filesystem has a split between (for our purposes) two layers:
1) key-value store where the key is an int (called an inode) and the value is the bytes in one file.
2) a translation layer that translates paths into inodes

Let's look at an example that has 3 directories (including root) and 4 regular files:

/
  foo/
    bar.txt: "hello world"
    baz.txt: "why doesn't the world ever say hello back?"
  README.md: "this is a small project"
  src/
    main.go: "package main\n"


The KVStore stores directories as files that list for each (direct) child the key in the KVStore and what type of file it is.
foo: {<key for /foo>, dir}
README.md: {<key for /README.md>, file}
src: {<key for /src>, dir}

0: "foo: {1, dir}; src: {2, dir}, README.md: {3, file}"
1: "bar.txt: {4, file}, baz.txt: {5, file}"
2: "main.go: {6, file}"
3: "this is a small project"
4: "hello world"
5: "why doesn't the world ever say hello back?"
6: "package main\n"

The interface for KVStore is in kv/store.go

(The key type is called an inode for historical reasons)

When the translation layer sees a request for "/foo/bar.txt", it has to:
*) lookup inode 0 (the root directory is always in inode 0)
*) parse its contents as a directory and lookup the file "foo" (it's a directory in inode 1)
*) lookup inode 1
*) parse its contents as a directory and lookup the file "bar.txt" (it's a regular file in inode 4)
*) lookup inode 4 and return the bytes

You'll need to edit fs/kvfs/kvfs.go to implement the fs.FS interface on top of kv.KVStore.

There are two binaries:

cmd/toydu: print the disk usage of a directory. This will use your DU function. If you pass the --memory flag, it will run the function twice: first with a RealFS and then it will create a KVFS and use your CP function to copy the data in and then run your DU function against KVFS.

cmd/toycp: copy from a source directory (in your real filesystem) to a destination directory (in your real filesystem). This will use your CP function. If you pass the --memory flag, it will create a KVFS as an intermediate destination and then source. This tool is useful for debugging as you go.

A handy suggestion is to use diff -r (e.g., "diff -r /source /dest") to check if the copy went well.

You'll need to modify 3 files (cp/cp.go, cmd/toydu/lib.go, fs/kvfs/kvfs.go). You're welcome to modify others for debugging, or if you want to change anything to feel better (or if you find bugs in this assignment).

During our interview, we'll talk about the code, the assignment, possible improvements/follow-ups, etc.