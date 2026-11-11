# openkv-client-go

OpenKV official **Go** client. Talks to Compute Gateway `kv.v1.KVService` / `DSService` over gRPC.

```bash
go get github.com/devotionalist/openkv-client-go@latest
```

```go
import "github.com/devotionalist/openkv-client-go/kv"

cli, err := kv.Dial(ctx, "127.0.0.1:19000")
rev, err := cli.PutTTL(ctx, []byte("k"), []byte("v"), 0)
```

## API (v0.1)

**KV** (`kv`): `Dial` · `Get` / `GetOpt` · `Put` / `PutTTL` · `PutCAS` · `Delete` / `DeleteCAS` · `RawGet` / `RawPut` / `RawPutTTL` / `RawPutCAS` / `RawDelete` / `RawDeleteCAS` · `BatchPut` · `Scan` / `ListPrefix` · `TTL` / `Expire` / `Persist` · `Watch` · `Begin` (txn)

Raw* 为无事务平面（无 TSO/MVCC，revision=Raft apply index）；**勿与 Put/Get 混用同 key**。

**DS** (`ds`): Hash / Set / List / ZSet (`HSet`/`HGet`/… · `SAdd`/… · `LPush`/… · `ZAdd`/…)

## Proto

`proto/kv/v1/{kv,ds}.proto` synced from [openkv](https://github.com/devotionalist/openkv):

```bash
./scripts/sync-proto-from-openkv.sh
./scripts/gen.sh
```

## License

Apache-2.0
