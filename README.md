# KVBench

cloned from [tidwall/kvbench](https://github.com/tidwall/kvbench)

KVBench is a Redis server clone backed by a few different Go databases. 

It's intended to be used with the `redis-benchmark` command to test the
performance of various Go databases.  It has support for redis pipelining. The
`redis-benchmark` can run as explained here https://github.com/tidwall/kvbench#examples.

This cloned version adds more kv databases and automatic scripts.

Features:

- Databases
  - [badger](https://github.com/dgraph-io/badger)
  - [BboltDB](https://go.etcd.io/bbolt)
  - [BoltDB](https://github.com/boltdb/bolt)
  - [buntdb](https://github.com/tidwall/buntdb)
  - [LevelDB](https://github.com/syndtr/goleveldb)
  - [cznic/kv](https://github.com/cznic/kv)
  - [rocksdb](https://github.com/tecbot/gorocksdb)
  - [pebble](https://github.com/petermattis/pebble)
  - [pogreb](https://github.com/akrylysov/pogreb)
  - [nutsdb](https://github.com/xujiajun/nutsdb)
  - [sniper](https://github.com/recoilme/sniper)
  - map (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
  - btree (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
- Option to disable fsync
- Compatible with Redis clients


## SSD benchmark
The following benchmarks show the throughput of inserting/reading keys (of size
9 bytes) and values (of size 256 bytes).

### nofsync

**throughputs** 

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory|     
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|      
|batch-write-test|2514000|432000|448000|931000|43000|1708000|972000|170000|991000|4648000|2076000|2267000|6934000|7206000|        
|set|209791|8152|8134|211668|74510|46138|529154|28800|128611|939618|198024|950633|191202|1318976|       
|setmixed|20218|5969|5995|79730|5273|16474|996085|19885|38217|65564|51758|64962|75396|73277|
|get|772390|974693|917747|965861|115413|5254873|600809|12228226|2786182|2100741|8212607|8513291|12936061|13110160|
|getmixed|640706|659506|607481|738533|84381|269054|944505|321072|616060|2231092|863158|1443472|1441690|1864505|
|del|259927|8281|8014|224542|115335|98340|1299895|211524|141109|911702|799415|1700040|1300773|1915467|

**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory| 
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|
|batch-write-test|5.0349167s|5.1190745s|5.5172247s|5.6663517s|8.4884269s|5.0601845s|5.2561413s|5.2989499s|5.0844088s|5.0186074s|5.0444394s|5.0413249s|5.0356747s|5.0277986s|
|set|297|7666|7683|295|838|1354|118|2170|485|66|315|65|326|47|
|setmixed|49460|167515|166788|12542|189628|60699|1003|50286|26166|15252|19320|15393|13263|13646|
|getmixed|97|94|102|84|740|232|66|194|101|28|72|43|43|33|
|get|80|64|68|64|541|11|104|5|22|29|7|7|4|4|
|del|240|7547|7798|278|541|635|48|295|442|68|78|36|48|32|


### fsync

**throughputs**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|
|batch-write-test|2330000|320000|160000|856000|43000|1236000|972000|16000|28000|4588000|1492000|2269000|3473000|7446000|
|get|787602|970840|946604|1011293|130083|7050809|845326|12873568|2982659|2409954|9204621|8879275|13696542|12382005|
|del|243913|546|899|104720|130612|1211|19355|2889|2465|935412|6693|1640712|31|2158668|
|set|210077|835|129|16856|70414|1302|17385|1576|2455|930298|2387|997449|2389|1296671|
|setmixed|14865|711|401|2083|5761|1389|2097|1758|2348|65261|2314|65336|322|72606|
|getmixed|628961|870772|887451|952747|92198|23432|761016|28612|37592|2230702|37052|1479172|5877|1847649|



**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|
|getmixed|99|71|70|65|677|2667|82|2184|1662|28|1686|42|10633|33|
|del|256|114420|69518|596|478|51569|3229|21632|25351|66|9337|38|1971955|28|
|set|297|74819|483958|3707|887|48001|3594|39646|25454|67|26182|62|26153|48|
|setmixed|67271|1405836|2488978|479992|173553|719898|476738|568791|425862|15322|432052|15305|3105572|13772|
|batch-write-test|5.0226171s|5.2235246s|5.1091474s|5.6505237s|8.4703233s|5.0930754s|5.2725028s|18.0664595s|11.3959573s|5.0100718s|5.1069484s|5.0407096s|5.0452549s|5.0224686s|
|get|79|64|66|61|480|8|73|4|20|25|6|7|4|5|


