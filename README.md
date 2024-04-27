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
  - [buntdb](https://github.com/tidwall/buntdb)
  - [LevelDB](https://github.com/syndtr/goleveldb)
  - [cznic/kv](https://github.com/cznic/kv)
  - [rocksdb](https://github.com/tecbot/gorocksdb)
  - [pebble](https://github.com/petermattis/pebble)
  - [pogreb](https://github.com/akrylysov/pogreb)
  - [nutsdb](https://github.com/nutsdb/nutsdb)
  - [sniper](https://github.com/recoilme/sniper)
  - map (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
  - btree (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
- Option to disable fsync
- Compatible with Redis clients


## SSD benchmark
The following benchmarks show the throughput of inserting/reading keys (of size
9 bytes) and values (of size 256 bytes).

nofsync - throughputs

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|
|batch-write-test|3384000|660000|720000|1314000|55000|2645000|762000|685000|12214000|1874000|3244000|3453000|11071000|7102000|
|set|241790|17499|14602|347061|13181|118990|547942|126504|1391978|424581|254951|1967990|293287|1671227|
|get|420646|477658|509576|567713|129678|2964136|44271|3431627|1834835|893392|3476483|3778381|4389113|3368191|
|getmixed|398860|385322|368482|416508|85898|589247|584372|1349474|1355336|840717|1326735|1774149|2062676|2634345|
|del|377813|18244|15259|310687|129441|319978|579919|391222|1587662|358787|806879|2476898|290704|2834749|
|setmixed|3023|6467|5485|19857|3577|18031|813147|45260|23670|19802|36790|38833|43765|42800|


nofsync - time

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|
|getmixed|104|108|113|100|485|70|71|30|30|49|31|23|20|15|
|batch-write-test|5.023055292s|5.288278292s|5.198012375s|5.016727375s|9.167386s|5.092892375s|5.3318445s|5.101601292s|5.004326459s|5.026614625s|5.042461625s|5.041533083s|5.008375666s|5.022207s|
|set|172|2381|2853|120|3161|350|76|329|29|98|163|21|142|24|
|get|99|87|81|73|321|14|941|12|22|46|11|11|9|12|
|setmixed|330745|154610|182296|50358|279532|55458|1229|22094|42246|50498|27180|25750|22849|23363|
|del|110|2283|2730|134|321|130|71|106|26|116|51|16|143|14|


fsync - throughputs

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|
|set|14781|120|118|2875|13767|226|2745|235|1393568|429444|235|2005780|244|1662746|
|setmixed|6826|121|119|245|3521|222|247|232|23810|19142|244|38023|242|44382|
|getmixed|233708|340951|341819|615106|84561|5389|1176337|5632|1353489|835564|6175|1714657|6337|2518278|
|del|12401|118|120|43408|134910|239|2929|234|1571754|178756|243|2513004|246|2761175|
|batch-write-test|2635000|535000|557000|798000|55000|803000|651000|24000|12352000|2007000|749000|3371000|1098000|7048000|
|get|517107|531439|562038|606780|133530|3526076|1117863|3506479|1837513|886139|3443823|3784528|4421023|3421921|


fsync - time

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|pogreb|nutsdb|sniper|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|--|--|
|getmixed|178|122|121|67|492|7731|35|7397|30|49|6746|24|6574|16|
|del|3359|352141|344480|959|308|173866|14221|177974|26|233|171371|16|169202|15|
|batch-write-test|5.066608791s|5.179667709s|5.161143334s|6.069936083s|9.243754291s|5.167015666s|5.764243667s|1m39.120360083s|5.004542667s|5.035737041s|5.179057333s|5.046837417s|5.101120875s|5.023446625s|
|set|2818|347109|352946|14487|3026|183638|15174|176701|29|97|176974|20|170134|25|
|get|80|78|74|68|312|11|37|11|22|47|12|11|9|12|
|setmixed|146484|8199282|8361364|4075094|283975|4494632|4032421|4302106|41999|52240|4095404|26299|4116504|22531|



