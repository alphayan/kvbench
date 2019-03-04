package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/smallnest/hrtime"
	"github.com/smallnest/kvbench"
	"github.com/smallnest/log"
)

var (
	n      = flag.Int("n", 10000, "count")
	c      = flag.Int("c", runtime.GOMAXPROCS(-1), "concurrent goroutines")
	size   = flag.Int("size", 256, "data size")
	fsync  = flag.Bool("fsync", false, "fsync")
	memory = flag.Bool("memory", false, "fsync")
	s      = flag.String("s", "map", "store type")
)

func main() {
	flag.Parse()

	var path string
	if *memory {
		path = ":memory:"
	}
	store, path, err := getStore(*s, *fsync, path)
	if err != nil {
		panic(err)
	}
	if !*memory {
		defer os.RemoveAll(path)
	}

	data := make([]byte, *size)
	numPerG := *n / (*c)

	// test set
	{
		var wg sync.WaitGroup
		wg.Add(*c)
		benchmarks := make([]*hrtime.Benchmark, *c)

		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {
				b := hrtime.NewBenchmark(numPerG)
				benchmarks[index] = b
				i := index
				for b.Next() {
					store.Set(genKey(i), data)
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		bench := hrtime.MergeBenchmarks(benchmarks...)
		fmt.Println(bench.Histogram(10))
	}

	// test get
	{
		var wg sync.WaitGroup
		wg.Add(*c)
		benchmarks := make([]*hrtime.Benchmark, *c)

		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {
				b := hrtime.NewBenchmark(numPerG)
				benchmarks[index] = b

				i := index
				for b.Next() {
					store.Get(genKey(i))
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		bench := hrtime.MergeBenchmarks(benchmarks...)
		fmt.Println(bench.Histogram(10))
	}

	// test multiple get/one set
	{
		var wg sync.WaitGroup
		wg.Add(*c)
		benchmarks := make([]*hrtime.Benchmark, *c)

		ch := make(chan struct{})
		go func() {
			i := uint64(0)
			for {
				select {
				case <-ch:
					return
				default:
					store.Set(genKey(i), data)
					i++
				}
			}
		}()
		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {
				b := hrtime.NewBenchmark(numPerG)
				benchmarks[index] = b

				i := index
				for b.Next() {
					store.Get(genKey(i))
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(ch)
		bench := hrtime.MergeBenchmarks(benchmarks...)
		fmt.Println(bench.Histogram(10))
	}

	// test del
	{
		var wg sync.WaitGroup
		wg.Add(*c)
		benchmarks := make([]*hrtime.Benchmark, *c)

		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {
				b := hrtime.NewBenchmark(numPerG)
				benchmarks[index] = b

				i := index
				for b.Next() {
					store.Del(genKey(i))
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		bench := hrtime.MergeBenchmarks(benchmarks...)
		fmt.Println(bench.Histogram(10))
	}
}

func genKey(i uint64) []byte {
	r := make([]byte, 9)
	r[0] = 'k'
	binary.BigEndian.PutUint64(r[1:], i)
	return r
}

func getStore(s string, fsync bool, path string) (kvbench.Store, string, error) {
	var store kvbench.Store
	var err error
	switch s {
	default:
		err = fmt.Errorf("unknown store type: %v", s)
	case "map":
		if path == "" {
			path = "map.db"
		}
		store, err = kvbench.NewMapStore(path, fsync)
	case "btree":
		if path == "" {
			path = "btree.db"
		}
		store, err = kvbench.NewBTreeStore(path, fsync)
	case "bolt":
		if path == "" {
			path = "bolt.db"
		}
		store, err = kvbench.NewBoltStore(path, fsync)
	case "bbolt":
		if path == "" {
			path = "bbolt.db"
		}
		store, err = kvbench.NewBboltStore(path, fsync)
	case "leveldb":
		if path == "" {
			path = "leveldb.db"
		}
		store, err = kvbench.NewLevelDBStore(path, fsync)
	case "kv":
		log.Warningf("kv store is unstable")
		if path == "" {
			path = "kv.db"
		}
		store, err = kvbench.NewKVStore(path, fsync)
	case "badger":
		if path == "" {
			path = "badger"
		}
		store, err = kvbench.NewBadgerStore(path, fsync)
	case "buntdb":
		if path == "" {
			path = "buntdb.db"
		}
		store, err = kvbench.NewBuntdbStore(path, fsync)
	case "rocksdb":
		if path == "" {
			path = "rocksdb.db"
		}
		store, err = kvbench.NewRocksdbStore(path, fsync)
	}

	return store, path, err
}