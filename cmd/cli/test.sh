#!/bin/sh

SIZE=256

STORES=("badger" "bbolt" "bolt" "leveldb" "kv" "buntdb" "pebble" "pogreb" "nutsdb" "sniper"  "btree" "btree/memory" "map" "map/memory")

export LD_LIBRARY_PATH=/usr/local/lib

# CGO_CFLAGS="-I/usr/local/include/rocksdb" CGO_LDFLAGS="-L/usr/local/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4"   go build main.go

`rm  -fr .*db`
`rm  -fr *.db`
`rm  -fr pogreb.*`
`rm -f benchmarks/test.log`

echo "=========== test nofsync ==========="
for i in "${STORES[@]}"
do
	./cli.exe -d 5s -size ${SIZE} -s "$i" >> benchmarks/test.log 2>&1
done

`rm  -fr .*db`
`rm  -fr *.db`
`rm  -fr pogreb.*`

echo ""
echo "=========== test fsync ==========="

for i in "${STORES[@]}"
do
	./cli.exe -d 5s -size ${SIZE} -s "$i" -fsync >> benchmarks/test.log 2>&1
done

`rm  -fr .*db` 
`rm  -fr *.db`
`rm  -fr pogreb.*`