package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

type rootDir struct {
	name           string
	nfiles, nbytes int64
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel
	var rootDirs []*rootDir
	fileSizes := make(chan rootDir)
	var n sync.WaitGroup
	for _, root := range roots {
		rd := rootDir{root, int64(0), int64(0)}
		rootDirs = append(rootDirs, &rd)
		n.Add(1)
		go walkDir(&rd, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case _, ok := <-fileSizes:
			if !ok {
				break loop
			}
		case <-tick:
			for _, rd := range rootDirs {
				printDiskUsage(rd.name, rd.nfiles, rd.nbytes)
			}
		}
	}
}

func printDiskUsage(rd string, nfiles, nbytes int64) {
	fmt.Printf("root: %s %d files %.1f GB\n", rd, nfiles, float64(nbytes)/1e9)
}

func walkDir(rd *rootDir, dir string, n *sync.WaitGroup, fileSizes chan<- rootDir) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(rd, subdir, n, fileSizes)
		} else {
			rd.nfiles++
			rd.nbytes += entry.Size()
			fileSizes <- *rd
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
